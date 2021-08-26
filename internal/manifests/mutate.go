package manifests

import (
	"reflect"

	"github.com/ViaQ/logerr/log"

	"github.com/imdario/mergo"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"

	"github.com/ViaQ/logerr/kverrors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// MutateFuncFor returns a mutate function based on the
// existing resource's concrete type. It supports currently
// only the following types or else panics:
// - ConfigMap
// - Service
// - Deployment
// - StatefulSet
// - ServiceMonitor
func MutateFuncFor(existing, desired client.Object) controllerutil.MutateFn {
	return func() error {
		existingAnnotations := existing.GetAnnotations()
		mergeWithOverride(&existingAnnotations, desired.GetAnnotations())
		existing.SetAnnotations(existingAnnotations)

		existingLabels := existing.GetLabels()
		mergeWithOverride(&existingLabels, desired.GetLabels())
		existing.SetLabels(existingLabels)

		switch existing.(type) {
		case *corev1.ConfigMap:
			cm := existing.(*corev1.ConfigMap)
			wantCm := desired.(*corev1.ConfigMap)
			mutateConfigMap(cm, wantCm)

		case *corev1.Service:
			svc := existing.(*corev1.Service)
			wantSvc := desired.(*corev1.Service)
			mutateService(svc, wantSvc)

		case *appsv1.Deployment:
			dpl := existing.(*appsv1.Deployment)
			wantDpl := desired.(*appsv1.Deployment)
			mutateDeployment(dpl, wantDpl)

		case *appsv1.StatefulSet:
			sts := existing.(*appsv1.StatefulSet)
			wantSts := desired.(*appsv1.StatefulSet)
			mutateStatefulSet(sts, wantSts)

		case *monitoringv1.ServiceMonitor:
			svcMonitor := existing.(*monitoringv1.ServiceMonitor)
			wantSvcMonitor := desired.(*monitoringv1.ServiceMonitor)
			mutateServiceMonitor(svcMonitor, wantSvcMonitor)

		default:
			t := reflect.TypeOf(existing).String()
			return kverrors.New("missing mutate implementation for resource type", "type", t)
		}
		return nil
	}
}

func mergeWithOverride(dst, src interface{}) {
	err := mergo.Merge(dst, src, mergo.WithOverride)
	if err != nil {
		log.Error(err, "unable to mergeWithOverride", "dst", dst, "src", src)
	}
}

func mutateConfigMap(existing, desired *corev1.ConfigMap) {
	existing.BinaryData = desired.BinaryData
}

func mutateService(existing, desired *corev1.Service) {
	existing.Spec.Ports = desired.Spec.Ports
	mergeWithOverride(&existing.Spec.Selector, desired.Spec.Selector)
}

func mutateDeployment(existing, desired *appsv1.Deployment) {
	// Deployment selector is immutable so we set this value only if
	// a new object is going to be created
	if existing.CreationTimestamp.IsZero() {
		mergeWithOverride(existing.Spec.Selector, desired.Spec.Selector)
	}
	existing.Spec.Replicas = desired.Spec.Replicas
	mergeWithOverride(&existing.Spec.Template, desired.Spec.Template)
	mergeWithOverride(&existing.Spec.Strategy, desired.Spec.Strategy)
}

func mutateStatefulSet(existing, desired *appsv1.StatefulSet) {
	// StatefulSet selector is immutable so we set this value only if
	// a new object is going to be created
	if existing.CreationTimestamp.IsZero() {
		existing.Spec.Selector = desired.Spec.Selector
	}
	existing.Spec.PodManagementPolicy = desired.Spec.PodManagementPolicy
	existing.Spec.Replicas = desired.Spec.Replicas
	mergeWithOverride(&existing.Spec.Template, desired.Spec.Template)
	for i := range existing.Spec.VolumeClaimTemplates {
		existing.Spec.VolumeClaimTemplates[i].TypeMeta = desired.Spec.VolumeClaimTemplates[i].TypeMeta
		existing.Spec.VolumeClaimTemplates[i].ObjectMeta = desired.Spec.VolumeClaimTemplates[i].ObjectMeta
		existing.Spec.VolumeClaimTemplates[i].Spec = desired.Spec.VolumeClaimTemplates[i].Spec
	}
}

func mutateServiceMonitor(existing, desired *monitoringv1.ServiceMonitor) {
	// ServiceMonitor selector is immutable so we set this value only if
	// a new object is going to be created
}
