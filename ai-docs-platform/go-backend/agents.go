package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

type Agent struct {
	Model string
}

func NewAgent(model string) *Agent {
	return &Agent{Model: model}
}

func (a *Agent) Query(prompt string, context string) (string, error) {
	url := "http://localhost:11434/api/generate"
	fullPrompt := fmt.Sprintf("Context:\n%s\n\nTask:\n%s", context, prompt)

	reqBody, _ := json.Marshal(OllamaRequest{
		Model:  a.Model,
		Prompt: fullPrompt,
		Stream: false,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", err
	}

	return ollamaResp.Response, nil
}

// Specialized Agents
type CodeAnalyzer struct {
	*Agent
}

func NewCodeAnalyzer() *CodeAnalyzer {
	return &CodeAnalyzer{NewAgent("qwen2.5-coder:7b")}
}

func (ca *CodeAnalyzer) AnalyzeFile(fileName string, content string) (string, error) {
	prompt := fmt.Sprintf("Analyze the following source code from file '%s'. Explain its purpose, main functions, and any important logic. Format the output as a concise summary.", fileName)
	return ca.Query(prompt, content)
}
