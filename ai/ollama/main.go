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
	output, err := exec.Command("docker", "exec", container, "ollama", "pull", model).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
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
