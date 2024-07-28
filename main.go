package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Define the variable value
	value := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>My Simple Website</title>
	</head>
	<body>
		<h1>Welcome to My Simple Website</h1>
		<p>This is a paragraph on my simple website.</p>
	</body>
	</html>
	`

	// Run the shell script with the variable value
	cmd := exec.Command("shell_script/editHtml.sh", value)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output from the shell script
	fmt.Println(string(output))
}
