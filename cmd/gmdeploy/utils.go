package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Save(name, data string) {
	err := ioutil.WriteFile(name, []byte(data), 0644)
	if err != nil {
		log.Fatalf("Failed to save %s:%v\n", name, err)
	}
}

func MkdirAll(name string) {
	err := os.MkdirAll(name, 0755)
	if err != nil {
		log.Fatalf("Failed to make all dir, %s:%v\n", name, err)
	}
}

func RunCmd(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute command:%v\n", err)
	}
}
