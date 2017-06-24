package abstract_torrent_client

import (
	"encoding/json"
	"net/http"
)

type Server struct {
}

func (server *Server) Start(client TorrentClient) {
	http.HandleFunc("/magnet/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			magnet := r.PostForm.Get("magnet")
			data := client.HandleMagnet(magnet)
			str, _ := json.Marshal(data)
			w.Write(str)
		}
	})
	http.ListenAndServe(":80", nil)
}
