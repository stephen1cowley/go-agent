package server

import (
	"fmt"
	"os/exec"
)

func EditWebsite(AppJSCode string) {
	// Run the shell script with the variable value
	cmd := exec.Command("shell_script/editAppJS.sh", AppJSCode)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output from the shell script
	fmt.Println(string(output))
}
