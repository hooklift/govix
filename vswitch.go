package vix

import (
	"net"
)

type VSwitch struct {
	id string
	// Provide addresses on this vswitch via DHCP
	DHCP        bool
	DCHPNetwork net.IPNet

	//Connects host machine to this vswitch
	VirtualAdapter bool

	// Allow virtual machines on this vswitch to connect
	// to external networks using NAT
	NAT bool
}

/*
answer VNET_4_DHCP yes
answer VNET_4_VIRTUAL_ADAPTER yes
answer VNET_4_HOSTONLY_NETMASK 255.255.255.0
answer VNET_4_HOSTONLY_SUBNET 192.168.60.0
answer VNET_4_NAT yes
*/

//http://thornelabs.net/2013/10/18/manually-add-and-remove-vmware-fusion-virtual-adapters.html
func AddVSwitch(vswitch VSwitch) (string, error) {
	return "", nil
}

//http://thornelabs.net/2013/10/18/manually-add-and-remove-vmware-fusion-virtual-adapters.html
func RemoveVSwitch(id string) error {
	return nil
}

func ListVSwitches() ([]*VSwitch, error) {
	return nil, nil
}

func ExistVSwitch(id string) bool {
	return false
}

func GetVSwitch(id string) (VSwitch, error) {
	return VSwitch{}, nil
}

// Source http://kb.vmware.com/selfservice/microsites/search.do?language=en_US&cmd=displayKC&externalId=1026510
func restartVMNetServices() {
	//vmnet-cli --configure
	//vmnet-cli --stop
	//vmnet-cli --start
}
