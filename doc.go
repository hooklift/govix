/*
Govix is a set of Go bindings for VMware VIX.

The VIX API allows you to automate virtual machine operations on most
current VMware platform products, especially hosted VMware products such as:
vmware workstation, player, fusion and server.

vSphere API, starting from 5.0, merged VIX API in the so-called
GuestOperationsManager managed object. So, we encourage you to use
govsphere for vSphere instad.


Features

This API supports:

  * Adding, removing and listing virtual networks adapters attached to a VM
  * Adding and removing virtual CPUs and memory from a VM
  * Managing virtual switches
  * Managing virtual machines life cycle: power on, power off, reset, pause and resume.
  * Adding and removing shared folders
  * Taking screenshots from a running VM
  * Cloning VMs
  * Creating and removing Snapshots as well as restoring a VM from a Snapshot
  * Upgrading virtual hardware
  * Guest management: login, logout, install vmware tools, etc.

Dynamic library loading

In order for Go to find libvix when running your compiled binary, a govix
path has to be added to the LD_LIBRARY_PATH environment variable.

Example:

  * OSX: export DYLD_LIBRARY_PATH=${GOPATH}/src/github.com/c4milo/govix
  * Linux: export LD_LIBRARY_PATH=${GOPATH}/src/github.com/c4milo/govix
  * Windows: append the path to the PATH environment variable

Be aware that the previous example assumes $GOPATH only has a path set.


Debugging

In order to enable VIX debugging in VMware, you have set the following setting:

  * OSX: echo "vix.debugLevel = \"9\"" >> ~/Library/Preferences/VMware\ Fusion/config
  * Linux: ?
  * Windows: ?


Logs

For logging, the following are the paths in each operating system:

  * OSX: `~/Library/Logs/VMware/*.log`
  * Linux: `/tmp/vmware-<username>/vix-<pid>.log`
  * Windows: `%TEMP%\vmware-<username>\vix-<pid>.log`
*/
package vix
