// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vix

import (
	"testing"
)

func TestFindItems(t *testing.T) {
	t.Log("Here we go!")

	host, err := Connect(ConnectConfig{
		Provider: VMWARE_WORKSTATION,
		Options:  VERIFY_SSL_CERT,
	})

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	defer host.Disconnect()

	t.Log("Searching for running vms...")

	urls, err := host.FindItems(FIND_RUNNING_VMS)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	for _, url := range urls {
		t.Logf("%s \n", url)
	}
}
