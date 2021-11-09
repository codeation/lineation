package palette

type Config struct {
	Window configWindow `json:"window"`
	Boxes  configBoxes  `json:"boxes"`
	Fonts  configFonts  `json:"fonts"`
}

type configWindow struct {
	Size configSize `json:"size"`
}

type configSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type configBoxes struct {
	Align  configSize `json:"align"`
	Offset configSize `json:"offset"`
	Widths [3]int     `json:"widths"`
}

type configFonts struct {
	Default configFont   `json:"default"`
	Cursor  configCursor `json:"cursor"`
}

type configFont struct {
	Height     int               `json:"height"`
	Attributes map[string]string `json:"attributes"`
	Align      configSize        `json:"align"`
	Offset     int               `json:"offset"`
}

type configCursor struct {
	Width int `json:"width"`
}

func defaultConfig() *Config {
	return &Config{
		Window: configWindow{
			Size: configSize{Width: 1280, Height: 768},
		},
		Boxes: configBoxes{
			Align:  configSize{Width: 20, Height: 16},
			Offset: configSize{Width: 50, Height: 16},
			Widths: [3]int{100, 140, 300},
		},
		Fonts: configFonts{
			Default: configFont{
				Attributes: map[string]string{"family": "Verdana"},
				Height:     12,
				Align:      configSize{Width: 10, Height: 6}, Offset: 2},
			Cursor: configCursor{Width: 2},
		},
	}
}
