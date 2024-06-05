package palette

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"os"
	"path"
	"sync"

	"github.com/codeation/impress"
)

type Palette struct {
	config *Config
}

func NewPalette() *Palette {
	isCreate := false
	if p, err := newPalette(); err != nil {
		isCreate = os.IsNotExist(errors.Unwrap(err))
	} else {
		return p
	}
	p := &Palette{config: defaultConfig()}
	if isCreate {
		_ = p.save()
	}
	return p
}

func newPalette() (*Palette, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("os.UserHomeDir: %w", err)
	}
	data, err := os.ReadFile(path.Join(home, ".lineation", "config.json"))
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return &Palette{config: &config}, nil
}

func (p *Palette) save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("os.UserHomeDir: %w", err)
	}
	filename := path.Join(home, ".lineation", "config.json")
	if _, err := os.Stat(path.Dir(filename)); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(path.Dir(filename), 0755); err != nil {
			return fmt.Errorf("os.MkdirAll: %w", err)
		}
	}

	data, err := json.MarshalIndent(p.config, "", "    ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("os.WriteFile: %w", err)
	}

	return nil
}

func (p *Palette) DefaultAppRect() image.Rectangle {
	return image.Rectangle{Max: p.config.Window.Size.Point()}
}

// BoxAlign returns space from application edge and any map node border
func (p *Palette) BoxAlign() image.Point {
	return p.config.Boxes.Align.Point()
}

// BoxOffset return space between two map node
func (p *Palette) BoxOffset() image.Point {
	return p.config.Boxes.Offset.Point()
}

func (p *Palette) Columns() int {
	return len(p.config.Boxes.Widths)
}

func (p *Palette) BoxWidth(level int) int {
	if level <= len(p.config.Boxes.Widths) {
		return p.config.Boxes.Widths[level-1]
	}
	return p.config.Boxes.Widths[len(p.config.Boxes.Widths)-1]
}

// TextAlign returns space from border to text for a map node
func (p *Palette) TextAlign() image.Point {
	return p.config.Fonts.Default.Align.Point()
}

// TextLineOffset returns extra space between two text line for a map node
func (p *Palette) TextLineOffset() int {
	return p.config.Fonts.Default.Offset
}

// CursorSize returns cursor size
func (p *Palette) CursorSize() image.Point {
	return image.Pt(p.config.Fonts.Cursor.Width, p.DefaultFont().LineHeight)
}

var once sync.Once
var font *impress.Font

func (p *Palette) DefaultFont() *impress.Font {
	once.Do(func() {
		font = impress.NewFont(p.config.Fonts.Default.Height, p.config.Fonts.Default.Attributes)
	})
	return font
}
