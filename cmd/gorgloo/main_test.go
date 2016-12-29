package main

import (
	"bytes"
	"strings"
	"testing"
	"fmt"
)

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("gorgloo -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus\n exp=%s\n got=%s\n\n", ExitCodeOK, status)
	}

	expected := fmt.Sprintf("gorgloo version: %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("Output\n exp=%q\n got=%q\n\n", expected, errStream.String())
	}
}
