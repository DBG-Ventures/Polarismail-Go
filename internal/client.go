package internal

import (
	"encoding/json"
	"fmt"
	"io"
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

	_, err := client.authenticate()
	if err != nil {
		return nil, err
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
	return Domains{c: &c}
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
	err := c.newRequest(payload, &respValue)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("invalid credentials")
	}

	return respValue.ReturnData, nil
}

func (c *Client) requestWrapper(formData url.Values, v interface{}) error {
	apiKey, err := c.authenticate()
	if err != nil {
		return err
	}

	formData.Add("token", apiKey)

	payload := strings.NewReader(formData.Encode())

	return c.newRequest(payload, v)
}

func (c *Client) newRequest(payload *strings.Reader, v interface{}) error {
	req, err := http.NewRequest("POST", c.baseUrl, payload)
	if err != nil {
		return err
	}

	requestParams := url.Values{}

	req.URL.RawQuery = requestParams.Encode()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", resp.Status)
	}

	var errCheck itypes.ErrorCheck
	err = json.Unmarshal(body, &errCheck)
	if err != nil {
		fmt.Println("error check")
		return err
	}

	if errCheck.ReturnCode == 0 {
		var errResp itypes.StatusResponse
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return err
		}

		return fmt.Errorf("%v", errResp.ReturnData)
	}

	return json.Unmarshal(body, v)
}
