package abstract_tracker

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
}

type UrlResponse struct {
	Url string
}

func (s *Server) Start(trackerURL string, tracker Tracker) {
	mux := http.NewServeMux()

	mux.HandleFunc("/distributions", func(w http.ResponseWriter, r *http.Request) {
		phrase := r.URL.Query().Get("phrase")

		if phrase == "" {
			w.WriteHeader(400)
			return
		}

		data := tracker.Search(phrase)
		str, _ := json.Marshal(data)
		w.Write(str)
	})

	mux.HandleFunc("/distributions/", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(strings.Trim(r.URL.Path, "/distributions/"))
		data := tracker.Get(id)
		str, _ := json.Marshal(data)
		w.Write(str)
	})

	mux.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		str, _ := json.Marshal(UrlResponse{
			Url: trackerURL,
		})

		w.Write(str)
	})

	http.ListenAndServe("/", mux)
}
