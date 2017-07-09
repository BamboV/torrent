package client

import (
	"bytes"
	"fmt"
	"net/http"
)

type Client struct {
	client    http.Client
	clientURL string
}

func NewClient(client http.Client, clientURL string) Client {
	return Client{
		client:    client,
		clientURL: clientURL,
	}
}

func (c *Client) DownloadByMagnet(magnet string) bool {
	var jsonStr = []byte(fmt.Sprintf(`{"magnet":"%v"}`, magnet))
	resp, err := c.client.Post(c.clientURL+"/magnet", "application/json", bytes.NewBuffer(jsonStr))

	if err != nil {
		return false
	}

	return resp.StatusCode == 200
}
