package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	itypes "github.com/dbgventures/polarismail-go/internal/types"
	"github.com/dbgventures/polarismail-go/types"
)

const (
	DefaultURL = "https://cp.emailarray.com/admin/json.php"
)

type Client struct {
	baseUrl    string
	httpClient *http.Client
	creds      *types.Credentials
}

func NewClient(creds *types.Credentials, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		baseUrl:    DefaultURL,
		httpClient: httpClient,
		creds:      creds,
	}

	return client, nil
}

func (c Client) Admin() types.Admin {
	return Admin{c: &c}
}

func (c Client) Users() types.Users {
	return Users{c: &c}
}

func (c Client) Domains() types.Domains {
	return nil
}

func (c Client) Lists() types.Lists {
	return nil
}

func (c Client) Aliases() types.Aliases {
	return nil
}

func (c *Client) authenticate() (string, error) {
	formData := url.Values{
		"action":   {"login"},
		"username": {c.creds.Username},
		"password": {c.creds.Password},
	}

	payload := strings.NewReader(formData.Encode())

	var respValue itypes.ApiKeyResponse
	resp, err := c.newRequest(payload, &respValue)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid credentials: %v", resp.Status)
	}

	return respValue.ReturnData, nil
}

func (c *Client) newRequest(payload *strings.Reader, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.baseUrl, payload)
	if err != nil {
		return nil, err
	}

	requestParams := url.Values{}

	req.URL.RawQuery = requestParams.Encode()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
