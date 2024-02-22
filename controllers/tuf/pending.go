package tuf

import (
	"context"
	"errors"
	"fmt"
	rhtasv1alpha1 "github.com/securesign/operator/api/v1alpha1"
	"github.com/securesign/operator/controllers/common/action"
	"github.com/securesign/operator/controllers/constants"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewPendingAction() action.Action[rhtasv1alpha1.Tuf] {
	return &pendingAction{}
}

type pendingAction struct {
	action.BaseAction
}

func (i pendingAction) Name() string {
	return "pending"
}

func (i pendingAction) CanHandle(tuf *rhtasv1alpha1.Tuf) bool {
	return tuf.Status.Phase == rhtasv1alpha1.PhaseNone || tuf.Status.Phase == rhtasv1alpha1.PhasePending
}

func (i pendingAction) Handle(ctx context.Context, instance *rhtasv1alpha1.Tuf) (*rhtasv1alpha1.Tuf, error) {
	if instance.Status.Phase == rhtasv1alpha1.PhaseNone {
		instance.Status.Phase = rhtasv1alpha1.PhasePending
		meta.SetStatusCondition(&instance.Status.Conditions, v1.Condition{
			Type:    string(rhtasv1alpha1.PhaseReady),
			Status:  v1.ConditionFalse,
			Reason:  (string)(rhtasv1alpha1.PhasePending),
			Message: "Resolving keys",
		})
		return instance, nil
	}

	var errs []error

	for index, key := range instance.Spec.Keys {
		meta.SetStatusCondition(&instance.Status.Conditions, v1.Condition{
			Type:   key.Name,
			Status: v1.ConditionUnknown,
			Reason: "Resolving",
		})

		// key is copy of TufKey structure
		updated, err := i.handleKey(ctx, instance, &instance.Spec.Keys[index])
		if err != nil {
			errs = append(errs, err)
			meta.SetStatusCondition(&instance.Status.Conditions, v1.Condition{
				Type:    key.Name,
				Status:  v1.ConditionFalse,
				Reason:  "Failure",
				Message: err.Error(),
			})
			continue
		} else {
			if updated {
				if err = i.Client.Update(ctx, instance); err != nil {
					return instance, err
				}
			}
			meta.SetStatusCondition(&instance.Status.Conditions, v1.Condition{
				Type:   key.Name,
				Status: v1.ConditionTrue,
				Reason: "Ready",
			})
		}
	}

	if len(errs) > 0 {
		return instance, requeueError
	}

	instance.Status.Phase = rhtasv1alpha1.PhaseCreating
	return instance, nil
}

func (i pendingAction) handleKey(ctx context.Context, instance *rhtasv1alpha1.Tuf, key *rhtasv1alpha1.TufKey) (bool, error) {
	switch {
	case key.SecretRef == nil:
		sks, err := i.discoverSecret(ctx, instance.Namespace, key)
		if err != nil {
			return false, err
		}
		key.SecretRef = sks
		return true, nil
	case key.SecretRef != nil:
		return false, nil
	default:
		return false, errors.New(fmt.Sprintf("Unable to resolve %s key. Enable autodiscovery or set secret reference.", key.Name))
	}
}

func (i pendingAction) discoverSecret(ctx context.Context, namespace string, key *rhtasv1alpha1.TufKey) (*rhtasv1alpha1.SecretKeySelector, error) {
	labelName := constants.TufLabelNamespace + "/" + key.Name
	s, err := i.findSecret(ctx, namespace, labelName)
	if err != nil {
		return nil, err
	}
	if s != nil {
		keySelector := s.Labels[labelName]
		if keySelector == "" {
			err = errors.New(fmt.Sprintf("label %s is empty", labelName))
			return nil, err
		}
		return &rhtasv1alpha1.SecretKeySelector{
			Key: keySelector,
			LocalObjectReference: v12.LocalObjectReference{
				Name: s.Name,
			},
		}, nil
	}

	return nil, errors.New("secret not found")
}

func (i pendingAction) findSecret(ctx context.Context, namespace string, label string) (*v12.Secret, error) {
	list := &v12.SecretList{}

	selector, err := labels.Parse(label)
	listOptions := &client.ListOptions{
		LabelSelector: selector,
	}

	err = i.Client.List(ctx, list, client.InNamespace(namespace), listOptions)

	if err != nil {
		return nil, err
	}
	if len(list.Items) > 1 {
		for _, d := range list.Items {
			i.Logger.V(1).Info("list %s", d.Name, d.Labels)
		}

		return nil, errors.New("duplicate resource")
	}

	if len(list.Items) == 1 {
		return &list.Items[0], nil
	}
	return nil, nil
}