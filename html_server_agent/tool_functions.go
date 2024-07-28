package server

import (
	"fmt"
	"os/exec"
)

func RunServer() {
	cmd := exec.Command("/bin/sh", "script.sh")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func EditWebsite(htmlCode string) {
	// Run the shell script with the variable value
	cmd := exec.Command("shell_script/editHtml.sh", htmlCode)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output from the shell script
	fmt.Println(string(output))
}
