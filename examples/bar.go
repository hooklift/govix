package main

import (
	"fmt"
	"github.com/c4milo/govix"
)

func main() {
	host, err := vix.Connect(
		"",
		0,
		"",
		"",
		vix.VMWARE_WORKSTATION,
		vix.VERIFY_SSL_CERT)

	if err != nil {
		panic(err)
	}

	defer host.Disconnect()

	vm, err := host.OpenVm("/Users/camilo/Dropbox/Development/cloudescape/dobby-boxes/ubuntu/output-vmware-iso/ubuntu1404.vmx", "")
	if err != nil {
		panic(err)
	}

	// err = vm.AddNetworkAdapter(nil)
	// if err != nil {
	// 	panic(err)
	// }

	vm.RemoveNetworkAdapter(&vix.NetworkAdapter{Id: "7"})
	vm.RemoveNetworkAdapter(&vix.NetworkAdapter{Id: "6"})
	vm.RemoveNetworkAdapter(&vix.NetworkAdapter{Id: "5"})
	vm.RemoveNetworkAdapter(&vix.NetworkAdapter{Id: "4"})
	vm.RemoveNetworkAdapter(&vix.NetworkAdapter{Id: "3"})
	vm.RemoveNetworkAdapter(&vix.NetworkAdapter{Id: "2"})
	vm.RemoveNetworkAdapter(&vix.NetworkAdapter{Id: "1"})

	netAdapters, err := vm.NetworkAdapters()
	for _, adapter := range netAdapters {
		fmt.Printf("%v\n", adapter)
	}

	// toolState, err := vm.ToolState()
	// if err != nil {
	// 	panic(err)
	// }

	// if toolState != vix.TOOLSSTATE_RUNNING {
	// 	fmt.Printf("VMware Tools is not present!!! %d", toolState)
	// }
}
