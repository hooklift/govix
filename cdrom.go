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
	// Either IDE, SCSI or SATA
	Bus BusType
	// Used only when attaching image files. Ex: ISO images
	Filename string
}

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
func (v *VM) DetachCdrom(config *CdromConfig) error {
	if running, _ := v.IsRunning(); running {
		return &VixError{
			Operation: "vm.DetachCdrom",
			Code:      200002,
			Text:      "Virtual machine must be powered off in order to detach CD/DVD drive.",
		}
	}
	return nil
}

// Returns the list of currently attached CDROM devices
func (v *VM) Cdrom() (*CdromConfig, error) {
	return nil, nil
}
