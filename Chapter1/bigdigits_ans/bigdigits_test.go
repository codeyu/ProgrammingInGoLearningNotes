package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestBigDigits(t *testing.T) {
	log.SetFlags(0)
	log.Println("TEST bigdigits_ans")

	path, _ := os.Getwd()
	expected, err := ioutil.ReadFile(filepath.Join(path, "0123456789.txt"))
	if err != nil {
		t.Fatal(err)
	}
	executable := filepath.Join(path, "bigdigits_ans")
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	command := exec.Command(executable, "-b", "0123456789")
	command.Stdout = writer
	err = command.Run()
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()
	actual, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	reader.Close()
	if bytes.Compare(actual, expected) != 0 {
		t.Fatal("actual != expected")
	}
}
