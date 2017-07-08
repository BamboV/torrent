package client

import (
	"bytes"
	"fmt"
	"net/http"
)

type Client struct {
	client http.Client
}

func NewClient(client http.Client) Client {
	return Client{
		client: client,
	}
}

func (c *Client) DownloadByMagnet(clientURL string, magnet string) bool {
	var jsonStr = []byte(fmt.Sprintf(`{"magnet":"%v"}`, magnet))
	resp, err := c.client.Post(clientURL, "application/json", bytes.NewBuffer(jsonStr))

	if err != nil {
		return false
	}

	return resp.StatusCode == 200
}
