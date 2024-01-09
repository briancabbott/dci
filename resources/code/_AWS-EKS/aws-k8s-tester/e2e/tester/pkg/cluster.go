package pkg

import (
	"errors"
	"fmt"
)

type ClusterCreator interface {
	// Initialize the creator such as downloading dependencies
	Init() (Step, error)

	// Create and wait for cluster creation
	Up() (Step, error)

	// Teardown the cluster
	TearDown() (Step, error)
}

func NewClusterCreator(config *TestConfig, testDir string, testId string) (ClusterCreator, error) {
	cluster := config.Cluster

	if cluster.Kops == nil && cluster.Eks == nil {
		return nil, fmt.Errorf("TestConfig.Cluster is not set")
	}

	if cluster.Kops != nil && cluster.Eks != nil {
		return nil, fmt.Errorf("Both Kops and Eks cluster is set")
	}

	if cluster.Kops != nil {
		return NewKopsClusterCreator(cluster.Kops, testDir, testId), nil
	}

	if cluster.Eks != nil {
		return NewEksctlClusterCreator(cluster.Eks, testDir, testId), nil
	}

	return nil, errors.New("Cluster is not specified")
}
