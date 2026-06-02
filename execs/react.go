package execs

import (
	"deadlift3008/ai-agents/clients"
	"fmt"
	"os"
	"strings"
)

const REACT_SYSTEM = `
Ты решаешь задачу по циклу think → act → observe, без внешних инструментов.
На каждом шаге выводи РОВНО один блок:

Thought: <короткое рассуждение>
Action: PROPOSE: <текущий вариант ответа>
   - или -
Action: FINAL: <ответ, в котором ты уверен>

После PROPOSE я пришлю Observation с просьбой перепроверить.
Перепроверь и либо исправь (снова PROPOSE), либо зафиксируй (FINAL).
Делай по одному шагу за раз, не выкладывай все решение сразу.
На каждом шаге ОБЯЗАТЕЛЬНО найди хотя бы одну возможную ошибку
`

func ReAct(question string, retryCount int) (string, error) {
	openRouterClient := clients.NewOpenRouter(os.Getenv("OPEN_ROUTER_KEY"))

	dialog := []clients.Message{{Role: "user", Content: question}}

	var resText string
	var err error

	for range retryCount {
		resText, err = openRouterClient.RequestLLM(REACT_SYSTEM, dialog)

		if err != nil {
			fmt.Println("GAVNO!", err)
			return "", err
		}

		fmt.Println(resText, "\n---")

		dialog = append(dialog, clients.Message{Role: "assistant", Content: resText})

		if strings.Contains(resText, "Action: FINAL") {
			break
		}

		dialog = append(dialog, clients.Message{Role: "user", Content: "Observation: перечитай свой вариант. Есть ошибка - исправь, иначе зафиксируй FINAL."})
	}

	return resText, nil
}
