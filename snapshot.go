package vix

/*
#cgo darwin LDFLAGS: -L /Users/camilo/Dropbox/Development/go/src/github.com/c4milo/govix -lvixAllProducts -ldl -lpthread
#include "vix.h"
#include "helper.h"
*/
import "C"

import (
	"runtime"
)

type Snapshot struct {
	// Internal VIX handle
	handle C.VixHandle

	// User defined name for the snapshot.
	Name string

	// User defined description for the snapshot.
	Description string

	// Whether the snapshot is replayable.
	IsReplayable bool
}

// This function returns the specified child snapshot.
//
// Parameters:
//
// index: Index into the list of snapshots.
//
// Remarks:
//
// * Snapshots are indexed from 0 to n-1, where n is the number of child
//   snapshots. Use the function Snapshot.NumChildren() to get the value of n.
// * This function is not supported when using the VMWARE_PLAYER provider.
//
// Since VMware Workstation 6.0
func (s *Snapshot) Child(index int) (*Snapshot, error) {
	var snapshotHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	err = C.VixSnapshot_GetChild(s.handle,
		C.int(index),    //index
		&snapshotHandle) //(output) A handle to the child snapshot.

	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	snapshot := &Snapshot{
		handle: snapshotHandle,
	}

	runtime.SetFinalizer(snapshot, cleanupSnapshot)

	return snapshot, nil
}

// This function returns the number of child snapshots of a specified snapshot.
//
// Remarks:
//
// * This function is not supported when using the VMWARE_PLAYER provider.
//
// Since VMware Workstation 6.0.
func (s *Snapshot) NumChildren() (int, error) {
	var err C.VixError = C.VIX_OK
	var numChildren *C.int

	err = C.VixSnapshot_GetNumChildren(s.handle, numChildren)

	if C.VIX_OK != err {
		return 0, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return int(*numChildren), nil
}

// This function returns the parent of a snapshot.
//
// Remarks:
//
// * This function is not supported when using the VMWARE_PLAYER provider
//
// Since VMware Workstation 6.0
func (s *Snapshot) Parent() (*Snapshot, error) {
	var snapshotHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	err = C.VixSnapshot_GetParent(s.handle,
		&snapshotHandle) //(output) A handle to the child snapshot.

	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	snapshot := &Snapshot{
		handle: snapshotHandle,
	}

	runtime.SetFinalizer(snapshot, cleanupSnapshot)

	return snapshot, nil
}

// Private function to clean up snapshot handle
func cleanupSnapshot(s *Snapshot) {
	if s.handle != C.VIX_INVALID_HANDLE {
		C.Vix_ReleaseHandle(s.handle)
		s.handle = C.VIX_INVALID_HANDLE
	}
}
