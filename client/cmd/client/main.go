package main

import (
	"github.com/IceRain26/ToolHub/client/internals/ai"
	"os"
	"fmt"
	"bufio"
)


func main() {
	client := ai.NewOllama("http://localhost:11434", "llama3.1")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("ToolHub CLI v0.1")
	fmt.Println("-------------------")

	for {
		fmt.Print("\n> ")

		input, _ := reader.ReadString('\n')
		response, err := client.Chat(input)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Println("\nAI:", response)


	}
}