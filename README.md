# Overview

* Go client for the Proxmox-VE REST API (https://pve.proxmox.com/wiki/Proxmox_VE_API)

* This code is based on Go package [go-proxmox](https://github.com/luthermonson/go-proxmox)


# Proxmox API Client Go Package
A Go package containing a client for [Proxmox VE](https://www.proxmox.com/). The client implements [/api2/json](https://pve.proxmox.com/pve-docs/api-viewer/index.html) and aims to provide better sdk solution for especially [cluster-api-provider-proxmox](https://github.com/sp-yduck/cluster-api-provider-proxmox) project.

## Example Usage
```go
package main

import (
	"fmt"
	"github.com/sp-yduck/proxmox/pkg/service"
)

func main() {
    // create new client with user login
	svc, err := service.NewServiceWithLogin("proxmox_url", "root@pam", "password")
	if err != nil {
		panic(err)
	}

	// get version
	version, err := svc.Version()
	if err != nil {
		panic(err)
	}
	fmt.Println(version.Release)

	// get node with name "mynode"
	node, err := svc.Node("mynode")
	if err != nil {
		panic(err)
	}
	fmt.Println(node.Node) // mynode

	// list all virtual machines on "mynode" node
	vms, err := node.VirtualMachines()
	if err != nil {
		panic(err)
	}
	for _, vm := range vms {
		fmt.Println(vm.VMID, vm.Name)
	}
}

```
