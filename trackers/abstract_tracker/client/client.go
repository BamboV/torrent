package client

import (
	"encoding/json"
	"github.com/bamboV/torrent"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TrackerClient struct {
	client http.Client
}

func (t *TrackerClient) GetTorrent(trackerURL string, id int) (torrent.Distribution, error) {
	resp, err := t.client.Get(trackerURL + "/" + strconv.Itoa(int(id)))

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	tr := torrent.Distribution{}
	err = json.Unmarshal(body, &tr)

	if err != nil {
		return nil, err

	}

	return tr, nil
}

func (t *TrackerClient) Search(trackerURL string, phrase string) ([]torrent.Distribution, error) {
	resp, err := t.client.Get(trackerURL + "/search/" + phrase)

	if err != nil {
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
