package jackett

// SearchResults is the struct of a jackett search response
type SearchResults struct {
	Results  []result  `json:"Results"`
	Indexers []indexer `json:"Indexers"`
}

type result struct {
	//PublishDate          time.Time   `json:"PublishDate"`
	//FirstSeen            string      `json:"FirstSeen"`
	Tracker              string      `json:"Tracker"`
	TrackerID            string      `json:"TrackerId"`
	CategoryDesc         string      `json:"CategoryDesc"`
	BlackholeLink        interface{} `json:"BlackholeLink"`
	Title                string      `json:"Title"`
	GUID                 string      `json:"Guid"`
	Link                 interface{} `json:"Link"`
	Comments             string      `json:"Comments"`
	Category             []int       `json:"Category"`
	Size                 int64       `json:"Size"`
	Files                interface{} `json:"Files"`
	Grabs                interface{} `json:"Grabs"`
	Description          interface{} `json:"Description"`
	RageID               interface{} `json:"RageID"`
	TVDBID               interface{} `json:"TVDBId"`
	Imdb                 interface{} `json:"Imdb"`
	TMDb                 interface{} `json:"TMDb"`
	Seeders              int         `json:"Seeders"`
	Peers                int         `json:"Peers"`
	BannerURL            interface{} `json:"BannerUrl"`
	InfoHash             interface{} `json:"InfoHash"`
	MagnetURI            string      `json:"MagnetUri"`
	MinimumRatio         float64     `json:"MinimumRatio"`
	MinimumSeedTime      int         `json:"MinimumSeedTime"`
	DownloadVolumeFactor float64     `json:"DownloadVolumeFactor"`
	UploadVolumeFactor   float64     `json:"UploadVolumeFactor"`
	Gain                 float64     `json:"Gain"`
}

type indexer struct {
	ID      string      `json:"ID"`
	Name    string      `json:"Name"`
	Status  int         `json:"Status"`
	Results int         `json:"Results"`
	Error   interface{} `json:"Error"`
}
