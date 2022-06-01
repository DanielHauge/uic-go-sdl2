package uic_go_sdl2

import "encoding/xml"

// https://www.onlinetool.io/xmltogo/

type Rect struct {
	Text   string `xml:",chardata"`
	X      string `xml:"x"`
	Y      string `xml:"y"`
	Width  string `xml:"width"`
	Height string `xml:"height"`
}

type Property struct {
	Text   string `xml:",chardata"`
	Name   string `xml:"name,attr"`
	Rect   Rect   `xml:"rect"`
	String string `xml:"string"`
	Bool   string `xml:"bool"`
}

type Widget struct {
	Text     string     `xml:",chardata"`
	Class    string     `xml:"class,attr"`
	Name     string     `xml:"name,attr"`
	Property []Property `xml:"property"`
	Widget   []Widget   `xml:"widget"`
}

type QtUi struct {
	XMLName     xml.Name `xml:"ui"`
	Text        string   `xml:",chardata"`
	Version     string   `xml:"version,attr"`
	Class       string   `xml:"class"`
	Widget      Widget   `xml:"widget"`
	Resources   string   `xml:"resources"`
	Connections string   `xml:"connections"`
}
