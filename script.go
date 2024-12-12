package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Build Docker image
	cmd := exec.Command("docker", "build", "-t", "my-go-app", ".")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Docker image built successfully.")

	// Run Docker container
	cmd = exec.Command("docker", "run", "-d", "--name", "go-app-container", "my-go-app")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Docker container is running.")
}
