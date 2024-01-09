package eksconfig

import (
	"errors"

	"github.com/aws/aws-k8s-tester/pkg/timeutil"
)

// AddOnALB2048 defines parameters for EKS cluster
// add-on ALB 2048 service.
type AddOnALB2048 struct {
	// Enable is 'true' to create this add-on.
	Enable bool `json:"enable"`
	// Created is true when the resource has been created.
	// Used for delete operations.
	Created         bool               `json:"created" read-only:"true"`
	TimeFrameCreate timeutil.TimeFrame `json:"time-frame-create" read-only:"true"`
	TimeFrameDelete timeutil.TimeFrame `json:"time-frame-delete" read-only:"true"`

	// Namespace is the namespace to create objects in.
	Namespace string `json:"namespace"`

	// DeploymentReplicasALB is the number of ALB replicas to deploy using "Deployment" object.
	DeploymentReplicasALB int32 `json:"deployment-replicas-alb"`
	// DeploymentReplicas2048 is the number of 2048 replicas to deploy using "Deployment" object.
	DeploymentReplicas2048 int32 `json:"deployment-replicas-2048"`
	// DeploymentNodeSelector2048 is configured to overwrite existing node selector
	// for ALB 2048 deployment. If left empty, tester sets default selector.
	DeploymentNodeSelector2048 map[string]string `json:"deployment-node-selector-2048"`

	// ALBARN is the ARN of the ALB created from the service.
	ALBARN string `json:"alb-arn" read-only:"true"`
	// ALBName is the name of the ALB created from the service.
	ALBName string `json:"alb-name" read-only:"true"`
	// URL is the URL for ALB 2048 Service.
	URL string `json:"url" read-only:"true"`
}

// EnvironmentVariablePrefixAddOnALB2048 is the environment variable prefix used for "eksconfig".
const EnvironmentVariablePrefixAddOnALB2048 = AWS_K8S_TESTER_EKS_PREFIX + "ADD_ON_ALB_2048_"

// IsEnabledAddOnALB2048 returns true if "AddOnALB2048" is enabled.
// Otherwise, nil the field for "omitempty".
func (cfg *Config) IsEnabledAddOnALB2048() bool {
	if cfg.AddOnALB2048 == nil {
		return false
	}
	if cfg.AddOnALB2048.Enable {
		return true
	}
	cfg.AddOnALB2048 = nil
	return false
}

func getDefaultAddOnALB2048() *AddOnALB2048 {
	return &AddOnALB2048{
		Enable:                     false,
		DeploymentReplicasALB:      3,
		DeploymentReplicas2048:     3,
		DeploymentNodeSelector2048: make(map[string]string),
	}
}

func (cfg *Config) validateAddOnALB2048() error {
	if !cfg.IsEnabledAddOnALB2048() {
		return nil
	}
	if !cfg.IsEnabledAddOnNodeGroups() && !cfg.IsEnabledAddOnManagedNodeGroups() {
		return errors.New("AddOnALB2048.Enable true but no node group is enabled")
	}
	if cfg.AddOnALB2048.Namespace == "" {
		cfg.AddOnALB2048.Namespace = cfg.Name + "-alb-2048"
	}
	return nil
}
