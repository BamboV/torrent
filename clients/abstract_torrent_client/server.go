package abstract_torrent_client

import (
	"encoding/json"
	"net/http"
)

type Server struct {
}

type magnet struct {
	Magnet string `json:"magnet"`
}

func (server *Server) Start(client TorrentClient) {
	http.HandleFunc("/magnet/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			decoder := json.NewDecoder(r.Body)
			magnet := magnet{}

			decoder.Decode(&magnet)
			data := client.HandleMagnet(magnet.Magnet)
			str, _ := json.Marshal(data)
			w.Write(str)
		}
	})
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		panic(err)
	}
}
