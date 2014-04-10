# VMware VIX API for GO

The VIX API allows you to automate virtual machine operations on most current VMware platform products, especially hosted VMware products such as: vmware workstation, player, fusion and server. 

vSphere API, starting from 5.0, merged VIX API in the so-called GuestOperationsManager managed object. So, we encourage you to use govsphere for vSphere.

## To keep in mind
In order for Go to find libvix when running your compiled binary, a govix path has to be added to the LD_LIBRARY_PATH environment variable. Example:

* **OSX:** export DYLD_LIBRARY_PATH=${GOPATH}/src/github.com/c4milo/govix
* **Linux:** export LD_LIBRARY_PATH=${GOPATH}/src/github.com/c4milo/govix

Be aware that the previous example assumes $GOPATH only has a path set.

Additionally, due to an [issue](https://code.google.com/p/go/issues/detail?id=5428) in Go, the CGO LDFLAGS
directive has to have an absolute path in order for govix to compile.


## License
Copyright 2014 Cloudescape. All rights reserved.
