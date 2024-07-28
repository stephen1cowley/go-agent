package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "script.sh")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}
