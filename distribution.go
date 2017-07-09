package torrent

type Distribution struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Size         string `json:"size"`
	MagnetLink   string `json:"magnetLink"`
	DownloadLink string `json:"downloadLink"`
	TopicLink    string `json:"topicLink"`
	LastUpdated  string `json:"lastUpdated"`
}
