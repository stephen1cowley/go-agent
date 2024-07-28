package server

import (
	"fmt"
	"os/exec"
)

func Mult3(num1 int, num2 int, num3 int) int {
	return (num1 * num2 * num3)
}

func Mult4(num1 int, num2 int, num3 int, num4 int) int {
	return (num1 * num2 * num3 * num4)
}

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

}
