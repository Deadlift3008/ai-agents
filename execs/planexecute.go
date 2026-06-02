package execs

import (
	"deadlift3008/ai-agents/clients"
	"fmt"
	"os"
	"strings"
)

const PLANNER_SYSTEM = "Ты - планировщик. Разбей задачу на 3–5 атомарных пунктов, по одному в строке. Только план, не выполняй."
const WRITER_SYSTEM = "Ты - исполнитель. Напиши ОДИН пункт кратко (2–3 предложения), не повторяя уже написанное."

func PlanExecute(question string) (string, error) {
	openRouterClient := clients.NewOpenRouter(os.Getenv("OPEN_ROUTER_KEY"))

	textPlan, _ := openRouterClient.RequestLLM(PLANNER_SYSTEM, []clients.Message{{Role: "user", Content: question}})

	done := make([]string, 0)

	steps := strings.Split(textPlan, "\n")

	for _, step := range steps {
		content := fmt.Sprintf("Задача: %s\nПункт: %s\nУже написано: %s", question, step, strings.Join(done, "\n"))
		part, _ := openRouterClient.RequestLLM(WRITER_SYSTEM, []clients.Message{{Role: "user", Content: content}})
		done = append(done, part)
	}

	return strings.Join(done, "\n"), nil
}
