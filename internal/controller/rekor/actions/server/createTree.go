package server

import (
	"context"
	"fmt"

	rhtasv1alpha1 "github.com/securesign/operator/api/v1alpha1"
	"github.com/securesign/operator/internal/controller/common"
	"github.com/securesign/operator/internal/controller/common/action"
	k8sutils "github.com/securesign/operator/internal/controller/common/utils/kubernetes"
	"github.com/securesign/operator/internal/controller/constants"
	"github.com/securesign/operator/internal/controller/rekor/actions"
	trillian "github.com/securesign/operator/internal/controller/trillian/actions"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCreateTrillianTreeAction() action.Action[rhtasv1alpha1.Rekor] {
	return &createTrillianTreeAction{}
}

type createTrillianTreeAction struct {
	action.BaseAction
}

func (i createTrillianTreeAction) Name() string {
	return "create Trillian tree"
}

func (i createTrillianTreeAction) CanHandle(_ context.Context, instance *rhtasv1alpha1.Rekor) bool {
	c := meta.FindStatusCondition(instance.Status.Conditions, constants.Ready)
	return c.Reason == constants.Creating && instance.Status.TreeID == nil
}

func (i createTrillianTreeAction) Handle(ctx context.Context, instance *rhtasv1alpha1.Rekor) *action.Result {
	if instance.Spec.TreeID != nil && *instance.Spec.TreeID != int64(0) {
		instance.Status.TreeID = instance.Spec.TreeID
		return i.StatusUpdate(ctx, instance)
	}
	var err error

	trillUrl, err := k8sutils.GetInternalUrl(ctx, i.Client, instance.Namespace, trillian.LogserverDeploymentName)
	if err != nil {
		return i.Failed(err)
	}
	tree, err := common.CreateTrillianTree(ctx, "rekor-tree", trillUrl+":8091", constants.CreateTreeDeadline)
	if err != nil {
		meta.SetStatusCondition(&instance.Status.Conditions, metav1.Condition{
			Type:    actions.ServerCondition,
			Status:  metav1.ConditionFalse,
			Reason:  constants.Failure,
			Message: err.Error(),
		})
		meta.SetStatusCondition(&instance.Status.Conditions, metav1.Condition{
			Type:    constants.Ready,
			Status:  metav1.ConditionFalse,
			Reason:  constants.Failure,
			Message: err.Error(),
		})
		return i.FailedWithStatusUpdate(ctx, fmt.Errorf("could not create trillian tree: %w", err), instance)
	}
	i.Recorder.Event(instance, v1.EventTypeNormal, "TreeID", "New Trillian tree created")
	instance.Status.TreeID = &tree.TreeId

	return i.StatusUpdate(ctx, instance)
}
