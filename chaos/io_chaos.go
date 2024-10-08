package chaos

import (
	"errors"
	chaosmeshv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

func NewIOChaos(opts ...OptChaos) (*chaosmeshv1alpha1.IOChaos, error) {
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
	if config.IOChaos == nil {
		return nil, errors.New("IOChaos is required")
	}

	ioChaos := chaosmeshv1alpha1.IOChaos{}
	ioChaos.Name = config.Name
	ioChaos.Namespace = config.Namespace
	config.IOChaos.DeepCopyInto(&ioChaos.Spec)

	if config.Labels != nil {
		ioChaos.Labels = config.Labels
	}
	if config.Annotations != nil {
		ioChaos.Annotations = config.Annotations
	}

	return &ioChaos, nil
}
