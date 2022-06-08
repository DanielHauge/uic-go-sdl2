package uic_go_sdl2

import "encoding/xml"

type Rect struct {
	Text   string `xml:",chardata"`
	X      string `xml:"x"`
	Y      string `xml:"y"`
	Width  string `xml:"width"`
	Height string `xml:"height"`
}

type Property struct {
	Text    string  `xml:",chardata"`
	Name    string  `xml:"name,attr"`
	Rect    Rect    `xml:"rect"`
	String  string  `xml:"string"`
	Bool    string  `xml:"bool"`
	Font    Font    `xml:"font"`
	Set     string  `xml:"set"`
	Palette Palette `xml:"palette"`
}

type Font struct {
	XMLName   xml.Name `xml:"font"`
	Text      string   `xml:",chardata"`
	PointSize string   `xml:"pointsize"`
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

type Palette struct {
	XMLName  xml.Name     `xml:"palette"`
	Text     string       `xml:",chardata"`
	Active   PaletteState `xml:"active"`
	Inactive PaletteState `xml:"inactive"`
	Disabled PaletteState `xml:"disabled"`
}

type PaletteState struct {
	Text      string    `xml:",chardata"`
	ColorRole ColorRole `xml:"colorrole"`
}

type ColorRole struct {
	Text  string `xml:",chardata"`
	Role  string `xml:"role,attr"`
	Brush Brush  `xml:"brush"`
}

type Brush struct {
	Text       string `xml:",chardata"`
	BrushStyle string `xml:"brushstyle,attr"`
	Color      Color  `xml:"color"`
}

type Color struct {
	Text  string `xml:",chardata"`
	Alpha string `xml:"alpha,attr"`
	Red   string `xml:"red"`
	Green string `xml:"green"`
	Blue  string `xml:"blue"`
}
