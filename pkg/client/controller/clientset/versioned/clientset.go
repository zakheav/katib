/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	commonv1alpha3 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/common/v1alpha3"
	commonv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/common/v1beta1"
	experimentv1alpha3 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/experiments/v1alpha3"
	experimentv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/experiments/v1beta1"
	suggestionv1alpha3 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/suggestions/v1alpha3"
	suggestionv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/suggestions/v1beta1"
	trialv1alpha3 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/trials/v1alpha3"
	trialv1beta1 "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/typed/trials/v1beta1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	CommonV1alpha3() commonv1alpha3.CommonV1alpha3Interface
	CommonV1beta1() commonv1beta1.CommonV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Common() commonv1beta1.CommonV1beta1Interface
	ExperimentV1alpha3() experimentv1alpha3.ExperimentV1alpha3Interface
	ExperimentV1beta1() experimentv1beta1.ExperimentV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Experiment() experimentv1beta1.ExperimentV1beta1Interface
	SuggestionV1alpha3() suggestionv1alpha3.SuggestionV1alpha3Interface
	SuggestionV1beta1() suggestionv1beta1.SuggestionV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Suggestion() suggestionv1beta1.SuggestionV1beta1Interface
	TrialV1alpha3() trialv1alpha3.TrialV1alpha3Interface
	TrialV1beta1() trialv1beta1.TrialV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Trial() trialv1beta1.TrialV1beta1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	commonV1alpha3     *commonv1alpha3.CommonV1alpha3Client
	commonV1beta1      *commonv1beta1.CommonV1beta1Client
	experimentV1alpha3 *experimentv1alpha3.ExperimentV1alpha3Client
	experimentV1beta1  *experimentv1beta1.ExperimentV1beta1Client
	suggestionV1alpha3 *suggestionv1alpha3.SuggestionV1alpha3Client
	suggestionV1beta1  *suggestionv1beta1.SuggestionV1beta1Client
	trialV1alpha3      *trialv1alpha3.TrialV1alpha3Client
	trialV1beta1       *trialv1beta1.TrialV1beta1Client
}

// CommonV1alpha3 retrieves the CommonV1alpha3Client
func (c *Clientset) CommonV1alpha3() commonv1alpha3.CommonV1alpha3Interface {
	return c.commonV1alpha3
}

// CommonV1beta1 retrieves the CommonV1beta1Client
func (c *Clientset) CommonV1beta1() commonv1beta1.CommonV1beta1Interface {
	return c.commonV1beta1
}

// Deprecated: Common retrieves the default version of CommonClient.
// Please explicitly pick a version.
func (c *Clientset) Common() commonv1beta1.CommonV1beta1Interface {
	return c.commonV1beta1
}

// ExperimentV1alpha3 retrieves the ExperimentV1alpha3Client
func (c *Clientset) ExperimentV1alpha3() experimentv1alpha3.ExperimentV1alpha3Interface {
	return c.experimentV1alpha3
}

// ExperimentV1beta1 retrieves the ExperimentV1beta1Client
func (c *Clientset) ExperimentV1beta1() experimentv1beta1.ExperimentV1beta1Interface {
	return c.experimentV1beta1
}

// Deprecated: Experiment retrieves the default version of ExperimentClient.
// Please explicitly pick a version.
func (c *Clientset) Experiment() experimentv1beta1.ExperimentV1beta1Interface {
	return c.experimentV1beta1
}

// SuggestionV1alpha3 retrieves the SuggestionV1alpha3Client
func (c *Clientset) SuggestionV1alpha3() suggestionv1alpha3.SuggestionV1alpha3Interface {
	return c.suggestionV1alpha3
}

// SuggestionV1beta1 retrieves the SuggestionV1beta1Client
func (c *Clientset) SuggestionV1beta1() suggestionv1beta1.SuggestionV1beta1Interface {
	return c.suggestionV1beta1
}

// Deprecated: Suggestion retrieves the default version of SuggestionClient.
// Please explicitly pick a version.
func (c *Clientset) Suggestion() suggestionv1beta1.SuggestionV1beta1Interface {
	return c.suggestionV1beta1
}

// TrialV1alpha3 retrieves the TrialV1alpha3Client
func (c *Clientset) TrialV1alpha3() trialv1alpha3.TrialV1alpha3Interface {
	return c.trialV1alpha3
}

// TrialV1beta1 retrieves the TrialV1beta1Client
func (c *Clientset) TrialV1beta1() trialv1beta1.TrialV1beta1Interface {
	return c.trialV1beta1
}

// Deprecated: Trial retrieves the default version of TrialClient.
// Please explicitly pick a version.
func (c *Clientset) Trial() trialv1beta1.TrialV1beta1Interface {
	return c.trialV1beta1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.commonV1alpha3, err = commonv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.commonV1beta1, err = commonv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.experimentV1alpha3, err = experimentv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.experimentV1beta1, err = experimentv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.suggestionV1alpha3, err = suggestionv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.suggestionV1beta1, err = suggestionv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.trialV1alpha3, err = trialv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.trialV1beta1, err = trialv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.commonV1alpha3 = commonv1alpha3.NewForConfigOrDie(c)
	cs.commonV1beta1 = commonv1beta1.NewForConfigOrDie(c)
	cs.experimentV1alpha3 = experimentv1alpha3.NewForConfigOrDie(c)
	cs.experimentV1beta1 = experimentv1beta1.NewForConfigOrDie(c)
	cs.suggestionV1alpha3 = suggestionv1alpha3.NewForConfigOrDie(c)
	cs.suggestionV1beta1 = suggestionv1beta1.NewForConfigOrDie(c)
	cs.trialV1alpha3 = trialv1alpha3.NewForConfigOrDie(c)
	cs.trialV1beta1 = trialv1beta1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.commonV1alpha3 = commonv1alpha3.New(c)
	cs.commonV1beta1 = commonv1beta1.New(c)
	cs.experimentV1alpha3 = experimentv1alpha3.New(c)
	cs.experimentV1beta1 = experimentv1beta1.New(c)
	cs.suggestionV1alpha3 = suggestionv1alpha3.New(c)
	cs.suggestionV1beta1 = suggestionv1beta1.New(c)
	cs.trialV1alpha3 = trialv1alpha3.New(c)
	cs.trialV1beta1 = trialv1beta1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}