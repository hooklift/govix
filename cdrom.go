// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vix

import (
	"fmt"
	"io/ioutil"

	"github.com/cloudescape/govmx"
)

// Bus type to use when attaching CD/DVD drives and disks.
type BusType string

const (
	IDE  BusType = "ide"
	SCSI BusType = "scsi"
	SATA BusType = "sata"
)

// CDROM configuration
type CdromConfig struct {
	ID string
	// Either IDE, SCSI or SATA
	Bus BusType
	// Used only when attaching image files. Ex: ISO images
	// If you just want to attach a raw cdrom device leave it empty
	Filename string
}

// Attaches a CD/DVD drive to the virtual machine.
// TODO(c4milo): make it thread safe
func (v *VM) AttachCdrom(config *CdromConfig) error {
	if running, _ := v.IsRunning(); running {
		return &VixError{
			Operation: "vm.AttachCdrom",
			Code:      200000,
			Text:      "Virtual machine must be powered off in order to attach a CD/DVD drive.",
		}
	}

	v.vmxfile.Seek(0, 0)
	data, err := ioutil.ReadAll(v.vmxfile)
	if err != nil {
		return err
	}

	vm := new(vmx.VirtualMachine)

	err = vmx.Unmarshal(data, vm)
	if err != nil {
		return err
	}

	switch config.Bus {
	case IDE:
		device := vmx.IDEDevice{}
		if config.Filename != "" {
			device.Filename = config.Filename
			device.Type = "cdrom-image"
		} else {
			device.Type = "cdrom-raw"
			device.Autodetect = true
		}

		device.Present = true
		device.StartConnected = true
		vm.IDEDevices = append(vm.IDEDevices, device)
	case SCSI:
		device := vmx.SCSIDevice{}
		if config.Filename != "" {
			device.Filename = config.Filename
			device.Type = "cdrom-image"
		} else {
			device.Type = "cdrom-raw"
			device.Autodetect = true
		}

		device.Present = true
		device.StartConnected = true
		vm.SCSIDevices = append(vm.SCSIDevices, device)
	case SATA:
		device := vmx.SATADevice{}
		if config.Filename != "" {
			device.Filename = config.Filename
			device.Type = "cdrom-image"
		} else {
			device.Type = "cdrom-raw"
			device.Autodetect = true
		}

		device.Present = true
		device.StartConnected = true
		vm.SATADevices = append(vm.SATADevices, device)
	default:
		return &VixError{
			Operation: "vm.AttachCdrom",
			Code:      200001,
			Text:      fmt.Sprintf("Unrecognized bus type: %s\n", config.Bus),
		}
	}

	newdata, err := vmx.Marshal(vm)
	if err != nil {
		return err
	}

	v.vmxfile.Seek(0, 0)
	_, err = v.vmxfile.Write(newdata)

	return err
}

// Detaches a CDROM device from the virtual machine
// TODO(c4milo): make it thread safe
func (v *VM) DetachCdrom(config *CdromConfig) error {
	if running, _ := v.IsRunning(); running {
		return &VixError{
			Operation: "vm.DetachCdrom",
			Code:      200002,
			Text:      "Virtual machine must be powered off in order to detach CD/DVD drive.",
		}
	}

	v.vmxfile.Seek(0, 0)
	data, err := ioutil.ReadAll(v.vmxfile)
	if err != nil {
		return err
	}

	vm := new(vmx.VirtualMachine)

	err = vmx.Unmarshal(data, vm)
	if err != nil {
		return err
	}

	switch config.Bus {
	case IDE:
		for i, device := range vm.IDEDevices {
			if config.ID == device.VMXID {
				// This method of removing the element avoids memory leaks
				copy(vm.IDEDevices[i:], vm.IDEDevices[i+1:])
				vm.IDEDevices[len(vm.IDEDevices)-1] = vmx.IDEDevice{}
				vm.IDEDevices = vm.IDEDevices[:len(vm.IDEDevices)-1]
			}
		}
	case SCSI:
		for i, device := range vm.SCSIDevices {
			if config.ID == device.VMXID {
				copy(vm.SCSIDevices[i:], vm.SCSIDevices[i+1:])
				vm.SCSIDevices[len(vm.SCSIDevices)-1] = vmx.SCSIDevice{}
				vm.SCSIDevices = vm.SCSIDevices[:len(vm.SCSIDevices)-1]
			}
		}
	case SATA:
		for i, device := range vm.SATADevices {
			if config.ID == device.VMXID {
				copy(vm.SATADevices[i:], vm.SATADevices[i+1:])
				vm.SATADevices[len(vm.SATADevices)-1] = vmx.SATADevice{}
				vm.SATADevices = vm.SATADevices[:len(vm.SATADevices)-1]
			}
		}
	default:
		return &VixError{
			Operation: "vm.DetachCdrom",
			Code:      200003,
			Text:      fmt.Sprintf("Unrecognized bus type: %s\n", config.Bus),
		}
	}

	newdata, err := vmx.Marshal(vm)
	if err != nil {
		return err
	}

	v.vmxfile.Seek(0, 0)
	_, err = v.vmxfile.Write(newdata)

	return nil
}

// Returns the list of currently attached CD/DVD devices
func (v *VM) Cdroms() (*CdromConfig, error) {
	return nil, nil
}
