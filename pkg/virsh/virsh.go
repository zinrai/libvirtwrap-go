package virsh

import (
    "fmt"
    "os/exec"
    "strings"
    "github.com/zinrai/libvirtwrap-go/internal/xmlutil"
)

type DomainDisk struct {
    Device string `xml:"device,attr"`
    Source struct {
        File string `xml:"file,attr"`
    } `xml:"source"`
    Target struct {
        Dev string `xml:"dev,attr"`
    } `xml:"target"`
}

type Domain struct {
    Devices struct {
        Disks []DomainDisk `xml:"disk"`
    } `xml:"devices"`
}

func GetVMDiskPaths(vmName string) ([]string, error) {
    cmd := exec.Command("sudo", "virsh", "dumpxml", vmName)
    output, err := cmd.Output()
    if err != nil {
        return nil, fmt.Errorf("failed to execute virsh dumpxml: %v", err)
    }

    var domain Domain
    if err := xmlutil.ParseXML(output, &domain); err != nil {
        return nil, fmt.Errorf("failed to parse XML: %v", err)
    }

    var diskPaths []string
    for _, disk := range domain.Devices.Disks {
        if disk.Device == "disk" {
            diskPaths = append(diskPaths, disk.Source.File)
        }
    }

    if len(diskPaths) == 0 {
        return nil, fmt.Errorf("no disks found for VM '%s'", vmName)
    }

    return diskPaths, nil
}

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
