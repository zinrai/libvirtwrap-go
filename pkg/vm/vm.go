package vm

import (
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
