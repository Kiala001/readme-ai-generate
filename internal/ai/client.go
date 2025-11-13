package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GeminiRequest struct {
	Contents []struct {
		Role  string `json:"role"`
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content string `json:"content"`
	} `json:"candidates"`
}

func GenerateReadmeWithGemini(prompt string) (string, error) {
	apiKey := "AIzaSyDVcOgj3e75hKeNRrMolNK1usuWIGM0Y90"
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY não definido")
	}

	reqBody := GeminiRequest{
		Contents: []struct {
			Role  string `json:"role"`
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Role: "user",
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: prompt},
				},
			},
		},
	}

	bodyBytes, _ := json.Marshal(reqBody)
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/gemini-1.5-flash:generateContent?key=%s", apiKey)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("erro Gemini API: %s", resp.Status)
	}

	respBytes, _ := ioutil.ReadAll(resp.Body)
	var gemResp map[string]interface{}
	if err := json.Unmarshal(respBytes, &gemResp); err != nil {
		return "", err
	}

	// A resposta real pode ter outputs[0].content[0].text
	outputs, ok := gemResp["outputs"].([]interface{})
	if !ok || len(outputs) == 0 {
		return "", fmt.Errorf("nenhuma resposta da Gemini")
	}

	firstOutput := outputs[0].(map[string]interface{})
	contentArr := firstOutput["content"].([]interface{})
	text := contentArr[0].(map[string]interface{})["text"].(string)

	return text, nil
}
