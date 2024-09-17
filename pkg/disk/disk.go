package disk

import (
	"fmt"
	"os/exec"
)

func ResizeAndExpandDisk(imagePath, device string, partition int, newSize string) error {
	newImagePath := imagePath + ".new"

	cmd := exec.Command("sudo", "qemu-img", "create", "-f", "qcow2", "-o", "preallocation=metadata", newImagePath, newSize)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create new image: %v", err)
	}

	cmd = exec.Command("sudo", "virt-resize", "--expand", fmt.Sprintf("/dev/%s%d", device, partition), imagePath, newImagePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to resize disk: %v", err)
	}

	cmd = exec.Command("sudo", "mv", newImagePath, imagePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to replace original image: %v", err)
	}

	return nil
}
