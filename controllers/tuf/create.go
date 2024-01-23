package tuf

import (
	"context"
	"fmt"

	rhtasv1alpha1 "github.com/securesign/operator/api/v1alpha1"
	"github.com/securesign/operator/controllers/common/action"
	"github.com/securesign/operator/controllers/common/utils/kubernetes"
	"github.com/securesign/operator/controllers/fulcio/utils"
	tufutils "github.com/securesign/operator/controllers/tuf/utils"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	tufDeploymentName  = "tuf"
	componentName      = "tuf"
	serviceAccountName = "tus-sa"
)

func NewCreateAction() action.Action[rhtasv1alpha1.Tuf] {
	return &createAction{}
}

type createAction struct {
	action.BaseAction
}

func (i createAction) Name() string {
	return "create"
}

func (i createAction) CanHandle(tuf *rhtasv1alpha1.Tuf) bool {
	return tuf.Status.Phase == rhtasv1alpha1.PhaseCreating
}

func (i createAction) Handle(ctx context.Context, instance *rhtasv1alpha1.Tuf) (*rhtasv1alpha1.Tuf, error) {
	var err error
	labels := kubernetes.FilterCommonLabels(instance.Labels)
	labels[kubernetes.ComponentLabel] = ComponentName
	labels[kubernetes.NameLabel] = tufDeploymentName

	_, err = utils.FindFulcio(ctx, i.Client, instance.Namespace, kubernetes.FilterCommonLabels(instance.Labels))
	if err != nil {
		instance.Status.Phase = rhtasv1alpha1.PhaseError
		return instance, fmt.Errorf("could not find Fulcio: %s", err)
	}

	sa := kubernetes.CreateServiceAccount(instance.Namespace, serviceAccountName, labels)
	db := tufutils.CreateTufDeployment(instance, tufDeploymentName, labels, sa.Name)
	controllerutil.SetControllerReference(instance, db, i.Client.Scheme())
	if err = i.Client.Create(ctx, db); err != nil {
		instance.Status.Phase = rhtasv1alpha1.PhaseError
		return instance, fmt.Errorf("could not create TUF: %w", err)
	}

	svc := kubernetes.CreateService(instance.Namespace, ComponentName, 8080, labels)
	//patch the pregenerated service
	svc.Spec.Ports[0].Port = 80
	controllerutil.SetControllerReference(instance, svc, i.Client.Scheme())
	if err = i.Client.Create(ctx, svc); err != nil {
		instance.Status.Phase = rhtasv1alpha1.PhaseError
		return instance, fmt.Errorf("could not create service: %w", err)
	}

	if instance.Spec.ExternalAccess.Enabled {
		ingress, err := kubernetes.CreateIngress(ctx, i.Client, *svc, instance.Spec.ExternalAccess, "tuf", labels)
		if err != nil {
			instance.Status.Phase = rhtasv1alpha1.PhaseError
			return instance, fmt.Errorf("could not create ingress: %w", err)
		}
		controllerutil.SetControllerReference(instance, ingress, i.Client.Scheme())
		if err = i.Client.Create(ctx, ingress); err != nil {
			instance.Status.Phase = rhtasv1alpha1.PhaseError
			return instance, fmt.Errorf("could not create route: %w", err)
		}
	}

	instance.Status.Phase = rhtasv1alpha1.PhaseInitialize
	return instance, nil
}
