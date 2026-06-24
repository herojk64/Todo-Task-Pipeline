package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Base  string
	Email string
	Token string
	HTTP  *http.Client
}

func NewClient(base, email, token string) *Client {
	return &Client{
		Base:  base,
		Email: email,
		Token: token,
		HTTP:  &http.Client{},
	}
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.Email, c.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("jira error %d: %s", resp.StatusCode, string(b))
	}

	return b, nil
}

type Issue struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Summary string `json:"summary"`
	Status  Status `json:"status"`
}

type Status struct {
	Name string `json:"name"`
}

func (c *Client) FetchIssues() ([]Issue, error) {
	body := map[string]any{
		"jql":    "project IS NOT EMPTY",
		"fields": []string{"key", "summary", "status"},
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		c.Base+"/rest/api/3/search/jql",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Issues []Issue `json:"issues"`
	}

	err = json.Unmarshal(respBytes, &result)
	return result.Issues, err
}
