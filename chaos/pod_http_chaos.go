package chaos

import (
	"errors"
	chaosmeshv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

func init() {
	template := &chaosmeshv1alpha1.PodHttpChaosSpec{
		Rules: nil,
		TLS:   nil,
	}
	podHttpChaosRuleTemplate := chaosmeshv1alpha1.PodHttpChaosRule{
		PodHttpChaosBaseRule: chaosmeshv1alpha1.PodHttpChaosBaseRule{
			Target: chaosmeshv1alpha1.PodHttpRequest, // or chaosmeshv1alpha1.PodHttpResponse
			Selector: chaosmeshv1alpha1.PodHttpChaosSelector{
				Port:            nil,
				Path:            nil,
				Method:          nil,
				Code:            nil,
				RequestHeaders:  nil,
				ResponseHeaders: nil,
			},
			Actions: chaosmeshv1alpha1.PodHttpChaosActions{
				Abort:   nil,
				Delay:   nil,
				Replace: nil,
				Patch:   nil,
			},
		},
		Source: "",
		Port:   0,
	}

}

func NewPodHttpChaos(opts ...OptChaos) (*chaosmeshv1alpha1.PodHttpChaos, error) {
	config := ConfigChaos{}
	for _, opt := range opts {
		if opt != nil {
			opt(&config)
		}
	}

	if config.Name == "" {
		return nil, errors.New("the resource name is required")
	}
	if config.Namespace == "" {
		return nil, errors.New("the namespace is required")
	}
	if config.PodHttpChaos == nil {
		return nil, errors.New("PodHttpChaos is required")
	}

	podHttpChaos := chaosmeshv1alpha1.PodHttpChaos{}
	podHttpChaos.Name = config.Name
	podHttpChaos.Namespace = config.Namespace
	config.PodHttpChaos.DeepCopyInto(&podHttpChaos.Spec)

	if config.Labels != nil {
		podHttpChaos.Labels = config.Labels
	}
	if config.Annotations != nil {
		podHttpChaos.Annotations = config.Annotations
	}

	return &podHttpChaos, nil
}
