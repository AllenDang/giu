package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func save(name, data string) {
	const newFileMode = 0o644
	if err := ioutil.WriteFile(name, []byte(data), newFileMode); err != nil {
		log.Fatalf("Failed to save %s:%v\n", name, err)
	}
}

func mkdirAll(name string) {
	const newDirMode = 0o755
	if err := os.MkdirAll(name, newDirMode); err != nil {
		log.Fatalf("Failed to make all dir, %s:%v\n", name, err)
	}
}

func runCmd(cmd *exec.Cmd) {
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to execute command:%s with error %v\n", cmd.String(), err)
	}
}
