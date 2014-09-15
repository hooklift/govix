// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vix

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cloudescape/govmx"
)

// Bus type to use when attaching CD/DVD drives and disks.
type BusType string

// Disk controllers
const (
	IDE  BusType = "ide"
	SCSI BusType = "scsi"
	SATA BusType = "sata"
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
func (v *VM) AttachCDDVD(config *CDDVDConfig) error {
	if running, _ := v.IsRunning(); running {
		return &VixError{
			Operation: "vm.AttachCDDVD",
			Code:      200000,
			Text:      "Virtual machine must be powered off in order to attach a CD/DVD drive.",
		}
	}

	// Loads VMX file in memory
	v.vmxfile.Read()
	model := v.vmxfile.model

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
		model.IDEDevices = append(model.IDEDevices, vmx.IDEDevice{Device: device})
	case SCSI:
		model.SCSIDevices = append(model.SCSIDevices, vmx.SCSIDevice{Device: device})
	case SATA:
		model.SATADevices = append(model.SATADevices, vmx.SATADevice{Device: device})
	default:
		return &VixError{
			Operation: "vm.AttachCDDVD",
			Code:      200001,
			Text:      fmt.Sprintf("Unrecognized bus type: %s\n", config.Bus),
		}
	}

	return v.vmxfile.Write()
}

// Detaches a CD/DVD device from the virtual machine
func (v *VM) DetachCDDVD(config *CDDVDConfig) error {
	if running, _ := v.IsRunning(); running {
		return &VixError{
			Operation: "vm.DetachCDDVD",
			Code:      200002,
			Text:      "Virtual machine must be powered off in order to detach CD/DVD drive.",
		}
	}

	// Loads VMX file in memory
	err := v.vmxfile.Read()
	if err != nil {
		return err
	}

	model := v.vmxfile.model

	switch config.Bus {
	case IDE:
		for i, device := range model.IDEDevices {
			if config.ID == device.VMXID {
				// This method of removing the element avoids memory leaks
				copy(model.IDEDevices[i:], model.IDEDevices[i+1:])
				model.IDEDevices[len(model.IDEDevices)-1] = vmx.IDEDevice{}
				model.IDEDevices = model.IDEDevices[:len(model.IDEDevices)-1]
			}
		}
	case SCSI:
		for i, device := range model.SCSIDevices {
			if config.ID == device.VMXID {
				copy(model.SCSIDevices[i:], model.SCSIDevices[i+1:])
				model.SCSIDevices[len(model.SCSIDevices)-1] = vmx.SCSIDevice{}
				model.SCSIDevices = model.SCSIDevices[:len(model.SCSIDevices)-1]
			}
		}
	case SATA:
		for i, device := range model.SATADevices {
			if config.ID == device.VMXID {
				copy(model.SATADevices[i:], model.SATADevices[i+1:])
				model.SATADevices[len(model.SATADevices)-1] = vmx.SATADevice{}
				model.SATADevices = model.SATADevices[:len(model.SATADevices)-1]
			}
		}
	default:
		return &VixError{
			Operation: "vm.DetachCDDVD",
			Code:      200003,
			Text:      fmt.Sprintf("Unrecognized bus type: %s\n", config.Bus),
		}
	}

	return v.vmxfile.Write()
}

// Returns an unordered slice of currently attached CD/DVD devices on any bus.
func (v *VM) CDDVDs() ([]*CDDVDConfig, error) {
	// Loads VMX file in memory
	err := v.vmxfile.Read()
	if err != nil {
		return nil, err
	}

	model := v.vmxfile.model
	cddvds := make([]*CDDVDConfig, 0)

	var wg sync.WaitGroup
	wg.Add(1)
	// go iterate ide devices
	go func() {
		defer wg.Done()
		for _, device := range model.IDEDevices {
			if device.Type == CDROM_IMAGE || device.Type == CDROM_RAW {
				cddvds = append(cddvds, &CDDVDConfig{
					ID:       device.VMXID,
					Bus:      IDE,
					Filename: device.Filename,
				})
			}
		}
	}()

	// go iterate scsi devices
	go func() {
		defer wg.Done()
		for _, device := range model.SCSIDevices {
			if device.Type == CDROM_IMAGE || device.Type == CDROM_RAW {
				cddvds = append(cddvds, &CDDVDConfig{
					ID:       device.VMXID,
					Bus:      IDE,
					Filename: device.Filename,
				})
			}
		}
	}()

	// go iterate sata devices
	go func() {
		defer wg.Done()
		for _, device := range model.SATADevices {
			if device.Type == CDROM_IMAGE || device.Type == CDROM_RAW {
				cddvds = append(cddvds, &CDDVDConfig{
					ID:       device.VMXID,
					Bus:      IDE,
					Filename: device.Filename,
				})
			}
		}
	}()
	wg.Wait()

	return cddvds, nil
}

// Returns the CD/DVD drive identified by ID
// This function depends entirely on how GoVMX identifies array's elements
func (v *VM) CDDVD(ID string) (*CDDVDConfig, error) {
	err := v.vmxfile.Read()
	if err != nil {
		return nil, err
	}

	model := v.vmxfile.model
	cddvd := &CDDVDConfig{}

	if strings.HasPrefix(ID, string(IDE)) {
		for _, device := range model.IDEDevices {
			if ID == device.VMXID {
				cddvd.Bus = IDE
				cddvd.Filename = device.Filename
				return cddvd, nil
			}
		}
	}

	if strings.HasPrefix(ID, string(SCSI)) {
		for _, device := range model.SCSIDevices {
			if ID == device.VMXID {
				cddvd.Bus = SCSI
				cddvd.Filename = device.Filename
				return cddvd, nil
			}
		}
	}

	if strings.HasPrefix(ID, string(SATA)) {
		for _, device := range model.SATADevices {
			if ID == device.VMXID {
				cddvd.Bus = SATA
				cddvd.Filename = device.Filename
				return cddvd, nil
			}
		}
	}

	return nil, nil
}
