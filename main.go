package main

import (
	"deadlift3008/ai-agents/clients"
	"fmt"
	"os"
)

func main() {
	openRouterClient := clients.NewOpenRouter(os.Getenv("OPEN_ROUTER_KEY"))

	resText, err := openRouterClient.RequestLLM("Какой у тебя размер контекста?")

	if err != nil {
		fmt.Println("ну параша", err)
	}

	fmt.Println(resText)
}
