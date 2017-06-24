package abstract_tracker

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
}

type urlResponse struct {
	url string
}

func (s *Server) Start(trackerURL string, tracker Tracker) {
	mux := http.NewServeMux()

	mux.HandleFunc("/search/", func(w http.ResponseWriter, r *http.Request) {
		phrase := strings.Trim(r.URL.Path, "/search/")
		data := tracker.Search(phrase)
		str, _ := json.Marshal(data)
		w.Write(str)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(strings.Trim(r.URL.Path, "/"))
		data := tracker.Get(id)
		str, _ := json.Marshal(data)
		w.Write(str)
	})

	mux.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		str, _ := json.Marshal(urlResponse{
			url: trackerURL,
		})

		w.Write(str)
	})

	http.ListenAndServe("/", mux)
}
