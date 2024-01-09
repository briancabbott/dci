package eksconfig

import (
	"errors"

	"github.com/aws/aws-k8s-tester/pkg/timeutil"
)

// AddOnNLBHelloWorld defines parameters for EKS cluster
// add-on NLB hello-world service.
type AddOnNLBHelloWorld struct {
	// Enable is 'true' to create this add-on.
	Enable bool `json:"enable"`
	// Created is true when the resource has been created.
	// Used for delete operations.
	Created         bool               `json:"created" read-only:"true"`
	TimeFrameCreate timeutil.TimeFrame `json:"time-frame-create" read-only:"true"`
	TimeFrameDelete timeutil.TimeFrame `json:"time-frame-delete" read-only:"true"`

	// Namespace is the namespace to create objects in.
	Namespace string `json:"namespace"`

	// DeploymentReplicas is the number of replicas to deploy using "Deployment" object.
	DeploymentReplicas int32 `json:"deployment-replicas"`
	// DeploymentNodeSelector is configured to overwrite existing node selector
	// for NLB hello world deployment. If left empty, tester sets default selector.
	DeploymentNodeSelector map[string]string `json:"deployment-node-selector"`

	// NLBARN is the ARN of the NLB created from the service.
	NLBARN string `json:"nlb-arn" read-only:"true"`
	// NLBName is the name of the NLB created from the service.
	NLBName string `json:"nlb-name" read-only:"true"`
	// URL is the host name for hello-world service.
	URL string `json:"url" read-only:"true"`
}

// EnvironmentVariablePrefixAddOnNLBHelloWorld is the environment variable prefix used for "eksconfig".
const EnvironmentVariablePrefixAddOnNLBHelloWorld = AWS_K8S_TESTER_EKS_PREFIX + "ADD_ON_NLB_HELLO_WORLD_"

// IsEnabledAddOnNLBHelloWorld returns true if "AddOnNLBHelloWorld" is enabled.
// Otherwise, nil the field for "omitempty".
func (cfg *Config) IsEnabledAddOnNLBHelloWorld() bool {
	if cfg.AddOnNLBHelloWorld == nil {
		return false
	}
	if cfg.AddOnNLBHelloWorld.Enable {
		return true
	}
	cfg.AddOnNLBHelloWorld = nil
	return false
}

func getDefaultAddOnNLBHelloWorld() *AddOnNLBHelloWorld {
	return &AddOnNLBHelloWorld{
		Enable:                 false,
		DeploymentReplicas:     3,
		DeploymentNodeSelector: make(map[string]string),
	}
}

func (cfg *Config) validateAddOnNLBHelloWorld() error {
	if !cfg.IsEnabledAddOnNLBHelloWorld() {
		return nil
	}
	if !cfg.IsEnabledAddOnNodeGroups() && !cfg.IsEnabledAddOnManagedNodeGroups() {
		return errors.New("AddOnNLBHelloWorld.Enable true but no node group is enabled")
	}
	if cfg.AddOnNLBHelloWorld.Namespace == "" {
		cfg.AddOnNLBHelloWorld.Namespace = cfg.Name + "-nlb-hello-world"
	}
	return nil
}
