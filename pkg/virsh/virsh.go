package virsh

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsVMRunning(name string) (bool, error) {
	cmd := exec.Command("sudo", "virsh", "list", "--name", "--state-running")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to get list of running VMs: %v", err)
	}

	runningVMs := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, vm := range runningVMs {
		if vm == name {
			return true, nil
		}
	}
	return false, nil
}

func SetVMCPUCount(name string, count int) error {
	cmd := exec.Command("sudo", "virsh", "setvcpus", name, fmt.Sprintf("%d", count), "--config", "--maximum")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set maximum CPU count: %v", err)
	}

	cmd = exec.Command("sudo", "virsh", "setvcpus", name, fmt.Sprintf("%d", count), "--config")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set current CPU count: %v", err)
	}

	return nil
}

func SetVMMemorySize(name, size string) error {
	cmd := exec.Command("sudo", "virsh", "setmaxmem", name, size, "--config")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set maximum memory: %v", err)
	}

	cmd = exec.Command("sudo", "virsh", "setmem", name, size, "--config")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set current memory: %v", err)
	}

	return nil
}

func GetVMList() ([]string, error) {
	cmd := exec.Command("sudo", "virsh", "list", "--name", "--all")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get list of VMs: %v", err)
	}

	vms := strings.Split(strings.TrimSpace(string(output)), "\n")
	return vms, nil
}
