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
