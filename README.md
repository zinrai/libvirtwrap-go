# libvirtwrap-go

libvirtwrap-go is a Go library that wraps libvirt-related commands such as virsh, virt-customize, virt-resize, and qemu-img. It provides a interface for common VM operations across various virtualization technologies supported by libvirt.

## Features

- Check if a VM is running
- Set CPU count for a VM
- Set memory size for a VM
- Resize and expand VM disks
- Get a list of all defined VMs

## Installation

To install libvirtwrap-go, use `go get`:

```
$ go get github.com/zinrai/libvirtwrap-go
```

## Usage

Here's a quick example of how to use libvirtwrap-go:

```go
package main

import (
	"fmt"
	"log"

	"github.com/zinrai/libvirtwrap-go/pkg/vm"
)

func main() {
	myVM := vm.New("my-vm-name")

	running, err := myVM.IsRunning()
	if err != nil {
		log.Fatalf("Error checking VM status: %v", err)
	}

	fmt.Printf("VM is running: %v\n", running)

	err = myVM.SetCPUCount(2)
	if err != nil {
		log.Fatalf("Error setting CPU count: %v", err)
	}

	err = myVM.SetMemorySize("4G")
	if err != nil {
		log.Fatalf("Error setting memory size: %v", err)
	}

	fmt.Println("VM configuration updated successfully")
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
