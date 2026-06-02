package main

import (
	"deadlift3008/ai-agents/clients"
	"deadlift3008/ai-agents/execs"
	"fmt"
	"os"
)

// Если в 12 часов ночи идет дождь, то можно ли ожидать, что через 72 часа будет солнечная погода?
// Какое изобретение позволяет смотреть сквозь стены?
func main() {
	openRouterClient := clients.NewOpenRouter(os.Getenv("OPEN_ROUTER_KEY"))

	reactExecutor := execs.NewReactExecutor(openRouterClient)
	res, _ := reactExecutor.ReAct("Если в 12 часов ночи идет дождь, то можно ли ожидать, что через 72 часа будет солнечная погода?", 5)

	fmt.Println(res)

	planExecutor := execs.NewPlanExecutor(openRouterClient)
	answer, _ := planExecutor.PlanExecute("Объясни новичку, что такое AI-агент")

	fmt.Println("\n", answer)
}
