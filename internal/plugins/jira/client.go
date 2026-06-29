package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const searchEndPoint = "/rest/api/3/search/jql"

var fields = []string{
	"key",
	"summary",
	"status",
	"project",
	"description",
}

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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("jira error %d: %s", resp.StatusCode, string(b))
	}

	fmt.Println(b)

	return b, nil
}

type Issue struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Project struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Fields struct {
	Summary     string  `json:"summary"`
	Status      Status  `json:"status"`
	Project     Project `json:"project"`
	Description any     `json:"description"`
}

type Status struct {
	Name string `json:"name"`
}

func (c *Client) FetchIssues(jql string) ([]Issue, error) {
	body := map[string]any{
		"jql":    jql,
		"fields": fields,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		c.Base+searchEndPoint,
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

	if err := json.Unmarshal(respBytes, &result); err != nil {
		return nil, err
	}

	fmt.Println(result.Issues)

	return result.Issues, nil
}
