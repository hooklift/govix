// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vix

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cloudescape/govmx"
)

// Device Type
const (
	CDROM_IMAGE string = "cdrom-image"
	CDROM_RAW   string = "cdrom-raw"
)

// CD/DVD configuration
type CDDVDConfig struct {
	ID string
	// Either IDE, SCSI or SATA
	Bus BusType
	// Used only when attaching image files. Ex: ISO images
	// If you just want to attach a raw cdrom device leave it empty
	Filename string
}

// Attaches a CD/DVD drive to the virtual machine.
// TODO(c4milo): make it thread safe
// TODO(c4milo): Avoid unmarshaling in this function, it should be done somewhere
// in host.OpenVM just once.
func (v *VM) AttachCDDVD(config *CDDVDConfig) error {
	if running, _ := v.IsRunning(); running {
		return &VixError{
			Operation: "vm.AttachCDDVD",
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

	device := vmx.Device{}
	if config.Filename != "" {
		device.Filename = config.Filename
		device.Type = CDROM_IMAGE
	} else {
		device.Type = CDROM_RAW
		device.Autodetect = true
	}

	device.Present = true
	device.StartConnected = true

	switch config.Bus {
	case IDE:
		vm.IDEDevices = append(vm.IDEDevices, vmx.IDEDevice{device})
	case SCSI:
		vm.SCSIDevices = append(vm.SCSIDevices, vmx.SCSIDevice{device})
	case SATA:
		vm.SATADevices = append(vm.SATADevices, vmx.SATADevice{device})
	default:
		return &VixError{
			Operation: "vm.AttachCDDVD",
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

// Detaches a CD/DVD device from the virtual machine
// TODO(c4milo): make it thread safe
// TODO(c4milo): Avoid unmarshaling in this function, it should be done somewhere
// in host.OpenVM just once.
func (v *VM) DetachCDDVD(config *CDDVDConfig) error {
	if running, _ := v.IsRunning(); running {
		return &VixError{
			Operation: "vm.DetachCDDVD",
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
			Operation: "vm.DetachCDDVD",
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

// Returns an unordered slice of currently attached CD/DVD devices on any bus.
// TODO(c4milo): Avoid unmarshaling in this function, it should be done somewhere
// in host.OpenVM just once.
func (v *VM) CDDVDs() ([]*CDDVDConfig, error) {
	// unmarshal vmx
	v.vmxfile.Seek(0, 0)
	data, err := ioutil.ReadAll(v.vmxfile)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}

	vm := new(vmx.VirtualMachine)
	if err := vmx.Unmarshal(data, vm); err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err)
	}

	var cddvds []*CDDVDConfig
	vm.WalkDevices(func(d vmx.Device) {
		if d.Type == CDROM_IMAGE || d.Type == CDROM_RAW {
			cddvds = append(cddvds, &CDDVDConfig{
				ID:       d.VMXID,
				Bus:      IDE,
				Filename: d.Filename,
			})
		}
	})
	return cddvds, nil
}

// Returns the CD/DVD drive identified by ID
// This function depends entirely on how GoVMX identifies array's elements
// TODO(c4milo): Avoid unmarshaling in this function, it should be done somewhere
// in host.OpenVM just once.
func (v *VM) CDDVD(ID string) (*CDDVDConfig, error) {
	// unmarshal vmx
	v.vmxfile.Seek(0, 0)
	data, err := ioutil.ReadAll(v.vmxfile)
	if err != nil {
		return nil, err
	}

	vm := new(vmx.VirtualMachine)

	err = vmx.Unmarshal(data, vm)
	if err != nil {
		return nil, err
	}

	var bus string
	switch {
	case strings.HasPrefix(ID, string(vmx.IDE)):
		bus = vmx.IDE
	case strings.HasPrefix(ID, string(vmx.SCSI)):
		bus = vmx.SCSI
	case strings.HasPrefix(ID, string(vmx.SATA)):
		bus = vmx.SATA
	}
	var filename string
	found := vm.FindDevice(func(d vmx.Device) bool {
		if ID == d.VMXID {
			filename = d.Filename
		}
		return ID == d.VMXID
	}, bus)

	if !found {
		return nil, nil
	}
	return &CDDVDConfig{Bus: bus, Filename: filename}, nil
}
