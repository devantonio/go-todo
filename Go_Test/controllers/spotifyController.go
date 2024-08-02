package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Response map[string]interface{}
	Status   int
}

func getSpotifyApiToken(method, url, credentials string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte("grant_type=client_credentials")))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if credentials != "" {
		req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(credentials)))
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return "", fmt.Errorf("error parsing error response body: %v", err)
		}
		return "", fmt.Errorf("error: %d - %v", resp.StatusCode, errorResponse)
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return "", fmt.Errorf("error parsing response body: %v", err)
	}

	accessToken, ok := jsonResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access token not found or not a string")
	}

	return accessToken, nil
}

func makeSpotifyApiRequest(accessToken, url string) (Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return Response{}, fmt.Errorf("error parsing error response body: %v", err)
		}
		return Response{}, fmt.Errorf("error: %d - %v", resp.StatusCode, errorResponse)
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return Response{}, fmt.Errorf("error parsing response body: %v", err)
	}

	return Response{
		Response: jsonResponse,
		Status:   resp.StatusCode,
	}, nil
}

func spotifyApiRequest() (Response, error) {
	credentials := ""
	tokenUrl := "https://accounts.spotify.com/api/token"
	apiUrl := "https://api.spotify.com/v1/search?query=bedouine*&type=track,episode,show,audiobook&market=US&offset=0&limit=2"

	accessToken, err := getSpotifyApiToken("POST", tokenUrl, credentials)
	if err != nil {
		fmt.Println("Request failed while getting token:", err)
	}

	response, err := makeSpotifyApiRequest(accessToken, apiUrl)
	if err != nil {
		fmt.Println("Request failed while making API request:", err)
	}

	return response, nil

}
