package vm

import (
	"fmt"
	"github.com/zinrai/libvirtwrap-go/pkg/virsh"
)

type VM struct {
	Name string
}

func New(name string) *VM {
	return &VM{Name: name}
}

func (v *VM) IsRunning() (bool, error) {
	return virsh.IsVMRunning(v.Name)
}

func (v *VM) SetCPUCount(count int) error {
	return virsh.SetVMCPUCount(v.Name, count)
}

func (v *VM) SetMemorySize(size string) error {
	return virsh.SetVMMemorySize(v.Name, size)
}

func (v *VM) VerifyDiskBelongsToVM(imagePath string) (bool, error) {
    diskPaths, err := virsh.GetVMDiskPaths(v.Name)
    if err != nil {
        return false, fmt.Errorf("failed to get disk paths for VM '%s': %v", v.Name, err)
    }

    for _, path := range diskPaths {
        if path == imagePath {
            return true, nil
        }
    }

    return false, nil
}
