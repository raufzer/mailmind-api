package integrations

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
)

type GeminiAPIRequest struct {
    Contents []struct {
        Parts []struct {
            Text string `json:"text"`
        } `json:"parts"`
    } `json:"contents"`
}

type GeminiAPIResponse struct {
    Candidates []struct {
        Content struct {
            Parts []struct {
                Text string `json:"text"`
            } `json:"parts"`
        } `json:"content"`
    } `json:"candidates"`
}

func CallGeminiAPI(content, apiKey string) (string, error) {
    url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=%s", apiKey)

    requestBody, err := json.Marshal(GeminiAPIRequest{
        Contents: []struct {
            Parts []struct {
                Text string `json:"text"`
            } `json:"parts"`
        }{
            {
                Parts: []struct {
                    Text string `json:"text"`
                }{{Text: content}},
            },
        },
    })

    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var apiResponse GeminiAPIResponse
    if err := json.Unmarshal(body, &apiResponse); err != nil {
        return "", err
    }

    if len(apiResponse.Candidates) > 0 && len(apiResponse.Candidates[0].Content.Parts) > 0 {
        reply := apiResponse.Candidates[0].Content.Parts[0].Text
        return reply, nil
    }

    return "", errors.New("empty response from Gemini API")
}