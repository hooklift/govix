package vix

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func readVmx(path string) (map[string]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	vmx := make(map[string]string)

	for _, line := range strings.Split(string(data), "\n") {
		values := strings.Split(line, "=")
		if len(values) == 2 {
			vmx[strings.TrimSpace(values[0])] = strings.Trim(strings.TrimSpace(values[1]), `"`)
		}
	}

	return vmx, nil
}

func writeVmx(path string, vmx map[string]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	keys := make([]string, len(vmx))
	i := 0
	for k, _ := range vmx {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(key + " = " + `"` + vmx[key] + `"`)
		buf.WriteString("\n")
	}

	if _, err = io.Copy(f, &buf); err != nil {
		return err
	}

	return nil
}
