package actions

import (
	"context"
	"fmt"

	rhtasv1alpha1 "github.com/securesign/operator/api/v1alpha1"
	"github.com/securesign/operator/controllers/common/action"
	"github.com/securesign/operator/controllers/common/utils/kubernetes"
	"github.com/securesign/operator/controllers/constants"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func NewServiceAction() action.Action[rhtasv1alpha1.Tuf] {
	return &serviceAction{}
}

type serviceAction struct {
	action.BaseAction
}

func (i serviceAction) Name() string {
	return "create service"
}

func (i serviceAction) CanHandle(tuf *rhtasv1alpha1.Tuf) bool {
	return tuf.Status.Phase == rhtasv1alpha1.PhaseCreating || tuf.Status.Phase == rhtasv1alpha1.PhaseReady
}

func (i serviceAction) Handle(ctx context.Context, instance *rhtasv1alpha1.Tuf) *action.Result {
	var (
		err     error
		updated bool
	)

	labels := constants.LabelsFor(ComponentName, DeploymentName, instance.Name)

	svc := kubernetes.CreateService(instance.Namespace, DeploymentName, 8080, labels)
	//patch the pregenerated service
	svc.Spec.Ports[0].Port = instance.Spec.Port
	if err = controllerutil.SetControllerReference(instance, svc, i.Client.Scheme()); err != nil {
		return i.Failed(fmt.Errorf("could not set controller reference for Service: %w", err))
	}
	if updated, err = i.Ensure(ctx, svc); err != nil {
		instance.Status.Phase = rhtasv1alpha1.PhaseError
		meta.SetStatusCondition(&instance.Status.Conditions, metav1.Condition{
			Type:    string(rhtasv1alpha1.PhaseReady),
			Status:  metav1.ConditionFalse,
			Reason:  "Failure",
			Message: err.Error(),
		})
		return i.FailedWithStatusUpdate(ctx, fmt.Errorf("could not create service: %w", err), instance)
	}

	if updated {
		return i.Return()
	} else {
		return i.Continue()
	}

}
