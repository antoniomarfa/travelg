package flowapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type APIClient struct {
	apiKey    string
	secretKey string
	apiURL    string
}

func NewAPIClient(apiKey, secretKey, apiURL string) *APIClient {
	return &APIClient{
		apiKey:    apiKey,
		secretKey: secretKey,
		apiURL:    apiURL,
	}
}

func (client *APIClient) Send(service string, params map[string]string, method string) (map[string]interface{}, error) {
	method = strings.ToUpper(method)
	url := fmt.Sprintf("%s/%s", client.apiURL, service)

	params["apiKey"] = client.apiKey
	data := client.getPack(params, method)

	signature, err := client.sign(params)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	if method == "GET" {
		response, err = client.httpGet(url, data, signature)
	} else {
		response, err = client.httpPost(url, data, signature)
	}
	if err != nil {
		return nil, err
	}

	if info, ok := response["info"].(map[string]interface{}); ok {
		code := fmt.Sprintf("%v", info["http_code"])
		body, _ := response["output"].(map[string]interface{})
		if code == "200" {
			return body, nil
		} else if code == "400" || code == "401" {
			return nil, fmt.Errorf("%v: %v", body["message"], body["code"])
		} else {
			return nil, fmt.Errorf("unexpected error occurred. HTTP_CODE: %s", code)
		}
	}
	return nil, fmt.Errorf("unexpected error occurred")
}

func (client *APIClient) SetKeys(apiKey, secretKey, apiURL string) {
	client.apiKey = apiKey
	client.secretKey = secretKey
	client.apiURL = apiURL
}

func (client *APIClient) getPack(params map[string]string, method string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var data strings.Builder
	for _, key := range keys {
		if method == "GET" {
			data.WriteString("&" + url.QueryEscape(key) + "=" + url.QueryEscape(params[key]))
		} else {
			data.WriteString("&" + key + "=" + params[key])
		}
	}
	return data.String()[1:]
}

func (client *APIClient) sign(params map[string]string) (string, error) {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var toSign strings.Builder
	for _, key := range keys {
		toSign.WriteString("&" + key + "=" + params[key])
	}

	toSignStr := toSign.String()[1:]

	h := hmac.New(sha256.New, []byte(client.secretKey))
	h.Write([]byte(toSignStr))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (client *APIClient) httpGet(urlStr, data, sign string) (map[string]interface{}, error) {
	fullURL := fmt.Sprintf("%s?%s&s=%s", urlStr, data, sign)
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"output": string(body),
		"info": map[string]interface{}{
			"http_code": resp.StatusCode,
		},
	}, nil
}

func (client *APIClient) httpPost(urlStr, data, sign string) (map[string]interface{}, error) {
	payload := strings.NewReader(data + "&s=" + sign)
	resp, err := http.Post(urlStr, "application/x-www-form-urlencoded", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"output": string(body),
		"info": map[string]interface{}{
			"http_code": resp.StatusCode,
		},
	}, nil
}
