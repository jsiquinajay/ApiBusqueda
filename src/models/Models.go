package models

import "encoding/xml"

type FindPersonResponse struct {
	XMLName          xml.Name         `xml:"FindPersonResponse"`
	FindPersonResult FindPersonResult `xml:"FindPersonResult"`
}

type FindPersonResult struct {
	XMLName xml.Name `xml:"FindPersonResult"`
	Name    string   `xml:"Name"`
	SSN     string   `xml:"SSN"`
	DOB     string   `xml:"DOB"`
	Age     string   `xml:"Age"`
	Home    Home     `xml:"Home"`
	Office  Office   `xml:"Office"`
}

type Home struct {
	XMLName xml.Name `xml:"Home"`
	Street  string   `xml:"Street"`
	City    string   `xml:"City"`
	State   string   `xml:"State"`
	Zip     string   `xml:"Zip"`
}

type Office struct {
	XMLName xml.Name `xml:"Office"`
	Street  string   `xml:"Street"`
	City    string   `xml:"City"`
	State   string   `xml:"State"`
	Zip     string   `xml:"Zip"`
}

type ResponseMusic struct {
	Page    int `json:"resultCount"`
	Results []struct {
		WrapperType            string  `json:"wrapperType"`
		Kind                   string  `json:"kind"`
		ArtistId               int     `json:"artistId"`
		CollectionId           int     `json:"collectionId"`
		TrackId                int     `json:"trackId"`
		ArtistName             string  `json:"artistName"`
		CollectionName         string  `json:"collectionName"`
		TrackName              string  `json:"trackName"`
		CollectionCensoredName string  `json:"collectionCensoredName"`
		TrackCensoredName      string  `json:"trackCensoredName"`
		ArtistViewUrl          string  `json:"artistViewUrl"`
		CollectionViewUrl      string  `json:"collectionViewUrl"`
		TrackViewUrl           string  `json:"trackViewUrl"`
		PreviewUrl             string  `json:"previewUrl"`
		ArtworkUrl30           string  `json:"artworkUrl30"`
		ArtworkUrl60           string  `json:"artworkUrl60"`
		ArtworkUrl100          string  `json:"artworkUrl100"`
		CollectionPrice        float32 `json:"collectionPrice"`
		TrackPrice             float32 `json:"trackPrice"`
		ReleaseDate            string  `json:"releaseDate"`
		CollectionExplicitness string  `json:"collectionExplicitness"`
		TrackExplicitness      string  `json:"trackExplicitness"`
		DiscCount              int     `json:"discCount"`
		DiscNumber             int     `json:"discNumber"`
		TrackCount             int     `json:"trackCount"`
		TrackNumber            int     `json:"trackNumber"`
		TrackTimeMillis        int     `json:"trackTimeMillis"`
		Country                string  `json:"country"`
		Currency               string  `json:"currency"`
		PrimaryGenreName       string  `json:"primaryGenreName"`
	} `json:"results"`
}

type Image struct {
	Medio    string `json:"medium"`
	Original string `json:"original"`
}

type Schedule struct {
	Time string `json:"time"`
}

type Show struct {
	ID        int      `json:"id"`
	Url       string   `json:"url"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Language  string   `json:"language"`
	Status    string   `json:"status"`
	Premiered string   `json:"premiered"`
	Ended     string   `json:"ended"`
	Summary   string   `json:"summary"`
	Schedule  Schedule `json:"schedule"`
	Image     Image    `json:"image"`
}

type ResponseTV struct {
	Score float64 `json:"score"`
	Show  Show    `json:"show"`
}

type Country struct {
	ID   int    `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

type ShowPerson struct {
	ID        int     `json:"id"`
	Url       string  `json:"url"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Language  string  `json:"language"`
	Status    string  `json:"status"`
	Premiered string  `json:"premiered"`
	Ended     string  `json:"ended"`
	Summary   string  `json:"summary"`
	Country   Country `json:"country"`
}

type ResponseTVPerson struct {
	Score      float64    `json:"score"`
	ShowPerson ShowPerson `json:"person"`
}

type ReponseSearchTv struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Language  string `json:"language"`
	Status    string `json:"status"`
	Premiered string `json:"premiered"`
	Ended     string `json:"ended"`
	Summary   string `json:"summary"`
	Time      string `json:"time"`
	Original  string `json:"original"`
	UrlOrigin string `json:"urlOrigin"`
}

type ReponseSearchMusic struct {
	ArtistName      string  `json:"artistName"`
	CollectionName  string  `json:"collectionName"`
	TrackName       string  `json:"trackName"`
	CollectionPrice float32 `json:"collectionPrice"`
	TrackPrice      float32 `json:"trackPrice"`
	ReleaseDate     string  `json:"releaseDate"`
	DiscCount       int     `json:"discCount"`
	Country         string  `json:"country"`
	Currency        string  `json:"currency"`
	UrlOrigin       string  `json:"urlOrigin"`
}

type Person struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Address   *Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}
