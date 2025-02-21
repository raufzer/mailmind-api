package integrations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

)

type GenerateContentRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GenerateContentResponse struct {
	Candidates []struct {
		Content struct {
			Parts []Part `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// Free-tier specific limits
const (
	MaxRetries     = 3
	InitialBackoff = 1 * time.Second
)

func CallGeminiAPI(content, apiKey string) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=%s", apiKey)

	requestBody := GenerateContentRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: content},
				},
			},
		},
	}

	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	var resp *http.Response

	// Retry mechanism for rate limiting
	for attempt := 0; attempt < MaxRetries; attempt++ {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
		if err != nil {
			return "", fmt.Errorf("error creating request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		resp, err = client.Do(req)

		// Check if request failed before receiving a response
		if err != nil {
			time.Sleep(InitialBackoff * time.Duration(1<<attempt)) // Exponential backoff
			continue
		}

		// Ensure response body is closed on each retry attempt
		defer resp.Body.Close()

		// If successful, break out of retry loop
		if resp.StatusCode == http.StatusOK {
			break
		}

		// Handle rate limiting (429)
		if resp.StatusCode == 429 {
			time.Sleep(InitialBackoff * time.Duration(1<<attempt)) // Exponential backoff
			continue
		}

		// If another error code is received, no need to retry
		break
	}

	// Final check after retries
	if resp == nil {
		return "", errors.New("failed to get response after retries")
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Handle API errors
	if resp.StatusCode != http.StatusOK {
		var apiError struct {
			Error struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
				Status  string `json:"status"`
			} `json:"error"`
		}

		if err := json.Unmarshal(body, &apiError); err == nil {
			return "", fmt.Errorf("API error [%d %s]: %s",
				apiError.Error.Code,
				apiError.Error.Status,
				apiError.Error.Message)
		}
		return "", fmt.Errorf("unknown API error: %s", string(body))
	}

	// Parse successful response
	var geminiResp GenerateContentResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	if len(geminiResp.Candidates) == 0 ||
		len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("no content in API response")
	}

	return geminiResp.Candidates[0].Content.Parts[0].Text, nil
}
