package ssh

import (
	"context"

	"github.com/kris-nova/logger"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/awsapi"
	"github.com/weaveworks/eksctl/pkg/ssh/client"
	"github.com/weaveworks/eksctl/pkg/utils/file"
)

// LoadKey loads the SSH public key specified in NodeGroupSSH and returns it. The key should be specified
// in only one way: by name (for a key existing in EC2), by path (for a key in a local file)
// or by its contents (in the config-file). It also assumes that if ssh is enabled (SSH.Allow
// == true) then one key was specified
func LoadKey(ctx context.Context, sshConfig *api.NodeGroupSSH, clusterName, nodeGroupName string, ec2API awsapi.EC2) (string, error) {
	if sshConfig.Allow == nil || !*sshConfig.Allow {
		return "", nil
	}

	switch {

	// Load Key by content
	case sshConfig.PublicKey != nil:
		keyName, err := client.LoadKeyByContent(ctx, sshConfig.PublicKey, clusterName, nodeGroupName, ec2API)
		if err != nil {
			return "", err
		}
		return keyName, nil

	// Use key by name in EC2
	case sshConfig.PublicKeyName != nil && *sshConfig.PublicKeyName != "":
		if err := client.CheckKeyExistsInEC2(ctx, ec2API, *sshConfig.PublicKeyName); err != nil {
			return "", err
		}
		logger.Info("using EC2 key pair %q", *sshConfig.PublicKeyName)
		return *sshConfig.PublicKeyName, nil

	// Local ssh key file
	case file.Exists(*sshConfig.PublicKeyPath):
		keyName, err := client.LoadKeyFromFile(ctx, *sshConfig.PublicKeyPath, clusterName, nodeGroupName, ec2API)
		if err != nil {
			return "", err
		}
		return keyName, nil

	// A keyPath, when specified as a flag, can mean a local key (checked above) or a key name in EC2
	default:
		err := client.CheckKeyExistsInEC2(ctx, ec2API, *sshConfig.PublicKeyPath)
		if err != nil {
			return "", err
		}
		logger.Info("using EC2 key pair %q", *sshConfig.PublicKeyPath)
		return *sshConfig.PublicKeyPath, nil
	}

}
