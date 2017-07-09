package abstract_torrent_client

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
