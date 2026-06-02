package main

import (
	"deadlift3008/ai-agents/clients"
	"fmt"
	"os"
)

func main() {
	openRouterClient := clients.NewOpenRouter(os.Getenv("OPEN_ROUTER_KEY"))

	resText, err := openRouterClient.RequestLLM("ты пират, отвечай как пират", []string{"Сколько время?"})

	if err != nil {
		fmt.Println("ну параша", err)
	}

	fmt.Println(resText)
}
