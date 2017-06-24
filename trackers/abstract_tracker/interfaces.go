package abstract_tracker

import "github.com/bamboV/torrent"

type Tracker interface {
	Get(id int) torrent.Distribution
	Search(phrase string) []torrent.Distribution
}
