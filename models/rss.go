package models

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Title string `xml:"title"`
	Guid string `xml:"guid"`
	Link string `xml:"link"`
	ItunesSummary string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd summary"`
	ItunesDuration string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd duration"`
	Descripton string `xml:"description"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Channel struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Language string `xml:"language"`
	Copyright string `xml:"copyright"`
	Items[] Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}
