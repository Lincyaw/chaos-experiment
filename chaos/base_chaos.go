package chaos

import chaosmeshv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"

type ConfigChaos struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string

	PodHttpChaos *chaosmeshv1alpha1.PodHttpChaosSpec
}

type OptChaos func(opt *ConfigChaos)

func WithName(name string) OptChaos {
	return func(opt *ConfigChaos) {
		opt.Name = name
	}
}
func WithNamespace(namespace string) OptChaos {
	return func(opt *ConfigChaos) {
		opt.Namespace = namespace
	}
}
func WithLabels(labels map[string]string) OptChaos {
	return func(opt *ConfigChaos) {
		opt.Labels = labels
	}
}
func WithAnnotations(annotations map[string]string) OptChaos {
	return func(opt *ConfigChaos) {
		opt.Annotations = annotations
	}
}

func WithPodHttpChaosSpec(spec *chaosmeshv1alpha1.PodHttpChaosSpec) OptChaos {
	return func(opt *ConfigChaos) {
		opt.PodHttpChaos = spec
	}
}
