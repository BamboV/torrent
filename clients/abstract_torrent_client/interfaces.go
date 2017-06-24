package abstract_torrent_client

type TorrentClient interface {
	HandleMagnet(magnet string) Response
}
