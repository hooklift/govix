// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vix

import (
	"fmt"
	"strings"

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
type CDDVDDrive struct {
	ID string
	// Either IDE, SCSI or SATA
	Bus BusType
	// Used only when attaching image files. Ex: ISO images
	// If you just want to attach a raw cdrom device leave it empty
	Filename string
}

// Attaches a CD/DVD drive to the virtual machine.
func (v *VM) AttachCDDVD(drive *CDDVDDrive) error {
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
	if drive.Filename != "" {
		device.Filename = drive.Filename
		device.Type = CDROM_IMAGE
	} else {
		device.Type = CDROM_RAW
		device.Autodetect = true
	}

	device.Present = true
	device.StartConnected = true

	if drive.Bus == "" {
		drive.Bus = IDE
	}

	switch drive.Bus {
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
			Text:      fmt.Sprintf("Unrecognized bus type: %s\n", drive.Bus),
		}
	}

	return v.vmxfile.Write()
}

// Detaches a CD/DVD device from the virtual machine
func (v *VM) DetachCDDVD(drive *CDDVDDrive) error {
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

	switch drive.Bus {
	case IDE:
		for i, device := range model.IDEDevices {
			if drive.ID == device.VMXID {
				// This method of removing the element avoids memory leaks
				copy(model.IDEDevices[i:], model.IDEDevices[i+1:])
				model.IDEDevices[len(model.IDEDevices)-1] = vmx.IDEDevice{}
				model.IDEDevices = model.IDEDevices[:len(model.IDEDevices)-1]
			}
		}
	case SCSI:
		for i, device := range model.SCSIDevices {
			if drive.ID == device.VMXID {
				copy(model.SCSIDevices[i:], model.SCSIDevices[i+1:])
				model.SCSIDevices[len(model.SCSIDevices)-1] = vmx.SCSIDevice{}
				model.SCSIDevices = model.SCSIDevices[:len(model.SCSIDevices)-1]
			}
		}
	case SATA:
		for i, device := range model.SATADevices {
			if drive.ID == device.VMXID {
				copy(model.SATADevices[i:], model.SATADevices[i+1:])
				model.SATADevices[len(model.SATADevices)-1] = vmx.SATADevice{}
				model.SATADevices = model.SATADevices[:len(model.SATADevices)-1]
			}
		}
	default:
		return &VixError{
			Operation: "vm.DetachCDDVD",
			Code:      200003,
			Text:      fmt.Sprintf("Unrecognized bus type: %s\n", drive.Bus),
		}
	}

	return v.vmxfile.Write()
}

// Returns an unordered slice of currently attached CD/DVD devices on any bus.
func (v *VM) CDDVDs() ([]*CDDVDDrive, error) {
	// Loads VMX file in memory
	err := v.vmxfile.Read()
	if err != nil {
		return nil, err
	}

	model := v.vmxfile.model
	cddvds := make([]*CDDVDDrive, 0)

	handle := func(d vmx.Device, Bus BusType) {
		if d.Type == CDROM_IMAGE || d.Type == CDROM_RAW {
			cddvds = append(cddvds, &CDDVDDrive{
				ID:       d.VMXID,
				Bus:      Bus,
				Filename: d.Filename,
			})
		}
	}

	for _, d := range model.IDEDevices {
		handle(d.Device, IDE)
	}
	for _, d := range model.SCSIDevices {
		handle(d.Device, SCSI)
	}
	for _, d := range model.SATADevices {
		handle(d.Device, SATA)
	}

	return cddvds, nil
}

func (v *VM) RemoveAllCDDVDDrives() error {
	drives, err := v.CDDVDs()
	if err != nil {
		return &VixError{
			Operation: "vm.RemoveAllCDDVDDrives",
			Code:      200004,
			Text:      fmt.Sprintf("Error listing CD/DVD Drives: %s\n", err),
		}
	}

	for _, d := range drives {
		err := v.DetachCDDVD(d)
		if err != nil {
			return &VixError{
				Operation: "vm.RemoveAllCDDVDDrives",
				Code:      200004,
				Text:      fmt.Sprintf("Error removing CD/DVD Drive %v, error: %s\n", d, err),
			}
		}
	}

	return nil
}

// Returns the CD/DVD drive identified by ID
// This function depends entirely on how GoVMX identifies array's elements
func (v *VM) CDDVD(ID string) (*CDDVDDrive, error) {
	err := v.vmxfile.Read()
	if err != nil {
		return nil, err
	}

	model := v.vmxfile.model
	cddvd := &CDDVDDrive{}

	handle := func(ID string, d vmx.Device, Bus BusType) *CDDVDDrive {
		if ID == d.VMXID {
			cddvd.Bus = Bus
			cddvd.Filename = d.Filename
			return cddvd
		}
		return nil
	}

	if strings.HasPrefix(ID, string(IDE)) {
		for _, d := range model.IDEDevices {
			return handle(ID, d.Device, IDE), nil
		}
	}

	if strings.HasPrefix(ID, string(SCSI)) {
		for _, d := range model.SCSIDevices {
			return handle(ID, d.Device, SCSI), nil
		}
	}

	if strings.HasPrefix(ID, string(SATA)) {
		for _, d := range model.SATADevices {
			return handle(ID, d.Device, SATA), nil
		}
	}

	return nil, nil
}
