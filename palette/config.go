package palette

import (
	"image"
)

type Config struct {
	Window   window        `json:"window"` // main application window settings
	Nodes    nodes         `json:"nodes"`  // mind map nodes
	Dialog   dialog        `json:"dialog"` // dialog window settings
	Dialogs_ configDialogs `json:"dialogs"`
	Fonts    fonts         `json:"fonts"` // font selections
}

type window struct {
	Size    point `json:"size"`         // application default size
	Padding point `json:"padding"`      // application window padding
	Cursor  int   `json:"cursor_width"` // cursor width
}

type nodes struct {
	Widths   []int `json:"widths"`    // width of nodes depends on nesting level
	RowCount int   `json:"row_count"` // max node rows count
	Margin   point `json:"margin"`    // space between nodes
	Padding  point `json:"padding"`   // node internal padding
}

type dialog struct {
	Columns  columns  `json:"columns"`  // column widths
	Padding  point    `json:"padding"`  // dialog window padding
	Elements elements `json:"elements"` // element settings
}

type elements struct {
	Margin  point `json:"margin"`       // space between dialog elements
	Padding point `json:"padding"`      // element padding
	Button  int   `json:"button_width"` // button width
}

type fonts struct {
	Default font `json:"default"`
	Bold    font `json:"bold"`
}

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type columns struct {
	Left  int `json:"left_width"`
	Right int `json:"right_width"`
}

type font struct {
	Height     int               `json:"height"`
	Attributes map[string]string `json:"attributes"`
	LineHeight int               `json:"line_height"`
}

func (p point) Point() image.Point { return image.Pt(p.X, p.Y) }

type configDialogs struct {
	Align  configSize `json:"align"`
	Offset configSize `json:"offset"`
	Button int        `json:"button"`
	Width  int        `json:"width"`
}

type configSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (s configSize) Point() image.Point { return image.Pt(s.Width, s.Height) }

func defaultConfig() *Config {
	return &Config{
		Window: window{
			Size:    point{X: 1280, Y: 768},
			Padding: point{X: 20, Y: 16},
			Cursor:  2,
		},
		Nodes: nodes{
			Widths:   []int{100, 140, 300},
			RowCount: 0,
			Margin:   point{X: 20, Y: 16},
			Padding:  point{X: 10, Y: 6},
		},
		Dialog: dialog{
			Columns: columns{Left: 100, Right: 300},
			Padding: point{X: 20, Y: 16},
			Elements: elements{
				Margin:  point{X: 20, Y: 16},
				Padding: point{X: 10, Y: 6},
				Button:  160,
			},
		},
		Fonts: fonts{
			Default: font{
				Attributes: map[string]string{"family": "Verdana"},
				Height:     12,
				LineHeight: 14,
			},
			Bold: font{
				Attributes: map[string]string{"family": "Verdana", "weight": "bold"},
				Height:     12,
				LineHeight: 14,
			},
		},
	}
}
