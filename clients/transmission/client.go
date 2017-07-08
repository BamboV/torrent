package transmission

import (
	"github.com/bamboV/torrent/clients/abstract_torrent_client"
	"github.com/tubbebubbe/transmission"
)

type Client struct {
	client transmission.TransmissionClient
}

func NewClient(host string, login string, password string) Client {
	return Client{
		client: transmission.New(host, login, password),
	}
}

func (c Client) HandleMagnet(magnet string) abstract_torrent_client.Response {
	cmd, _ := transmission.NewAddCmdByMagnet(magnet)
	_, err := c.client.ExecuteAddCommand(cmd)

	if err != nil {
		return abstract_torrent_client.Response{
			Status:  500,
			Message: err.Error(),
		}
	}

	return abstract_torrent_client.Response{
		Status:  200,
		Message: "Ok",
	}
}

func New(host string, user string, password string) Client {
	client := transmission.New(host, user, password)

	return Client{
		client: client,
	}
}
