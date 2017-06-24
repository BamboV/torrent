package rutor

import (
	"github.com/bamboV/torrent"
	"github.com/bamboV/torrent/trackers/abstract_tracker"
)

type RutorTracker struct {
	parser Parser
}

func NewTracker(baseURL string) abstract_tracker.Tracker {
	return RutorTracker{
		parser: Parser{
			BaseURL: baseURL,
		},
	}
}

func (t RutorTracker) Get(id int) torrent.Distribution {
	distribution := t.parser.ParsePage(id)

	return t.transformDistribution(distribution)
}

func (t RutorTracker) Search(phrase string) []torrent.Distribution {
	distributions, _ := t.parser.Parse(phrase)

	return t.transformDistributions(distributions)
}

func (t *RutorTracker) transformDistribution(distribution torrent.Distribution) torrent.Distribution {
	return torrent.Distribution{
		LastUpdated:  distribution.LastUpdated,
		Size:         distribution.Size,
		DownloadLink: distribution.DownloadLink,
		TopicLink:    distribution.TopicLink,
		MagnetLink:   distribution.MagnetLink,
		Id:           distribution.Id,
		Title:        distribution.Title,
	}
}

func (t *RutorTracker) transformDistributions(distributions []torrent.Distribution) []torrent.Distribution {
	result := []torrent.Distribution{}

	for _, value := range distributions {
		result = append(result, t.transformDistribution(value))
	}

	return result
}
