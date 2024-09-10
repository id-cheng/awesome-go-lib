package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var container = "ollama"
var model = "qwen2:0.5b"

func InitModel() {
	cmd := exec.Command("docker", "exec", container, "ollama", "pull", model)
	fmt.Println("init........")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Fprintln(os.Stderr, scanner.Text())
		}
	}()
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

func SendMsg(text string) (string, error) {
	output, err := exec.Command("docker", "exec", container, "ollama", "run", model, text).Output()
	if err != nil {
		return "", err
	}
	return string(output), err
}

func main() {
	InitModel()
	fmt.Println("init model success")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		msg, _ := SendMsg(input)
		fmt.Print("bot:", msg)
	}
}
