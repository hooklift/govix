package vix

/*
#cgo LDFLAGS: -L . -lvixAllProducts -ldl -lpthread
#include "vix.h"
#include <stdio.h>

VixError getHandle(
	VixHandle jobHandle,
	VixPropertyID prop1,
	VixHandle* handle,
	VixPropertyID prop2) {

	return VixJob_Wait(jobHandle,
		prop1,
		handle,
		prop2);
}

VixError allocVmPasswordPropList(
	VixHandle handle,
	VixHandle* resultHandle,
	char* password
) {
	return VixPropertyList_AllocPropertyList(handle,
                                        resultHandle,
                                        VIX_PROPERTY_VM_ENCRYPTION_PASSWORD,
                                        password,
                                        VIX_PROPERTY_NONE);
}

VixError getScreenshotBytes(
	VixHandle handle,
	int* byte_count,
	char* screen_bits) {

	return VixJob_Wait(handle,
		VIX_PROPERTY_JOB_RESULT_SCREEN_IMAGE_DATA,
		byte_count, screen_bits,
		VIX_PROPERTY_NONE);
}

VixError getNumSharedFolders(VixHandle jobHandle, int* numSharedFolders) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_SHARED_FOLDER_COUNT,
		numSharedFolders,
		VIX_PROPERTY_NONE);
}

VixError readVariable(VixHandle jobHandle, char* readValue) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_VM_VARIABLE_STRING,
		readValue,
		VIX_PROPERTY_NONE);
}

VixError getTempFilePath(VixHandle jobHandle, char* tempFilePath) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_ITEM_NAME,
		tempFilePath,
		VIX_PROPERTY_NONE);
}

VixError isFileOrDir(VixHandle jobHandle, int* result) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_GUEST_OBJECT_EXISTS,
		result,
		VIX_PROPERTY_NONE);
}

VixError runProgramResult(
	VixHandle jobHandle,
	uint64* pid,
	int* elapsedTime,
	int* exitCode) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_PROCESS_ID,
		pid,
		VIX_PROPERTY_JOB_RESULT_GUEST_PROGRAM_ELAPSED_TIME,
		elapsedTime,
		VIX_PROPERTY_JOB_RESULT_GUEST_PROGRAM_EXIT_CODE,
		exitCode,
		VIX_PROPERTY_NONE);
}

VixError getSharedFolder(
	VixHandle jobHandle,
	char* folderName,
	char* folderHostPath,
	int* folderFlags) {

	return VixJob_Wait(jobHandle,
    	VIX_PROPERTY_JOB_RESULT_ITEM_NAME, folderName,
        VIX_PROPERTY_JOB_RESULT_SHARED_FOLDER_HOST, folderHostPath,
        VIX_PROPERTY_JOB_RESULT_SHARED_FOLDER_FLAGS, folderFlags,
		VIX_PROPERTY_NONE);
}


void findItemsResult(VixHandle jobHandle,
                   	VixEventType eventType,
					VixHandle moreEventInfo,
                    void *clientData)
{
   VixError err = VIX_OK;
   char *url = NULL;

   // Check callback event; ignore progress reports.
   if (VIX_EVENTTYPE_FIND_ITEM != eventType) {
      return;
   }

   // Found a virtual machine.
   err = Vix_GetProperties(moreEventInfo,
                           VIX_PROPERTY_FOUND_ITEM_LOCATION,
                           &url,
                           VIX_PROPERTY_NONE);
   if (VIX_OK != err) {
      // Handle the error...
      goto abort;
   }

   printf("\nFound virtual machine: %s", url);

abort:
   Vix_FreeBuffer(url);
}

VixError getFileInfo(VixHandle jobHandle,
					 int64* fsize,
					 int* flags,
					 int64* modtime) {

	return Vix_GetProperties(jobHandle,
		                VIX_PROPERTY_JOB_RESULT_FILE_SIZE,
                        fsize,
                        VIX_PROPERTY_JOB_RESULT_FILE_FLAGS,
                        flags,
                        VIX_PROPERTY_JOB_RESULT_FILE_MOD_TIME,
                        modtime,
                        VIX_PROPERTY_NONE);
}
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// VixPowerState
//
// These are the possible values reported for VIX_PROPERTY_VM_POWER_STATE
// property.
// These values are bitwise flags.
// The actual value returned for may be a bitwise OR of one more of these flags,
// along with other reserved values not documented here.
// They represent runtime information about the state of the virtual machine.
// To test the value of the property, use the Vix.GetProperties() function.
//
// Since VMware Server 1.0.
type VMPowerState int

const (
	// Indicates that VM.PowerOff() has been called, but the operation itself
	// has not completed.
	POWERSTATE_POWERING_OFF VMPowerState = C.VIX_POWERSTATE_POWERING_OFF

	// Indicates that the virtual machine is not running.
	POWERSTATE_POWERED_OFF VMPowerState = C.VIX_POWERSTATE_POWERED_OFF

	// Indicates that VM.PowerOn() has been called, but the operation itself
	// has not completed.
	POWERSTATE_POWERING_ON VMPowerState = C.VIX_POWERSTATE_POWERING_ON

	// Indicates that the virtual machine is running.
	POWERSTATE_POWERED_ON VMPowerState = C.VIX_POWERSTATE_POWERED_ON

	// Indicates that VM.Suspend() has been called, but the operation itself
	// has not completed.
	POWERSTATE_SUSPENDING VMPowerState = C.VIX_POWERSTATE_SUSPENDING

	// Indicates that the virtual machine is suspended. Use VM.PowerOn() to
	// resume the virtual machine.
	POWERSTATE_SUSPENDED VMPowerState = C.VIX_POWERSTATE_SUSPENDED

	// Indicates that the virtual machine is running and the VMware Tools
	// suite is active. See also the VixToolsState property.
	POWERSTATE_TOOLS_RUNNING VMPowerState = C.VIX_POWERSTATE_TOOLS_RUNNING

	// Indicates that VM.Reset() has been called, but the operation itself
	// has not completed.
	POWERSTATE_RESETTING VMPowerState = C.VIX_POWERSTATE_RESETTING

	// Indicates that a virtual machine state change is blocked, waiting for
	// user interaction.
	POWERSTATE_BLOCKED_ON_MSG VMPowerState = C.VIX_POWERSTATE_BLOCKED_ON_MSG
)

// VixFindItemType
//
// These are the types of searches you can do with Host.FindItems().
//
// Since VMware Server 1.0.
type SearchType int

const (
	// Finds all virtual machines currently running on the host.
	FIND_RUNNING_VMS SearchType = C.VIX_FIND_RUNNING_VMS

	// Finds all virtual machines registered on the host.
	// This search applies only to platform products that maintain a virtual
	// machine registry,
	// such as ESX/ESXi and VMware Server, but not Workstation or Player.
	FIND_REGISTERED_VMS SearchType = C.VIX_FIND_REGISTERED_VMS
)

// VixToolsState
//
// These are the possible values reported for VIX_PROPERTY_VM_TOOLS_STATE.
// They represent runtime information about the VMware Tools suite in the guest
// operating system.
// To test the value of the property, use the Vix.GetProperties() function.
//
// Since VMware Server 1.0.
type GuestToolsState int

const (
	// Indicates that Vix is unable to determine the VMware Tools status.
	TOOLSSTATE_UNKNOWN GuestToolsState = C.VIX_TOOLSSTATE_UNKNOWN

	// Indicates that VMware Tools is running in the guest operating system.
	TOOLSSTATE_RUNNING GuestToolsState = C.VIX_TOOLSSTATE_RUNNING

	// Indicates that VMware Tools is not installed in the guest operating system.
	TOOLSSTATE_NOT_INSTALLED GuestToolsState = C.VIX_TOOLSSTATE_NOT_INSTALLED
)

// Service Provider
type Provider int

const (
	// vCenter Server, ESX/ESXi hosts, and VMware Server 2.0
	VMWARE_VI_SERVER Provider = C.VIX_SERVICEPROVIDER_VMWARE_VI_SERVER

	// VMware Workstation
	VMWARE_WORKSTATION Provider = C.VIX_SERVICEPROVIDER_VMWARE_WORKSTATION

	// VMware Workstation (shared mode)
	VMWARE_WORKSTATION_SHARED Provider = C.VIX_SERVICEPROVIDER_VMWARE_WORKSTATION_SHARED

	// With VMware Player
	VMWARE_PLAYER Provider = C.VIX_SERVICEPROVIDER_VMWARE_PLAYER

	// VMware Server 1.0.x
	VMWARE_SERVER Provider = C.VIX_SERVICEPROVIDER_VMWARE_SERVER
)

type EventType int

const (
	JOB_COMPLETED EventType = C.VIX_EVENTTYPE_JOB_COMPLETED
	JOB_PROGRESS  EventType = C.VIX_EVENTTYPE_JOB_PROGRESS
	FIND_ITEM     EventType = C.VIX_EVENTTYPE_FIND_ITEM
)

type HostOption int

const (
	VERIFY_SSL_CERT = C.VIX_HOSTOPTION_VERIFY_SSL_CERT
)

type GuestLoginOption int

const (
	LOGIN_IN_GUEST_NONE                            GuestLoginOption = 0x0
	LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT GuestLoginOption = C.VIX_LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT
)

type GuestVarType int

const (
	VM_GUEST_VARIABLE          GuestVarType = C.VIX_VM_GUEST_VARIABLE
	VM_CONFIG_RUNTIME_ONLY     GuestVarType = C.VIX_VM_CONFIG_RUNTIME_ONLY
	GUEST_ENVIRONMENT_VARIABLE GuestVarType = C.VIX_GUEST_ENVIRONMENT_VARIABLE
)

type SharedFolderOption int

const (
	SHAREDFOLDER_WRITE_ACCESS = C.VIX_SHAREDFOLDER_WRITE_ACCESS
)

type CloneType int

const (
	CLONETYPE_FULL   CloneType = C.VIX_CLONETYPE_FULL
	CLONETYPE_LINKED CloneType = C.VIX_CLONETYPE_LINKED
)

type CreateSnapshotOption int

const (
	SNAPSHOT_INCLUDE_MEMORY CreateSnapshotOption = C.VIX_SNAPSHOT_INCLUDE_MEMORY
)

type VmDeleteOption int

const (
	VMDELETE_DISK_FILES VmDeleteOption = C.VIX_VMDELETE_DISK_FILES
)

type VMPowerOption int

const (
	VMPOWEROP_NORMAL                    VMPowerOption = C.VIX_VMPOWEROP_NORMAL
	VMPOWEROP_FROM_GUEST                VMPowerOption = C.VIX_VMPOWEROP_FROM_GUEST
	VMPOWEROP_SUPPRESS_SNAPSHOT_POWERON VMPowerOption = C.VIX_VMPOWEROP_SUPPRESS_SNAPSHOT_POWERON
	VMPOWEROP_LAUNCH_GUI                VMPowerOption = C.VIX_VMPOWEROP_LAUNCH_GUI
	VMPOWEROP_START_VM_PAUSED           VMPowerOption = C.VIX_VMPOWEROP_START_VM_PAUSED
)

type RunProgramOption int

const (
	RUNPROGRAM_WAIT               RunProgramOption = 0x0
	RUNPROGRAM_RETURN_IMMEDIATELY RunProgramOption = C.VIX_RUNPROGRAM_RETURN_IMMEDIATELY
	RUNPROGRAM_ACTIVATE_WINDOW    RunProgramOption = C.VIX_RUNPROGRAM_ACTIVATE_WINDOW
)

type InstallToolsOption int

const (
	INSTALLTOOLS_MOUNT_TOOLS_INSTALLER InstallToolsOption = C.VIX_INSTALLTOOLS_MOUNT_TOOLS_INSTALLER
	INSTALLTOOLS_AUTO_UPGRADE          InstallToolsOption = C.VIX_INSTALLTOOLS_AUTO_UPGRADE
	INSTALLTOOLS_RETURN_IMMEDIATELY    InstallToolsOption = C.VIX_INSTALLTOOLS_RETURN_IMMEDIATELY
)

type RemoveSnapshotOption int

const (
	SNAPSHOT_REMOVE_NONE     RemoveSnapshotOption = 0x0
	SNAPSHOT_REMOVE_CHILDREN RemoveSnapshotOption = C.VIX_SNAPSHOT_REMOVE_CHILDREN
)

type FileAttr int

const (
	FILE_ATTRIBUTES_DIRECTORY FileAttr = C.VIX_FILE_ATTRIBUTES_DIRECTORY
	FILE_ATTRIBUTES_SYMLINK   FileAttr = C.VIX_FILE_ATTRIBUTES_SYMLINK
)

type Host struct {
	handle C.VixHandle
}

// Connects to a Provider
//
// Parameters:
//
// Provider:
// * With vCenter Server, ESX/ESXi hosts, and VMware Server 2.0,
//   VMWARE_VI_SERVER.
// * With VMware Workstation, use VMWARE_WORKSTATION.
// * With VMware Workstation (shared mode), use VMWARE_WORKSTATION_SHARED.
// * With VMware Player, use VMWARE_PLAYER.
// * With VMware Server 1.0.x, use VMWARE_SERVER.
//
// Hostname:
// Varies by product platform. With vCenter Server, ESX/ESXi hosts,
// VMware Workstation (shared mode) and VMware Server 2.0,
// use a URL of the form "https://<hostName>:<port>/sdk"
// where <hostName> is either the DNS name or IP address.
// If missing, <port> may default to 443 (see Remarks below).
// In VIX API 1.10 and later, you can omit "https://" and "/sdk" specifying
// just the DNS name or IP address.
// Credentials are required even for connections made locally.
// With Workstation, use nil to connect to the local host.
// With VMware Server 1.0.x, use the DNS name or IP address for remote
// connections, or the same as Workstation for local connections.
//
// Port:
// TCP/IP port on the remote host.
// With VMware Workstation and VMware Player, use zero for the local host.
// With ESX/ESXi hosts, VMware Workstation (shared mode) and VMware Server 2.0
// you specify port number within the hostName parameter, so this parameter is
// ignored (see Remarks below).
//
// Username:
// Username for authentication on the remote machine.
// With VMware Workstation, VMware Player, and VMware Server 1.0.x,
// use nil to authenticate as the current user on local host.
// With vCenter Server, ESX/ESXi hosts, VMware Workstation (shared mode)
// and VMware Server 2.0, you must use a valid login.
//
// Password:
// Password for authentication on the remote machine.
// With VMware Workstation, VMware Player, and VMware Server 1.0.x,
// use nil to authenticate as the current user on local host.
// With ESX/ESXi, VMware Workstation (shared mode) and VMware Server 2.0, you
// must use a valid login.
//
// Remarks:
// * To specify the local host (where the API client runs) with VMware
//   Workstation and VMware Player, pass nil values for the hostname, port,
//   login, and password parameters.
// * With vCenter Server, ESX/ESXi hosts, and VMware Server 2.0, the URL for
//   the hostname argument may specify the port.
//   Otherwise an HTTPS connection is attempted on port 443. HTTPS is strongly
//   recommended.
//   Port numbers are set during installation of Server 2.0. The installer's
//   default HTTP and HTTPS values are 8222 and 8333 for Server on Windows, or
//   (if not already in use) 80 and 443 for Server on Linux, and 902 for the
//   automation socket, authd. If connecting to a virtual machine though a
//   firewall, port 902 and the communicating port must be opened to allow
//   guest operations.
// * If a VMware ESX host is being managed by a VMware VCenter Server, you
//   should call VixHost_Connect with the hostname or IP address of the VCenter
//   server, not the ESX host.
//   Connecting directly to an ESX host while bypassing its VCenter Server can
//   cause state inconsistency.
// * On Windows, this function should not be called multiple times with
//   different service providers in the same process; doing so will result in
//   a VIX_E_WRAPPER_MULTIPLE_SERVICEPROVIDERS error.
//   A single client process can connect to multiple hosts as long as it
//   connects using the same service provider type.
// * To enable SSL certificate verification, set the value of the options
//   parameter to include the bit flag specified by VERIFY_SSL_CERT.
//   This option can also be set in the VMware config file by assigning
//   vix.enableSslCertificateCheck as TRUE or FALSE.
//   The vix.sslCertificateFile config option specifies the path to a file
//   containing CA certificates in PEM format.
//   The vix.sslCertificateDirectory config option can specify a directory
//   containing files that each contain a CA certificate.
//   Upon encountering a SSL validation error, the host handle is not created
//   with a resulting error code of E_NET_HTTP_SSL_SECURITY.
// * With VMware vCenter Server and ESX/ESXi 4.0 hosts, an existing VI API
//   session can be used instead of the username/password pair to authenticate
//   when connecting. To use an existing VI API session, a VI "clone ticket"
//   is required; call the VI API AcquireCloneTicket() method of the
//   SessionManager object to get this ticket.
//   Using the ticket string returned by this method, call VixHost_Connect()
//   with NULL as the 'username' and the ticket as the 'password'.
//
// Since VMware Server 1.0
func Connect(
	hostname string,
	port uint,
	username, password string,
	provider, options int,
) (*Host, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var hostHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixHost_Connect(C.VIX_API_VERSION,
		C.VixServiceProvider(provider),
		C.CString(hostname),
		C.int(port),
		C.CString(username),
		C.CString(password),
		C.VixHostOptions(options),
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	err = C.getHandle(jobHandle,
		C.VIX_PROPERTY_JOB_RESULT_HANDLE,
		&hostHandle,
		C.VIX_PROPERTY_NONE)

	defer C.Vix_ReleaseHandle(jobHandle)

	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	host := &Host{
		handle: hostHandle,
	}

	runtime.SetFinalizer(host, cleanupHost)

	return host, nil
}

// Private function to clean up host handle
func cleanupHost(host *Host) {
	if host.handle != C.VIX_INVALID_HANDLE {
		host.Disconnect()
	}
}

// Private function to clean up vm handle
func cleanupVM(v *VM) {
	if v.handle != C.VIX_INVALID_HANDLE {
		C.Vix_ReleaseHandle(v.handle)
		v.handle = C.VIX_INVALID_HANDLE
	}
}

// Private function to clean up snapshot handle
func cleanupSnapshot(s *Snapshot) {
	if s.handle != C.VIX_INVALID_HANDLE {
		C.Vix_ReleaseHandle(s.handle)
		s.handle = C.VIX_INVALID_HANDLE
	}
}

// Destroys the state for a particular host instance
//
// Call this function to disconnect the host.
// After you call this function the Host object is no longer valid
// and you should not longer use it.
// Similarly, you should not use any other object instances
// obtained from the Host object while it was connected.
//
// Since VMware Server 1.0
func (h *Host) Disconnect() {
	C.VixHost_Disconnect(h.handle)
	h.handle = C.VIX_INVALID_HANDLE
}

// This function finds Vix objects. For example, when used to find all
// running virtual machines, Host.FindItems() returns a series of virtual
// machine file path names.
func (h *Host) FindItems(options SearchType) ([]string, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixHost_FindItems(h.handle,
		C.VixFindItemType(options), //searchType
		C.VIX_INVALID_HANDLE,       //searchCriteria
		-1,                         //timeout
		nil,                        //callbackProc
		nil)                        //clientData

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil, nil
}

// This function opens a virtual machine on the host
// and returns a VM instance.
//
// Parameters:
//
// VmxFile:
// The path name of the virtual machine configuration file on the local host.
//
// Password:
// If VM is encrypted, this is the password for VIX to be able to open it.
//
// Remarks:
//
// * This function opens a virtual machine on the host instance
//   The virtual machine is identified by vmxFile, which is a path name to the
//   configuration file (.VMX file) for that virtual machine.
// * The format of the path name depends on the host operating system.
//   For example, a path name for a Windows host requires backslash as a
//   directory separator, whereas a Linux host requires a forward slash. If the
//   path name includes backslash characters, you need to precede each one with
//   an escape character. For VMware Server 2.x, the path contains a preceeding
//   data store, for example [storage1] vm/vm.vmx.
// * For VMware Server hosts, a virtual machine must be registered before you
//   can open it. You can register a virtual machine by opening it with the
//   VMware Server Console, through the vmware-cmd command with the register
//   parameter, or with Host.RegisterVM().
// * For vSphere, the virtual machine opened may not be the one desired if more
//   than one Datacenter contains VmxFile.
// * To open an encrypted virtual machine, pass its correspondent password.
//
// Since VMware Workstation 7.0
func (h *Host) OpenVm(vmxFile, password string) (*VM, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var propertyHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var vmHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	defer C.Vix_ReleaseHandle(propertyHandle)
	defer C.Vix_ReleaseHandle(jobHandle)

	if password != "" {
		err = C.allocVmPasswordPropList(h.handle,
			&propertyHandle,
			C.CString(password))

		if C.VIX_OK != err {
			return nil, &VixError{
				code: int(err & 0xFFFF),
				text: C.GoString(C.Vix_GetErrorText(err, nil)),
			}
		}
	}

	jobHandle = C.VixHost_OpenVM(h.handle,
		C.CString(vmxFile),
		C.VIX_VMOPEN_NORMAL,
		propertyHandle,
		nil, // callbackProc
		nil) // clientData

	err = C.getHandle(jobHandle,
		C.VIX_PROPERTY_JOB_RESULT_HANDLE,
		&vmHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	vm := &VM{
		handle: vmHandle,
	}

	runtime.SetFinalizer(vm, cleanupVM)

	return vm, nil
}

// This function adds a virtual machine to the host's inventory.
//
// Parameters:
//
// VmxFile:
// The path name of the .vmx file on the host.
//
// Remarks:
//
// * This function registers the virtual machine identified by vmxFile, which
//   is a storage path to the configuration file (.vmx) for that virtual machine.
//   You can register a virtual machine regardless of its power state.
// * The format of the path name depends on the host operating system.
//   If the path name includes backslash characters, you need to precede each
//   one with an escape character. Path to storage [standard] or [storage1] may
//   vary.
// * For VMware Server 1.x, supply the full path name instead of storage path,
//   and specify provider VMWARE_SERVER to connect.
// * This function has no effect on Workstation or Player, which lack a virtual
//   machine inventory.
// * It is not a Vix error to register an already-registered virtual machine,
//   although the VMware Server UI shows an error icon in the Task pane.
//   Trying to register a non-existent virtual machine results in error 2000,
//   VIX_E_NOT_FOUND.
//
// Since VMware Server 1.0
func (h *Host) RegisterVm(vmxFile string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixHost_RegisterVM(h.handle,
		C.CString(vmxFile),
		nil, // callbackProc
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function removes a virtual machine from the host's inventory.
//
// Parameters:
//
// VmxFile:
// The path name of the .vmx file on the host.
//
// Remarks:
//
// * This function unregisters the virtual machine identified by vmxFile,
//   which is a storage path to the configuration file (.vmx) for that virtual
//   machine. A virtual machine must be powered off to unregister it.
// * The format of the storage path depends on the host operating system.
//   If the storage path includes backslash characters, you need to precede each
//   one with an escape character. Path to storage [standard] or [storage1] may
//   vary.
// * For VMware Server 1.x, supply the full path name instead of storage path,
//   and specify VMWARE_SERVER provider to connect.
// * This function has no effect on Workstation or Player, which lack a virtual
//   machine inventory.
// * It is not a Vix error to unregister an already-unregistered virtual machine,
//   nor is it a Vix error to unregister a non-existent virtual machine.
//
// Since VMware Server 1.0
func (h *Host) UnregisterVm(vmxFile string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixHost_UnregisterVM(h.handle,
		C.CString(vmxFile),
		nil, // callbackProc
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// Copies a file or directory from the local system (where the Vix client is
// running) to the guest operating system.
//
// Parameters:
//
// vm: A VM instance.
// src: The path name of a file on a file system available to the Vix client.
// dest: The path name of a file on a file system available to the guest.
//
// Remarks:
//
// * The virtual machine must be running while the file is copied from the Vix
//   client machine to the guest operating system.
// * Existing files of the same name are overwritten, and folder contents are
//   merged.
// * The copy operation requires VMware Tools to be installed and running in
//   the guest operating system.
// * You must call VM.LoginInGuest() before calling this function in order
//   to get a Guest instance.
// * The format of the file name depends on the guest or local operating system.
//   For example, a path name for a Microsoft Windows guest or host requires
//   backslash as a directory separator, whereas a Linux guest or host requires
//   a forward slash. If the path name includes backslash characters,
//   you need to precede each one with an escape character.
// * Only absolute paths should be used for files in the guest; the resolution
//   of relative paths is not specified.
// * If any file fails to be copied, Vix aborts the operation, does not attempt
//   to copy the remaining files, and returns an error.
// * In order to copy a file to a mapped network drive in a Windows guest
//   operating system,
//   it is necessary to call VixVM_LoginInGuest() with the
//   LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT flag set.
//   Using the interactive session option incurs an overhead in file transfer
//   speed.
//
//  Since VMware Server 1.0
//  Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
//
func (h *Host) CopyFileToGuest(src string, guest *Guest, dest string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_CopyFileFromHostToGuest(
		guest.handle,         //VM handle
		C.CString(src),       // src name
		C.CString(dest),      // dest name
		C.int(0),             // options
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

type VM struct {
	// Internal VIX handle
	handle C.VixHandle

	// The number of virtual CPUs configured for the virtual machine.
	VCpus uint8

	// The path to the virtual machine configuration file.
	VmxPath string

	// The path to the virtual machine team.
	VmTeamPath string

	// The memory size of the virtual machine.
	MemorySize uint

	ReadOnly bool

	// Whether the virtual machine is a member of a team.
	InVMTeam bool

	// The power state of the virtual machine.
	PowerState VMPowerState

	// The state of the VMware Tools suite in the guest.
	ToolState GuestToolsState

	// Whether the virtual machine is running.
	IsRunning bool
}

// This function enables or disables all shared folders as a feature for a
// virtual machine.
//
// Remarks:
//
// * This function enables/disables all shared folders as a feature on a
//   virtual machine.
//   In order to access shared folders on a guest, the feature has to be enabled,
//   and in addition, the individual shared folder has to be enabled.
// * It is not necessary to call VM.LoginInGuest() before calling this function.
// * In this release, this function requires the virtual machine to be powered
//   on with VMware Tools installed.
// * Shared folders are not supported for the following guest operating systems:
//   Windows ME, Windows 98, Windows 95, Windows 3.x, and DOS.
// * On Linux virtual machines, calling this function will automatically mount
//   shared folder(s) in the guest.
//
// Since VMware Workstation 6.0, not available on Server 2.0.
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
//
func (v *VM) EnableSharedFolders(enabled bool) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	var share C.Bool = C.FALSE
	if enabled {
		share = C.TRUE
	} else {
		share = C.FALSE
	}

	jobHandle = C.VixVM_EnableSharedFolders(v.handle,
		share,
		0,
		nil,
		nil)

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function mounts a new shared folder in the virtual machine.
//
// Parameters:
//
// guestpath: Specifies the guest path name of the new shared folder.
// hostpath: Specifies the host path of the shared folder.
// flags: The folder options.
//
// Remarks:
//
// * This function creates a local mount point in the guest file system and
//   mounts a shared folder exported by the host.
// * Shared folders will only be accessible inside the guest operating system
//   if shared folders are enabled for the virtual machine.
//   See the documentation for VM.EnableSharedFolders().
// * The folder options include:
// 	 SHAREDFOLDER_WRITE_ACCESS - Allow write access.
// * Only absolute paths should be used for files in the guest; the resolution
//   of relative paths is not specified.
// * The hostpath argument must specify a path to a directory that exists on the
//   host, or an error will result.
// * If a shared folder with the same name exists before calling this function,
//   the job handle returned by this function will return VIX_E_ALREADY_EXISTS.
// * It is not necessary to call VM.LoginInGuest() before calling this function.
// * When creating shared folders in a Windows guest, there might be a delay
//   before contents of a shared folder are visible to functions such as
//   Guest.IsFile() and Guest.RunProgram().
// * Shared folders are not supported for the following guest operating
//   systems: Windows ME, Windows 98, Windows 95, Windows 3.x, and DOS.
// * In this release, this function requires the virtual machine to be powered
//   on with VMware Tools installed.
// * To determine in which directory in the guest the shared folder will be,
//   query Guest.SharedFoldersParentDir(). When the virtual machine is powered
//   on and the VMware Tools are running, this property will contain the path to
//   the parent directory of the shared folders for that virtual machine.
//
// Since VMware Workstation 6.0, not available on Server 2.0.
func (v *VM) AddSharedFolder(guestpath, hostpath string, flags SharedFolderOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_AddSharedFolder(v.handle,
		C.CString(guestpath),
		C.CString(hostpath),
		C.VixMsgSharedFolderOptions(flags),
		nil, nil)

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function removes a shared folder in the virtual machine.
//
// Parameters:
//
// * guestpath: Specifies the guest pathname of the shared folder to delete.
//
// Remarks:
//
// * This function removes a shared folder in the virtual machine referenced by
//   the VM object
// * It is not necessary to call VM.LoginInGuest() before calling this function.
// * Shared folders are not supported for the following guest operating
//   systems: Windows ME, Windows 98, Windows 95, Windows 3.x, and DOS.
// * In this release, this function requires the virtual machine to be powered
//   on with VMware Tools installed.
// * Depending on the behavior of the guest operating system, when removing
//   shared folders, there might be a delay before the shared folder is no
//   longer visible to programs running within the guest operating system and
//   to functions such as Guest.IsFile()
//
// Since VMware Workstation 6.0

func (v *VM) RemoveSharedFolder(guestpath string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_RemoveSharedFolder(v.handle,
		C.CString(guestpath),
		0,
		nil, nil)

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function captures the screen of the guest operating system.
//
// Remarks:
//
// * This function captures the current screen image and returns it as a
//   []byte result.
// * For security reasons, this function requires a successful call to
//   VM.LoginInGuest() must be made.
//
// Since VMware Workstation 6.5
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (v *VM) Screenshot() ([]byte, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var byte_count C.int
	var screen_bits C.char

	jobHandle = C.VixVM_CaptureScreenImage(v.handle,
		C.VIX_CAPTURESCREENFORMAT_PNG,
		C.VIX_INVALID_HANDLE,
		nil,
		nil)
	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.getScreenshotBytes(jobHandle, &byte_count, &screen_bits)
	defer C.Vix_FreeBuffer(unsafe.Pointer(&screen_bits))

	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return C.GoBytes(unsafe.Pointer(&screen_bits), byte_count), nil
}

// Creates a copy of the virtual machine specified by the current VM instance
//
// Parameters:
//
// cloneType:
// Must be either CLONETYPE_FULL or CLONETYPE_LINKED.
// * CLONETYPE_FULL - Creates a full, independent clone of the virtual machine.
// * CLONETYPE_LINKED - Creates a linked clone, which is a copy of a virtual
//                      machine that shares virtual disks with the parent
//                      virtual machine in an ongoing manner.
//                      This conserves disk space as long as the parent and
//                      clone do not change too much from their original state.
//
// destVmxFile:
// The path name of the virtual machine configuration file that will
// be created for the virtual machine clone produced by this operation.
// This should be a full absolute path name, with directory names delineated
// according to host system convention: \ for Windows and / for Linux.
//
// Remarks:
//
// * The function returns a new VM instance which is a clone of its parent VM.
// * It is not possible to create a full clone of a powered on virtual machine.
//   You must power off or suspend a virtual machine before creating a full
//   clone of that machine.
// * With a suspended virtual machine, requesting a linked clone results in
//   error 3007 VIX_E_VM_IS_RUNNING.
//   Suspended virtual machines retain memory state, so proceeding with a
//   linked clone could cause loss of data.
// * A linked clone must have access to the parent's virtual disks. Without
//   such access, you cannot use a linked clone
//   at all because its file system will likely be incomplete or corrupt.
// * Deleting a virtual machine that is the parent of a linked clone renders
//   the linked clone useless.
// * Because a full clone does not share virtual disks with the parent virtual
//   machine, full clones generally perform better than linked clones.
//   However, full clones take longer to create than linked clones. Creating a
//   full clone can take several minutes if the files involved are large.
// * This function is not supported when using the VMWARE_PLAYER provider.
func (v *VM) Clone(cloneType CloneType, destVmxFile string) (*VM, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var clonedHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_Clone(v.handle,
		C.VIX_INVALID_HANDLE,      // snapshotHandle
		C.VixCloneType(cloneType), // cloneType
		C.CString(destVmxFile),    // destConfigPathName
		0,                    //options,
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.getHandle(jobHandle,
		C.VIX_PROPERTY_JOB_RESULT_HANDLE,
		&clonedHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	vm := &VM{
		handle: clonedHandle,
	}

	runtime.SetFinalizer(vm, cleanupVM)

	return vm, nil
}

// This function saves a copy of the virtual machine state as a snapshot object.
//
// Parameters:
//
// name:
// A user-defined name for the snapshot; need not be unique.
//
// description:
// A user-defined description for the snapshot.
//
// options:
// Flags to specify how the snapshot should be created. Any combination of the
// following or 0 to exclude memory:
//    * SNAPSHOT_INCLUDE_MEMORY - Captures the full state of a running virtual
//      machine, including the memory.
//
// Remarks:
//
// * This function creates a child snapshot of the current snapshot.
// * If a virtual machine is suspended, you cannot snapshot it more than once.
// * If a powered-on virtual machine gets a snapshot created with option 0
//   (exclude memory), the power state is not saved, so reverting to the
//   snapshot sets powered-off state.
// * The 'name' and 'description' parameters can be set but not retrieved
//   using the VIX API.
// * VMware Server supports only a single snapshot for each virtual machine.
//   The following considerations apply to VMware Server:
//    * If you call this function a second time for the same virtual machine
//      without first deleting the snapshot,
//      the second call will overwrite the previous snapshot.
//    * A virtual machine imported to VMware Server from another VMware product
//      might have more than one snapshot at the time it is imported. In that
//      case, you can use this function to add a new snapshot to the series.
// * Starting in VMware Workstation 6.5, snapshot operations are allowed on
//   virtual machines that are part of a team.
//   Previously, this operation failed with error code
//   VIX_PROPERTY_VM_IN_VMTEAM. Team members snapshot independently so they can
//   have different and inconsistent snapshot states.
// * This function is not supported when using the VMWARE_PLAYER provider.
// * If the virtual machine is open and powered off in the UI, this function now
//   closes the virtual machine in the UI before creating the snapshot.
//
// Since VMware Workstation 6.0
func (v *VM) CreateSnapshot(name, description string, options CreateSnapshotOption) (*Snapshot, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var snapshotHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_CreateSnapshot(v.handle,
		C.CString(name),                     // name
		C.CString(description),              // description
		C.VixCreateSnapshotOptions(options), // options
		C.VIX_INVALID_HANDLE,                // propertyListHandle
		nil,                                 // callbackProc
		nil)                                 // clientData

	err = C.getHandle(jobHandle,
		C.VIX_PROPERTY_JOB_RESULT_HANDLE,
		&snapshotHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	snapshot := &Snapshot{
		handle:      snapshotHandle,
		Name:        name,
		Description: description,
	}

	runtime.SetFinalizer(snapshot, cleanupSnapshot)

	return snapshot, nil
}

// This function permanently deletes a virtual machine from your host system.
//
// Parameters:
//
// options:
// For VMware Server 2.0 and ESX, this value must be VMDELETE_DISK_FILES.
// For all other versions it can be either 0 or VMDELETE_DISK_FILES.
// When option is VIX_VMDELETE_DISK_FILES, deletes all associated files.
// When option is 0, does not delete *.vmdk virtual disk file(s).
//
// Remarks:
//
// * This function permanently deletes a virtual machine from your host system.
//    You can accomplish the same effect by deleting all virtual machine files
//    using the host file system. This function simplifies the task by deleting
//    all VMware files in a single function call.
//    If a deleteOptions value of 0 is specified, the virtual disk (vmdk) files
//    will not be deleted.
//    This function does not delete other user files in the virtual machine
//    folder.
// * This function is successful only if the virtual machine is powered off or
//   suspended.
// * Deleting a virtual machine that is the parent of a linked clone renders
//   the linked clone useless.
//
// since VMware Server 1.0
func (v *VM) Delete(options VmDeleteOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_Delete(v.handle, C.VixVMDeleteOptions(options), nil, nil)

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	C.Vix_ReleaseHandle(v.handle)
	return nil
}

// This function returns the handle of the current active snapshot belonging to
// the virtual machine
//
// Remarks:
//
// * This function returns the handle of the current active snapshot belonging
//   to the virtual machine.
// * This function is not supported when using the VMWARE_PLAYER provider
//
// Since VMware Workstation 6.0
func (v *VM) CurrentSnapshot() (*Snapshot, error) {
	var snapshotHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	err = C.VixVM_GetCurrentSnapshot(v.handle, &snapshotHandle)
	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	snapshot := &Snapshot{handle: snapshotHandle}

	runtime.SetFinalizer(snapshot, cleanupSnapshot)

	return snapshot, nil
}

// This function returns a Snapshot object matching the given name
//
// Parameters:
//
// Name:
// Identifies a snapshot name.
//
// Remarks:
//
// * When the snapshot name is a duplicate, it returns error 13017
//   VIX_E_SNAPSHOT_NONUNIQUE_NAME.
// * When there are multiple snapshots with the same name, or the same path to
//   that name, you cannot specify a unique name, but you can to use the UI to
//   rename duplicates.
// * You can specify the snapshot name as a path using '/' or '\\' as path
//   separators, including snapshots in the tree above the named snapshot,
//   for example 'a/b/c' or 'x/x'. Do not mix '/' and '\\' in the same path
//   expression.
// * This function is not supported when using the VMWARE_PLAYER provider.
//
// Since VMware Workstation 6.0
func (v *VM) SnapshotByName(name string) (*Snapshot, error) {
	var snapshotHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	err = C.VixVM_GetNamedSnapshot(v.handle, C.CString(name), &snapshotHandle)
	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	snapshot := &Snapshot{handle: snapshotHandle}

	runtime.SetFinalizer(snapshot, cleanupSnapshot)

	return snapshot, nil
}

// This function returns the number of top-level (root) snapshots belonging to a
// virtual machine.
//
// Remarks:
//
// * This function returns the number of top-level (root) snapshots belonging to
//   a virtual machine.
//   A top-level snapshot is one that is not based on any previous snapshot.
//   If the virtual machine has more than one snapshot, the snapshots can be a
//   sequence in which each snapshot is based on the previous one, leaving only
//   a single top-level snapshot.
//   However, if applications create branched trees of snapshots, a single
//   virtual machine can have several top-level snapshots.
// * VMware Server can manage only a single snapshot for each virtual machine.
//   All other snapshots in a sequence are ignored. The return value is always
//   0 or 1.
// * This function is not supported when using the VMWARE_PLAYER provider
//
// Since VMware Workstation 6.0
func (v *VM) TotalRootSnapshots() (int, error) {
	var result C.int
	var err C.VixError = C.VIX_OK

	err = C.VixVM_GetNumRootSnapshots(v.handle, &result)
	if C.VIX_OK != err {
		return 0, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return int(result), nil
}

// This function returns the number of shared folders mounted in the virtual
// machine.
//
// Remarks:
//
// * This function returns the number of shared folders mounted in the virtual
//   machine.
// * It is not necessary to call VM.LoginInGuest() before calling this function.
// * Shared folders are not supported for the following guest operating systems:
//   Windows ME, Windows 98, Windows 95, Windows 3.x, and DOS.
// * In this release, this function requires the virtual machine to be powered
//   on with VMware Tools installed.
//
// Since VMware Workstation 6.0
func (v *VM) TotalSharedFolders() (int, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var numSharedFolders C.int = 0

	jobHandle = C.VixVM_GetNumSharedFolders(v.handle, nil, nil)
	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.getNumSharedFolders(jobHandle, &numSharedFolders)

	if C.VIX_OK != err {
		return 0, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return int(numSharedFolders), nil
}

// This function returns a root Snapshot instance
// belonging to the current virtual machine
//
// Parameters:
//
// Index:
// Identifies a root snapshot. See below for range of values.
//
// Remarks:
//
// * Snapshots are indexed from 0 to n-1, where n is the number of root
//   snapshots. Use the function VM.TotalRootSnapshots to get the value of n.
// * VMware Server can manage only a single snapshot for each virtual machine.
//   The value of index can only be 0.
// * This function is not supported when using the VMWARE_PLAYER provider
//
// Since VMware Server 1.0
func (v *VM) RootSnapshot(index int) (*Snapshot, error) {
	var snapshotHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	err = C.VixVM_GetRootSnapshot(v.handle,
		C.int(index),
		&snapshotHandle)

	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	snapshot := &Snapshot{handle: snapshotHandle}

	runtime.SetFinalizer(snapshot, cleanupSnapshot)

	return snapshot, nil
}

// This function modifies the state of a shared folder mounted in the virtual
// machine.
//
// Parameters:
//
// name: Specifies the name of the shared folder.
// hostpath: Specifies the host path of the shared folder.
// options: The new flag settings.
//
// Remarks:
//
// * This function modifies the state flags of an existing shared folder.
// * If the shared folder does not exist before calling
//   this function, the function will return a not found error.
// * Shared folders are not supported for the following guest operating
//   systems: Windows ME, Windows 98, Windows 95, Windows 3.x, and DOS.
// * In this release, this function requires the virtual machine to be powered
//   on with VMware Tools installed.
//
// Since VMware Workstation 6.0
func (v *VM) SetSharedFolderState(name, hostpath string, options SharedFolderOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_SetSharedFolderState(v.handle, //vmHandle
		C.CString(name),                      //shareName
		C.CString(hostpath),                  //hostPathName
		C.VixMsgSharedFolderOptions(options), //flags
		nil, //callbackProc
		nil) //clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	C.Vix_ReleaseHandle(v.handle)
	return nil
}

// This function returns the state of a shared folder mounted in the virtual
// machine.
//
// Parameters:
//
// index: Identifies the shared folder
//
// Remarks:
//
// * Shared folders are indexed from 0 to n-1, where n is the number of shared
//   folders. Use the function VM.NumSharedFolders() to get the value of n.
// * Shared folders are not supported for the following guest operating systems:
//   Windows ME, Windows 98, Windows 95, Windows 3.x, and DOS.
// * In this release, this function requires the virtual machine to be powered
//   on with VMware Tools installed.
//
// Since VMware Workstation 6.0
func (v *VM) GetSharedFolderState(index int) (string, string, int, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var folderName *C.char
	var folderHostPath *C.char
	var folderFlags *C.int

	jobHandle = C.VixVM_GetSharedFolderState(v.handle, //vmHandle
		C.int(index), //index
		nil,          //callbackProc
		nil)          //clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.getSharedFolder(jobHandle, folderName, folderHostPath, folderFlags)
	defer C.Vix_FreeBuffer(unsafe.Pointer(folderName))
	defer C.Vix_FreeBuffer(unsafe.Pointer(folderHostPath))

	if C.VIX_OK != err {
		return "", "", 0, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	C.Vix_ReleaseHandle(v.handle)
	return C.GoString(folderName), C.GoString(folderHostPath), int(*folderFlags),
		nil
}

// This function pauses a virtual machine. See Remarks section for pause
// behavior when used with different operations.
//
// * This stops execution of the virtual machine.
// * Functions that invoke guest operations should not be called when the
//   virtual machine is paused.
// * Call VM.Resume() to continue execution of the virtual machine.
// * Calling VM.Reset(), VM.PowerOff(), and VM.Suspend() will successfully
//   work when paused. The pause state is not preserved in a suspended virtual
//   machine; a subsequent VM.PowerOn() will not remember the previous
//   pause state.
// * This function is not supported when using the VMWARE_PLAYER provider
//
// Since VMware Workstation 6.5.
func (v *VM) Pause() error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	//var snapshotHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_Pause(v.handle,
		0,                    // options
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	// err = C.getHandle(jobHandle,
	// 	C.VIX_PROPERTY_JOB_RESULT_HANDLE,
	// 	&snapshotHandle,
	// 	C.VIX_PROPERTY_NONE)

	// defer C.Vix_ReleaseHandle(snapshotHandle)
	err = C.VixJob_Wait(jobHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function continues execution of a paused virtual machine.
//
// Remarks:
//
// * This operation continues execution of a virtual machine that was stopped
//   using VM.Pause().
// * Refer to VM.Pause() for pause/unpause behavior with different operations.
// * This function is not supported when using the VMWARE_PLAYER provider
//
// Since VMware Workstation 6.5
func (v *VM) Resume() error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	//var snapshotHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_Unpause(v.handle,
		0,                    // options
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	// err = C.getHandle(jobHandle,
	// 	C.VIX_PROPERTY_JOB_RESULT_HANDLE,
	// 	&snapshotHandle,
	// 	C.VIX_PROPERTY_NONE)

	// defer C.Vix_ReleaseHandle(snapshotHandle)
	err = C.VixJob_Wait(jobHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function powers off a virtual machine.
//
// Parameters:
//
// options:
// Set of VMPowerOption flags to consider when powering off the virtual machine.
//
// Remarks:
// * If you call this function while the virtual machine is suspended or powered
//   off, the operation returns a VIX_E_VM_NOT_RUNNING error.
//   If suspended, the virtual machine remains suspended and is not powered off.
//   If powered off, you can safely ignore the error.
// * If you pass VMPOWEROP_NORMAL as an option, the virtual machine is powered
//   off at the hardware level. Any state of the guest that was not committed
//   to disk will be lost.
// * If you pass VMPOWEROP_FROM_GUEST as an option, the function tries to power
//   off the guest OS, ensuring a clean shutdown of the guest. This option
//   requires that VMware Tools be installed and running in the guest.
// * After VMware Tools begin running in the guest, and VM.WaitForToolsInGuest()
//   returns, there is a short delay before VMPOWEROP_FROM_GUEST becomes
//   available.
//   During this time a job may return error 3009,
//   VIX_E_POWEROP_SCRIPTS_NOT_AVAILABLE.
//   As a workaround, add a short sleep after the WaitForTools call.
// * On a Solaris guest with UFS file system on the root partition, the
//   VMPOWEROP_NORMAL parameter causes an error screen at next power on, which
//   requires user intervention to update the Solaris boot archive by logging
//   into the failsafe boot session from the GRUB menu. Hence, although UFS file
//   systems are supported, VMware recommends using the ZFS file system for
//   Solaris guests.
//
// Since VMware Server 1.0
func (v *VM) PowerOff(options VMPowerOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_PowerOff(v.handle,
		C.VixVMPowerOpOptions(options), // powerOptions,
		nil, // callbackProc,
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// Powers on a virtual machine.
//
// Parameters:
//
// options: VMPOWEROP_NORMAL or VMPOWEROP_LAUNCH_GUI.
//
// Remarks:
// * This operation completes when the virtual machine has started to boot.
//   If the VMware Tools have been installed on this guest operating system, you
//   can call VM.WaitForToolsInGuest() to determine when the guest has finished
//   booting.
// * After powering on, you must call VM.WaitForToolsInGuest() before executing
//   guest operations or querying guest properties.
// * In Server 1.0, when you power on a virtual machine, the virtual machine is
//   powered on independent of a console window. If a console window is open,
//   it remains open. Otherwise, the virtual machine is powered on without a
//   console window.
// * To display a virtual machine with a Workstation user interface, the options
//   parameter must have the VMPOWEROP_LAUNCH_GUI flag, and you must be
//   connected to the host with the VMWARE_WORKSTATION provider flag. If there
//   is an existing instance of the Workstation user interface, the virtual
//   machine will power on in a new tab within that instance.
//   Otherwise, a new instance of Workstation will open, and the virtual machine
//   will power on there.
// * To display a virtual machine with a Player user interface, the options
//   parameter must have the VMPOWEROP_LAUNCH_GUI flag, and you must be
//   connected to the host with the VMWARE_PLAYER flag. A new instance of Player
//   will always open, and the virtual machine will power on there.
// * This function can also be used to resume execution of a suspended virtual
//   machine.
// * The VMPOWEROP_LAUNCH_GUI option is not supported for encrypted virtual
//   machines; attempting to power on with this option results in
//   VIX_E_NOT_SUPPORTED.
//
// Since VMware Server 1.0
func (v *VM) PowerOn(options VMPowerOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_PowerOn(v.handle,
		C.VixVMPowerOpOptions(options), // powerOptions,
		C.VIX_INVALID_HANDLE,
		nil, // callbackProc,
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function resets a virtual machine.
//
// Parameters:
//
// options: Must be VMPOWEROP_NORMAL or VMPOWEROP_FROM_GUEST.
//
// Remarks:
// * If the virtual machine is not powered on when you call this function, it
//   returns an error.
// * If you pass VMPOWEROP_NORMAL as an option, this function is the equivalent
//   of pressing the reset button on a physical machine.
// * If you pass VMPOWEROP_FROM_GUEST as an option, this function tries to reset
//   the guest OS, ensuring a clean shutdown of the guest.
//   This option requires that the VMware Tools be installed and running in the
//   guest.
// * After VMware Tools begin running in the guest, and VM.WaitForToolsInGuest()
//   returns, there is a short delay before VMPOWEROP_FROM_GUEST becomes
//   available. During this time the function may return error 3009,
//   VIX_E_POWEROP_SCRIPTS_NOT_AVAILABLE.
//   As a workaround, add a short sleep after the WaitForTools call.
// * After reset, you must call VM.WaitForToolsInGuest() before executing guest
//   operations or querying guest properties.
// * On a Solaris guest with UFS file system on the root partition, the
//   VMPOWEROP_NORMAL parameter causes an error screen at next power on, which
//   requires user intervention to update the Solaris boot archive by logging
//   into the failsafe boot session from the GRUB menu. Hence, although UFS file
//   systems are supported, VMware recommends using the ZFS file system for
//   Solaris guests.
//
// Since VMware Server 1.0
func (v *VM) Reset(options VMPowerOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_Reset(v.handle,
		C.VixVMPowerOpOptions(options), // powerOptions,
		nil, // callbackProc,
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function suspends a virtual machine.
//
// Remarks:
//
// * If the virtual machine is not powered on when you call this function, the
//   function returns the error VIX_E_VM_NOT_RUNNING.
// * Call VM.PowerOn() to resume running a suspended virtual machine.
//
// Since VMware Server 1.0
func (v *VM) Suspend() error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_Suspend(v.handle,
		0,   // powerOptions,
		nil, // callbackProc,
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle,
		C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function reads variables from the virtual machine state.
// This includes the virtual machine configuration,
// environment variables in the guest, and VMware "Guest Variables"
//
// Parameters:
//
// varType: The type of variable to read. The currently supported values are:
// 		* VM_GUEST_VARIABLE - 	A "Guest Variable". This is a runtime-only
// 								value; it is never stored persistently.
// 								This is the same guest variable that is exposed
// 								through the VMControl APIs, and is a simple way
// 								to pass runtime values in and out of the guest.
// 		* VM_CONFIG_RUNTIME_ONLY - 	The configuration state of the virtual
// 									machine. This is the .vmx file that is
//									stored on the host. You can read this and
// 									it will return the persistent data. If you
// 									write to this, it will only be a runtime
//									change, so changes will be lost when the VM
//									powers off.
// 		* GUEST_ENVIRONMENT_VARIABLE - 	An environment variable in the guest of
//										the VM. On a Windows NT series guest,
//										writing these values is saved
//										persistently so they are immediately
//										visible to every process.
//										On a Linux or Windows 9X guest, writing
//										these values is not persistent so they
//										are only visible to the VMware tools
//										process.
//
// name: The name of the variable.
//
// Remarks:
// * You must call VM.LoginInGuest() before calling this function to read a
//   GUEST_ENVIRONMENT_VARIABLE value.
//   You do not have to call VM.LoginInGuest() to use this function to read a
//   VM_GUEST_VARIABLE or a VVM_CONFIG_RUNTIME_ONLY value.
//
// Since Workstation 6.0
func (v *VM) ReadVariable(varType GuestVarType, name string) (string, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var readValue *C.char

	jobHandle = C.VixVM_ReadVariable(v.handle,
		C.int(varType),
		C.CString(name),
		0,   // options
		nil, // callbackProc
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.readVariable(jobHandle, readValue)

	defer C.Vix_FreeBuffer(unsafe.Pointer(readValue))

	if C.VIX_OK != err {
		return "", &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return C.GoString(readValue), nil
}

// This function writes variables to the virtual machine state.
// This includes the virtual machine configuration, environment variables in
// the guest, and VMware "Guest Variables".
//
// Parameters:
//
// varType:
// The type of variable to write. The currently supported values are:
//		* VM_GUEST_VARIABLE - A "Guest Variable". This is a runtime-only value;
//							  it is never stored persistently.
//							  This is the same guest variable that is exposed
// 							  through the VMControl APIs, and is a simple way to
// 							  pass runtime values in and out of the guest.
//		* VM_CONFIG_RUNTIME_ONLY - 	The configuration state of the virtual
//									machine. This is the .vmx file that is
//									stored on the host.
//									You can read this and it will return the
//									persistent data. If you write to this, it
//									will only be a runtime change, so changes
//									will be lost when the VM powers off.
//									Not supported on ESX hosts.
//		* GUEST_ENVIRONMENT_VARIABLE - An environment variable in the guest of
//									   the VM. On a Windows NT series guest,
//									   writing these values is saved
//									   persistently so they are immediately
//									   visible to every process.
//									   On a Linux or Windows 9X guest, writing
//									   these values is not persistent so they
//									   are only visible to the VMware tools
// 									   process. Requires root or Administrator
// 									   privilege.
//
// name: The name of the variable.
// value: The value to be written.
//
// Remarks:
//
// * The VM_CONFIG_RUNTIME_ONLY variable type is not supported on ESX hosts.
// * You must call VM.LoginInGuest() before calling this function to write a
//   GUEST_ENVIRONMENT_VARIABLE value.
//   You do not have to call VM.LoginInGuest() to use this function to write a
//   VM_GUEST_VARIABLE or a VM_CONFIG_RUNTIME_ONLY value.
// * Do not use the slash '/' character in a VM_GUEST_VARIABLE variable name;
//   doing so produces a VIX_E_INVALID_ARG error.
// * Do not use the equal '=' character in the value parameter; doing so
//   produces a VIX_E_INVALID_ARG error.
// * On Linux guests, you must login as root to change environment variables
//   (when variable type is GUEST_ENVIRONMENT_VARIABLE)
//   otherwise it produces a VIX_E_GUEST_USER_PERMISSIONS error.
// * On Windows Vista guests, when variable type is GUEST_ENVIRONMENT_VARIABLE,
//   you must turn off User Account Control (UAC) in Control Panel >
//   User Accounts > User Accounts > Turn User Account on or off,
//   in order for this function to work.
//
// Since VMware Workstation 6.0
func (v *VM) WriteVariable(varType GuestVarType, name, value string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_WriteVariable(v.handle,
		C.int(varType),
		C.CString(name),
		C.CString(value),
		0,
		nil, // callbackProc
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// Restores the virtual machine to the state when the specified snapshot was
// created.
//
// Parameters:
//
// snapshot: A Snapshot instance. Call VVM.GetRootSnapshot() to get a snapshot
// 			 instance.
// options: Any applicable VMPowerOption. If the virtual machine was powered on
//          when the snapshot was created, then this will determine how the
// 			virtual machine is powered back on. To prevent the virtual machine
// 			from being powered on regardless of the power state when the
// 			snapshot was created, use the  VMPOWEROP_SUPPRESS_SNAPSHOT_POWERON
//			flag. VMPOWEROP_SUPPRESS_SNAPSHOT_POWERON is mutually exclusive
// 			to all other VMPowerOpOptions.
//
// Remarks:
//
// * Restores the virtual machine to the state when the specified snapshot was
//   created.
// 	 This function can power on, power off, or suspend a virtual machine.
//   The resulting power state reflects the power state when the snapshot was
//   created.
// * When you revert a powered on virtual machine and want it to display in the
//   Workstation user interface, options must have the VMPOWEROP_LAUNCH_GUI
//   flag, unless the VMPOWEROP_SUPPRESS_SNAPSHOT_POWERON is used.
// * The ToolsState property of the virtual machine is undefined after the
//   snapshot is reverted.
// * Starting in VMware Workstation 6.5, snapshot operations are allowed on
//   virtual machines that are part of a team.
//   Previously, this operation failed with error code PROPERTY_VM_IN_VMTEAM.
//   Team members snapshot independently so they can have different and
//   inconsistent snapshot states.
// * This function is not supported when using the VMWARE_PLAYER provider
// * If the virtual machine is open and powered off in the UI, this function
//   now closes the virtual machine in the UI before reverting to the snapshot.
//   To refresh this property, you must wait for tools in the guest.
// * After reverting to a snapshot, you must call VM.WaitForToolsInGuest()
//   before executing guest operations or querying guest properties.
//
// Since VMware Server 1.0
func (v *VM) RevertToSnapshot(snapshot *Snapshot, options VMPowerOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_RevertToSnapshot(v.handle,
		snapshot.handle,
		0,                    // options
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// Upgrades the virtual hardware version of the virtual machine to match the
// version of the VIX library.
// This has no effect if the virtual machine is already at the same version or
// at a newer version than the VIX library.
//
// Remarks:
// * The virtual machine must be powered off to do this operation.
// * When the VM is already up-to-date, the function returns without errors.
// * This function is not supported when using the VMWARE_PLAYER provider.
//
// Since VMware Server 1.0
func (v *VM) UpgradeVHardware() error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_UpgradeVirtualHardware(v.handle,
		0,   // options
		nil, // callbackProc
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)

	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function establishes a guest operating system authentication context
// returning an instance of the Guest object
//
// Parameters:
//
// * username: The name of a user account on the guest operating system.
// * password: The password of the account identified by username
// * options: Must be LOGIN_IN_GUEST_NONE or
//            LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT - directs guest
//            commands invoked after the call to
// VM.LoginInGuest() to be run from within the session of the user who is
// interactively logged into the guest operating system.
// See the remarks below for more information about
// LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT.
//
// Remarks:
//
// This function validates the account name and password in the guest OS.
// You must call this function before calling functions that perform operations
// on the guest OS, such as those below. Otherwise you do not need to call this
// function.
// Logins are supported on Linux and Windows. To log in as a Windows Domain
// user, specify the 'username' parameter in the form "domain\username".
// This function does not respect access permissions on Windows 95, Windows 98,
// and Windows ME, due to limitations of the permissions model in those systems.
// Other guest operating systems are not supported for login, including Solaris,
// FreeBSD, and Netware.
// The option LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT should be used to
// ensure that the functions Guest.CaptureScreenImage(), and
// Guest.RunProgramInGuest() work correctly.
//
// All guest operations for a particular VM are done using the identity you
// provide to VM.LoginInGuest().
// As a result, guest operations are restricted by file system privileges in the
// guest OS that apply to the user specified in VM.LoginInGuest(). For example,
// Guest.RmDir() might fail if the user named in VM.LoginInGuest() does not have
// access permissions to the directory in the guest OS.
// VM.LoginInGuest() changes the behavior of Vix functions to use a user account.
// It does not log a user into a console session on the guest OS. As a result,
// you might not see
// the user logged in from within the guest OS. Moreover, operations such as
// rebooting the guest do not clear the guest credentials.
//
// The virtual machine must be powered on before calling this function.
// VMware Tools must be installed and running on the guest OS before calling
// this function.
// You can call VM.WaitForToolsInGuest() to wait for the tools to run.
// Once VM.LoginInGuest() has succeeded, the user session remains valid until
// Guest.Logout() is called successfully,
// VM.LoginInGuest() is called successfully with different user credentials,
// the virtual machine handle's reference count reaches 0, or
// the client applications exits.
// The special login type VIX_CONSOLE_USER_NAME is no longer supported.
// Calling VM.LoginInGuest() with LOGIN_IN_GUEST_NONE as 'options' can be done
// at any time when the VMware Tools are running in the guest.
// Calling VM.LoginInGuest() with the
// LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT flag can only be done when
// there is an interactive user logged into the guest OS. Specifically,
// the "interactive user" refers to the user who has logged into the guest OS
// through the console (for instance, the user who logged into the Windows
// log-in screen).
// The VIX user is the user whose credentials are being provided in the call to
// VM.LoginInGuest().
//
// With LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT, there must be an
// interactive user logged into the guest when the call to VM.LoginInGuest()
// is made, and the VIX user must match the interactive user (they must have
// same account in the guest OS).
//
// Using LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT will ensure that the
// environment in which guest commands are executed is as close as possible to
// the normal environment in which a user interacts with the guest OS. Without
// LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT, commands may be run in a more
// limited environment; however, omitting
// LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT will ensure that commands can
// be run regardless of whether an interactive user is present in the guest.
//
// On Linux guest operating systems, the
// LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT flag requires that X11 be
// installed and running.
//
// Since VMware Server 1.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (v *VM) LoginInGuest(username, password string, options GuestLoginOption) (*Guest, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_LoginInGuest(v.handle,
		C.CString(username), // username
		C.CString(password), // password
		C.int(options),      // options
		nil,                 // callbackProc
		nil)                 // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return nil, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	guest := &Guest{
		handle: v.handle,
	}
	return guest, nil
}

// This function signals the job handle when VMware Tools has successfully
// started in the guest operating system.
// VMware Tools is a collection of services that run in the guest.
//
// Parameters:
//
// * timeout: The timeout in seconds. If VMware Tools has not started by
//            this time, the operation completes with an error.
//            If the value of this argument is zero or negative, then this
//  		  operation will wait indefinitely until the VMware Tools start
//            running in the guest operating system.
//
// Remarks:
//
// * This function signals the job when VMware Tools has successfully started
//   in the guest operating system.
//   VMware Tools is a collection of services that run in the guest.
// * VMware Tools must be installed and running for some Vix functions to
//   operate correctly.
//   If VMware Tools is not installed in the guest operating system, or if the
//   virtual machine
//   is not powered on, this function reports an error.
// * The ToolsState property of the virtual machine object is undefined until
//   VM.WaitForToolsInGuest() reports that VMware Tools is running.
// * This function should be called after calling any function that resets or
//   reboots the state of the guest operating system, but before calling any
//   functions that require VMware Tools to be running. Doing so assures that
//   VMware Tools are once again up and running. Functions that reset the guest
//   operating system in this way include:
//   VM.PowerOn()
//   VM.Reset()
//   VM.RevertToSnapshot()
//
// Since VMware Server 1.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (v *VM) WaitForToolsInGuest(timeout uint) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_WaitForToolsInGuest(v.handle,
		C.int(timeout), // timeoutInSeconds
		nil,            // callbackProc
		nil)            // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

type Guest struct {
	handle                 C.VixHandle
	SharedFoldersParentDir string
}

// Copies a file or directory from the guest operating system to the local
// system (where the Vix client is running).
//
// Parameters:
//
// guestpath: The path name of a file on a file system available to the guest.
// hostpath: The path name of a file on a file system available to the Vix
//           client.
//
func (g *Guest) CopyFileToHost(guestpath, hostpath string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_CopyFileFromGuestToHost(g.handle,
		C.CString(guestpath), // src name
		C.CString(hostpath),  // dest name
		0,                    // options
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function creates a directory in the guest operating system.
//
// Parameters:
//
// path: The path to the directory to be created.
//
// Remarks:
//
// * If the parent directories for the specified path do not exist, this
//   function will create them.
// * If the directory already exists, the error will be set to
//   VIX_E_FILE_ALREADY_EXISTS.
// * Only absolute paths should be used for files in the guest; the resolution
//   of relative paths is not specified.
//
// Since Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) MkDir(path string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_CreateDirectoryInGuest(g.handle,
		C.CString(path),      // path name
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function creates a temporary file in the guest operating system.
// The user is responsible for removing the file when it is no longer needed.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) MkTemp() (string, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var tempFilePath *C.char

	jobHandle = C.VixVM_CreateTempFileInGuest(g.handle,
		0,                    // options
		C.VIX_INVALID_HANDLE, // propertyListHandle
		nil,                  // callbackProc
		nil)                  // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.getTempFilePath(jobHandle, tempFilePath)
	defer C.Vix_FreeBuffer(unsafe.Pointer(tempFilePath))

	if C.VIX_OK != err {
		return "", &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return C.GoString(tempFilePath), nil
}

// This function deletes a directory in the guest operating system.
// Any files or subdirectories in the specified directory will also be deleted.
//
// Parameters:
//
// path: The path to the directory to be deleted.
//
// Remarks:
//
// * Only absolute paths should be used for files in the guest;
//   the resolution of relative paths is not specified.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) RmDir(path string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_DeleteDirectoryInGuest(g.handle,
		C.CString(path), // path name
		0,               // options
		nil,             // callbackProc
		nil)             // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function deletes a file in the guest operating system.
//
// Parameters:
//
// filepath: The path to the file to be deleted.
//
// Remarks:
// * Only absolute paths should be used for files in the guest; the resolution
//   of relative paths is not specified.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) RmFile(filepath string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_DeleteFileInGuest(g.handle,
		C.CString(filepath), // file path name
		nil,                 // callbackProc
		nil)                 // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function tests the existence of a directory in the guest operating
// system.
//
// Parameters:
//
// path: The path to the directory in the guest to be checked.
//
// Remarks:
//
// * Only absolute paths should be used for files in the guest; the resolution
//   of relative paths is not specified.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) IsDir(path string) (bool, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var result C.int

	jobHandle = C.VixVM_DirectoryExistsInGuest(g.handle,
		C.CString(path), // dir path name
		nil,             // callbackProc
		nil)             // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.isFileOrDir(jobHandle, &result)
	if C.VIX_OK != err {
		return false, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	if int(result) == C.FALSE {
		return false, nil
	}

	return true, nil
}

// This function tests the existence of a file in the guest operating system.
// Parameters:
//
// filepath: The path to the file to be tested.
//
// Remarks:
//
// * Only absolute paths should be used for files in the guest; the resolution
//   of relative paths is not specified.
// * If filepath exists as a file system object, but is not a normal file (e.g.
//   it is a directory, device, UNIX domain socket, etc),
//   then VIX_OK is returned, and VIX_PROPERTY_JOB_RESULT_GUEST_OBJECT_EXISTS
//   is set to FALSE.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) IsFile(filepath string) (bool, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var result C.int

	jobHandle = C.VixVM_FileExistsInGuest(g.handle,
		C.CString(filepath), // dir path name
		nil,                 // callbackProc
		nil)                 // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.isFileOrDir(jobHandle, &result)
	if C.VIX_OK != err {
		return false, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	if int(result) == C.FALSE {
		return false, nil
	}

	return true, nil
}

// This function returns information about a file in the guest operating system.
//
// Parameters:
//
// filepath: The path name of the file in the guest.
//
// Remarks:
// * Only absolute paths should be used for files in the guest;
//   the resolution of relative paths is not specified.
// * The function returns the following info parameters:
//		* size: file size as a 64-bit integer. This is 0 for directories.
//		* flags: file attribute flags. The flags are:
//			* FILE_ATTRIBUTES_DIRECTORY: Set if the pathname identifies a
// 			  directory.
//			* FILE_ATTRIBUTES_SYMLINK: Set if the pathname identifies a symbolic
// 			  link file.
//		* modtime:  The modification time of the file or directory as a 64-bit
//					integer specifying seconds since the epoch.
//
// Since VMware Workstation 6.5
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) FileInfo(filepath string) (int64, FileAttr, int64, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var fsize *C.int64
	var flags *C.int
	var modtime *C.int64

	jobHandle = C.VixVM_GetFileInfoInGuest(g.handle,
		C.CString(filepath), // file path name
		nil,                 // callbackProc
		nil)                 // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.getFileInfo(jobHandle, fsize, flags, modtime)
	if C.VIX_OK != err {
		return 0, 0, 0, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return int64(*fsize), FileAttr(*flags), int64(*modtime), nil
}

// This function terminates a process in the guest operating system.
//
// Parameters:
//
// pid: The ID of the process to be killed.
//
// Remarks:
//
// * Depending on the behavior of the guest operating system, there may be a
//   short delay after the job completes before the process truly disappears.
// * Because of differences in how various Operating Systems handle process IDs,
//   Vix may return either VIX_E_INVALID_ARG or VIX_E_NO_SUCH_PROCESS
//   for invalid process IDs.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) Kill(pid uint64) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_KillProcessInGuest(g.handle,
		C.uint64(pid), // file path name
		0,             // options
		nil,           // callbackProc
		nil)           // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

//VixVM_ListDirectoryInGuest
func (g *Guest) Ls() {

}

//VixVM_ListProcessesInGuest
func (g *Guest) Ps() {

}

// This function removes any guest operating system authentication
// context created by a previous call to VM.LoginInGuest().
//
// Remarks:
// * This function has no effect and returns success if VM.LoginInGuest()
//   has not been called.
// * If you call this function while guest operations are in progress,
//   subsequent operations may fail with a permissions error.
// It is best to wait for guest operations to complete before logging out.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) Logout() error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_LogoutFromGuest(g.handle,
		nil, // callbackProc
		nil) // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function runs a program in the guest operating system.
// The program must be stored on a file system available to the guest
// before calling this function.
//
// Parameters:
//
// path: The path name of an executable file on the guest operating system.
// args: A string to be passed as command line arguments to the executable.
// options: Run options for the program. See the remarks below.
//
// Remarks:
//
// * This function runs a program in the guest operating system.
// 	 The program must be stored on a file system available to the guest
//	 before calling this function.
// * The current working directory for the program in the guest is not defined.
//   Absolute paths should be used for files in the guest, including
//   command-line arguments.
// * If the program to run in the guest is intended to be visible to the user
//   in the guest, such as an application with a graphical user interface,
//   you must call VM.LoginInGuest() with
//   LOGIN_IN_GUEST_REQUIRE_INTERACTIVE_ENVIRONMENT as the option before calling
//   this function. This will ensure that the program is run within a
//   graphical session that is visible to the user.
// * If the options parameter is RUNPROGRAM_WAIT, this function will block and
//   return only when the program exits in the guest operating system.
//   Alternatively, you can pass RUNPROGRAM_RETURN_IMMEDIATELY as the value of
//   the options parameter, and this function will return as soon as the program
//	 starts in the guest.
// * For Windows guest operating systems, when running a program with a
//   graphical user interface, you can pass RUNPROGRAM_ACTIVATE_WINDOW as the
//   value of the options parameter. This option will ensure that the
//   application's window is visible and not minimized on the guest's screen.
//   This can be combined with the RUNPROGRAM_RETURN_IMMEDIATELY flag using
//   the bitwise inclusive OR operator (|). RUNPROGRAM_ACTIVATE_WINDOW
//   has no effect on Linux guest operating systems.
// * On a Linux guest operating system, if you are running a program with a
//   graphical user interface, it must know what X Windows display to use,
//   for example host:0.0, so it can make the program visible on that display.
//   Do this by passing the -display argument to the program, if it supports
//   that argument, or by setting the DISPLAY environment variable on the guest.
//   See documentation on VM.WriteVariable()
// * This functions returns three parameters:
//   PROCESS_ID: the process id; however, if the guest has
//   an older version of Tools (those released with Workstation 6 and earlier)
//   and the RUNPROGRAM_RETURN_IMMEDIATELY flag is used, then the process ID
//   will not be returned from the guest and this property will be 0
//   ELAPSED_TIME: the process elapsed time in seconds;
//   EXIT_CODE: the process exit code.
//   If the option parameter is RUNPROGRAM_RETURN_IMMEDIATELY, the latter two
//   will both be 0.
// * Depending on the behavior of the guest operating system, there may be a
//   short delay after the job completes before the process is visible in the
//   guest operating system.
//
// Since VMware Server 1.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) RunProgram(path, args string, options RunProgramOption) (uint64, int, int, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var pid *C.uint64
	var elapsedtime *C.int
	var exitCode *C.int

	jobHandle = C.VixVM_RunProgramInGuest(g.handle,
		C.CString(path),                 //guestProgramName
		C.CString(args),                 //commandLineArgs
		C.VixRunProgramOptions(options), //options
		C.VIX_INVALID_HANDLE,            //propertyListHandle
		nil,                             // callbackProc
		nil)                             // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.runProgramResult(jobHandle, pid, elapsedtime, exitCode)

	if C.VIX_OK != err {
		return 0, 0, 0, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return uint64(*pid), int(*elapsedtime), int(*exitCode), nil
}

// This function runs a script in the guest operating system.
//
// Parameters:
//
// shell: The path to the script interpreter, or NULL to use cmd.exe as
//        the interpreter on Windows.
// script: The text of the script.
// options: Run options for the program. See the remarks below.
//
// Remarks:
//
// * This function runs the script in the guest operating system.
// * The current working directory for the script executed in the guest is
//   not defined. Absolute paths should be used for files in the guest,
//   including the path to the shell or interpreter, and any files referenced
//   in the script text.
// * If the options parameter is RUNPROGRAM_WAIT, this function will block and
//   return only when the program exits in the guest operating system.
//   Alternatively, you can pass RUNPROGRAM_RETURN_IMMEDIATELY as the value of
//   the options parameter, and this makes the function to return as soon as the
//   program starts in the guest.
// * The following properties will be returned:
// 	 PROCESS_ID: the process id; however, if the guest has an older version of
//               Tools (those released with Workstation 6 and earlier) and
//               the RUNPROGRAM_RETURN_IMMEDIATELY flag is used, then the
//               process ID will not be returned from the guest and this
//               property will return 0.
// 	 ELAPSED_TIME: the process elapsed time;
//   PROGRAM_EXIT_CODE: the process exit code.
// * If the option parameter is RUNPROGRAM_RETURN_IMMEDIATELY, the latter two
//   will both be 0.
// * Depending on the behavior of the guest operating system, there may be a
//   short delay after the function returns before the process is visible in the
//   guest operating system.
// * If the total size of the specified interpreter and the script text is
//   larger than 60536 bytes, then the error VIX_E_ARGUMENT_TOO_BIG is returned.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) RunScript(shell, args string, options RunProgramOption) (uint64, int, int, error) {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK
	var pid *C.uint64
	var elapsedtime *C.int
	var exitCode *C.int

	jobHandle = C.VixVM_RunProgramInGuest(g.handle,
		C.CString(shell),                //guestProgramName
		C.CString(args),                 //commandLineArgs
		C.VixRunProgramOptions(options), //options
		C.VIX_INVALID_HANDLE,            //propertyListHandle
		nil,                             // callbackProc
		nil)                             // clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.runProgramResult(jobHandle, pid, elapsedtime, exitCode)

	if C.VIX_OK != err {
		return 0, 0, 0, &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return uint64(*pid), int(*elapsedtime), int(*exitCode), nil
}

// Prepares to install VMware Tools on the guest operating system.
//
// Parameters:
//
// options: May be either INSTALLTOOLS_MOUNT_TOOLS_INSTALLER or
//          INSTALLTOOLS_AUTO_UPGRADE. Either flag can be combined with the
//          INSTALLTOOLS_RETURN_IMMEDIATELY flag using the bitwise inclusive
//          OR operator (|). See remarks for more information.
//
// Remarks:
//
// * If the option INSTALLTOOLS_MOUNT_TOOLS_INSTALLER is provided, the function
//   prepares an ISO image to install VMware Tools on the guest operating system.
//   If autorun is enabled, as it often is on Windows, installation begins,
//   otherwise you must initiate installation.
//   If VMware Tools is already installed, this function prepares to upgrade it
//   to the version matching the product.
// * If the option VIX_INSTALLTOOLS_AUTO_UPGRADE is provided, the function
//   attempts to automatically upgrade VMware Tools without any user interaction
//   required, and then reboots the virtual machine. This option requires that a
//   version of VMware Tools already be installed. If VMware Tools is not
//   already installed, the function will fail.
// * When the option INSTALLTOOLS_AUTO_UPGRADE is used on virtual machine with a
//   Windows guest operating system, the upgrade process may cause the Windows
//   guest to perform a controlled reset in order to load new device drivers.
//   If you intend to perform additional guest operations after upgrading the
//   VMware Tools, it is recommanded that after this task completes, that the
//   guest be reset using VM.Reset() with the VMPOWEROP_FROM_GUEST flag,
//   followed by calling VM.WaitForToolsInGuest() to ensure that the guest has
//   reached a stable state.
// * If the option INSTALLTOOLS_AUTO_UPGRADE is provided and the newest version
//   of tools is already installed, the function will return successfully.
//   Some older versions of Vix may return VIX_E_TOOLS_INSTALL_ALREADY_UP_TO_DATE.
// * If the INSTALLTOOLS_RETURN_IMMEDIATELY flag is set, this function will
//   return immediately after mounting the VMware Tools ISO image.
// * If the INSTALLTOOLS_RETURN_IMMEDIATELY flag is not set for a WS host,
//   this function will return only after the installation successfully completes
//   or is cancelled.
// * The virtual machine must be powered on to do this operation.
// * If the Workstation installer calls for an ISO file that is not downloaded,
//   this function returns an error, rather than attempting to download the ISO
//   file.
//
// Since VMware Server 1.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) InstallTools(options InstallToolsOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_InstallTools(g.handle,
		C.int(options), //options
		nil,            //commandLineArgs
		nil,            //callbackProc
		nil)            //clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

// This function renames a file or directory in the guest operating system.
//
// Parameters:
//
// path1: The path to the file to be renamed.
// path2: The path to the new file.
//
// Remarks:
//
// * Only absolute paths should be used for files in the guest; the resolution
//   of relative paths is not specified.
// * On Windows guests, it fails on directory moves when the destination is on a
//   different volume.
// * Because of the differences in how various operating systems handle
//   filenames, Vix may return either VIX_E_INVALID_ARG or
//   VIX_E_FILE_NAME_TOO_LONG for filenames longer than 255 characters.
//
// Since VMware Workstation 6.0
// Minimum Supported Guest OS: Microsoft Windows NT Series, Linux
func (g *Guest) Mv(path1, path2 string) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_RenameFileInGuest(g.handle,
		C.CString(path1),     //oldName
		C.CString(path2),     //newName
		0,                    //options
		C.VIX_INVALID_HANDLE, //propertyListHandle
		nil,                  //callbackProc
		nil)                  //clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

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

// This function deletes all saved states for the snapshot.
//
// Parameters:
//
// * snapshot: A Snapshot instance. Call VM.RootSnapshot() to get a snapshot
//             instance.
//
// Remarks:
//
// * This function deletes all saved states for the specified snapshot. If the
//   snapshot was based on another snapshot, the base snapshot becomes the new
//   root snapshot.
// * The VMware Server release of the VIX API can manage only a single snapshot
//   for each virtual machine. A virtual machine imported from another VMware
//   product can have more than one snapshot at the time it is imported. In that
//   case, you can delete only a snapshot subsequently added using the VIX API.
// * Starting in VMware Workstation 6.5, snapshot operations are allowed on
//   virtual machines that are part of a team. Previously, this operation
//   failed with error code VIX_PROPERTY_VM_IN_VMTEAM. Team members snapshot
//   independently so they can have different and inconsistent snapshot states.
// * This function is not supported when using the VMWARE_PLAYER provider
// * If the virtual machine is open and powered off in the UI, this function may
//   close the virtual machine in the UI before deleting the snapshot.
//
// Since VMware Server 1.0
func (v *VM) RemoveSnapshot(snapshot *Snapshot, options RemoveSnapshotOption) error {
	var jobHandle C.VixHandle = C.VIX_INVALID_HANDLE
	var err C.VixError = C.VIX_OK

	jobHandle = C.VixVM_RemoveSnapshot(v.handle, //vmHandle
		snapshot.handle,                     //snapshotHandle
		C.VixRemoveSnapshotOptions(options), //options
		nil, //callbackProc
		nil) //clientData

	defer C.Vix_ReleaseHandle(jobHandle)

	err = C.VixJob_Wait(jobHandle, C.VIX_PROPERTY_NONE)
	if C.VIX_OK != err {
		return &VixError{
			code: int(err & 0xFFFF),
			text: C.GoString(C.Vix_GetErrorText(err, nil)),
		}
	}

	return nil
}

type VixError struct {
	code int
	text string
}

func (e *VixError) Error() string {
	return e.text
}
