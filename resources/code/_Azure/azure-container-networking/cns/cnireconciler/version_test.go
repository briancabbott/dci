package cnireconciler

import (
	"testing"

	testutils "github.com/Azure/azure-container-networking/test/utils"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/exec"
)

func newCNIVersionFakeExec(ver string) exec.Interface {
	calls := []testutils.TestCmd{
		{Cmd: []string{"/opt/cni/bin/azure-vnet", "-v"}, Stdout: ver},
	}

	fake := testutils.GetFakeExecWithScripts(calls)
	return fake
}

func TestIsDumpStateVer(t *testing.T) {
	tests := []struct {
		name    string
		exec    exec.Interface
		want    bool
		wantErr bool
	}{
		{
			name:    "bad ver",
			exec:    newCNIVersionFakeExec(`Azure CNI Version v1.4.1`),
			want:    false,
			wantErr: false,
		},
		{
			name:    "bad dirty ver",
			exec:    newCNIVersionFakeExec(`Azure CNI Version v1.4.0-2-g984c5a5e-dirty`),
			want:    false,
			wantErr: false,
		},
		{
			name:    "good ver",
			exec:    newCNIVersionFakeExec(`Azure CNI Version v1.4.7`),
			want:    true,
			wantErr: false,
		},
		{
			name:    "good dirty ver",
			exec:    newCNIVersionFakeExec(`Azure CNI Version v1.4.7-7-g7b97e1eb`),
			want:    true,
			wantErr: false,
		},
		{
			name:    "non-semver",
			exec:    newCNIVersionFakeExec(`Azure CNI Version v1.4.35_Win2019OverlayFix`),
			want:    true,
			wantErr: false,
		},
		{
			name:    "non-semver hotfix ver",
			exec:    newCNIVersionFakeExec(`Azure CNI Version v1.4.44.4`),
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isDumpStateVer(tt.exec)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
