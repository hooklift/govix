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

	fmt.Println("Searching for running vms...")

	urls, err := host.FindItems(vix.FIND_RUNNING_VMS)
	if err != nil {
		panic(err)
	}

	fmt.Println("host.findItems returned!")

	for _, url := range urls {
		vm, _ := host.OpenVm(url, "")
		fmt.Println("Url: " + url)
		vcpus, _ := vm.Vcpus()
		memsize, _ := vm.MemorySize()
		vmxpath, _ := vm.VmxPath()
		teampath, _ := vm.VmTeamPath()
		guestos, _ := vm.GuestOS()
		//features, _ := vm.Features()
		fmt.Printf("vcpus: %d\n", vcpus)
		fmt.Printf("memory: %d\n", memsize)
		fmt.Println("vmx file: " + vmxpath)
		fmt.Println("vmteam file: " + teampath)
		fmt.Println("guest os: " + guestos)
		//fmt.Println("vm features: " + features)
		if err != nil {
			panic(err)
		}
	}
}
