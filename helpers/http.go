package helpers

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func DoRequest(method, url string, body io.Reader, headers map[string]string) (string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", fmt.Errorf("создание запроса: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("выполнение запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("статус ответа: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("чтение ответа: %w", err)
	}

	return string(respBody), nil
}
