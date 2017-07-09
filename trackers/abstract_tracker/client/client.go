package client

import (
	"encoding/json"
	"github.com/bamboV/torrent"
	"github.com/bamboV/torrent/trackers/abstract_tracker"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TrackerClient struct {
	client http.Client
}

func NewClient(client http.Client) TrackerClient {
	return TrackerClient{
		client: client,
	}
}

func (t *TrackerClient) GetTorrent(trackerURL string, id int) (torrent.Distribution, error) {
	resp, err := t.client.Get(trackerURL + "/distributions/" + strconv.Itoa(int(id)))

	tr := torrent.Distribution{}

	if err != nil {
		return tr, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return tr, err
	}

	err = json.Unmarshal(body, &tr)

	if err != nil {
		return tr, err

	}

	return tr, nil
}

func (t *TrackerClient) Search(trackerURL string, phrase string) ([]torrent.Distribution, error) {
	resp, err := t.client.Get(trackerURL + "/distributions?phrase=" + phrase)

	if err != nil {
		print(err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	tr := []torrent.Distribution{}
	err = json.Unmarshal(body, &tr)

	if err != nil {
		return nil, err
	}

	return tr, nil
}

func (t *TrackerClient) GetOriginalTrackerUrl(trackerURL string) (string, error) {
	resp, err := t.client.Get(trackerURL + "/url")

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	result := abstract_tracker.UrlResponse{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		return "", err
	}

	return result.Url, nil
}
