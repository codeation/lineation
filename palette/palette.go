package palette

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"
	"path"
	"sync"

	"github.com/codeation/impress"
)

type Palette struct {
	config *Config
}

type ColorType color.Color

var (
	black  = color.RGBA{0, 0, 0, 255}
	white  = color.RGBA{255, 255, 255, 255}
	gray   = color.RGBA{192, 192, 192, 255}
	silver = color.RGBA{239, 239, 239, 255}
	lite   = color.RGBA{224, 224, 224, 255}
	red    = color.RGBA{255, 0, 0, 255}
)

var (
	DefaultBackground    = ColorType(white)
	DefaultBoxBackground = ColorType(white)
	DefaultEdge          = ColorType(gray)
	ActiveBoxBackground  = ColorType(silver)
	ActiveEdge           = ColorType(red)
	DefaultText          = ColorType(black)
	CursorBlock          = ColorType(red)
	DefaultLine          = ColorType(lite)
)

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
		return nil, fmt.Errorf("UserHomeDir: %w", err)
	}
	data, err := os.ReadFile(path.Join(home, ".lineation", "config.json"))
	if err != nil {
		return nil, fmt.Errorf("ReadFile: %w", err)
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return &Palette{config: &config}, nil
}

func (p *Palette) save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("UserHomeDir: %w", err)
	}
	filename := path.Join(home, ".lineation", "config.json")
	if _, err := os.Stat(path.Dir(filename)); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(path.Dir(filename), 0755); err != nil {
			return fmt.Errorf("MadirAll: %w", err)
		}
	}

	data, err := json.MarshalIndent(p.config, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("WriteFile: %w", err)
	}

	return nil
}

func (p *Palette) DefaultAppRect() image.Rectangle {
	return image.Rect(0, 0, p.config.Window.Size.Width, p.config.Window.Size.Height)
}

func (p *Palette) Color(c ColorType) color.Color {
	return color.Color(c)
}

func (p *Palette) HorizontalBoxAlign() int {
	return p.config.Boxes.Align.Width
}

func (p *Palette) VerticalBoxAlign() int {
	return p.config.Boxes.Align.Height
}

func (p *Palette) VerticalBoxOffset() int {
	return p.config.Boxes.Offset.Height
}

func (p *Palette) Columns() int {
	return len(p.config.Boxes.Widths)
}

func (p *Palette) HorizontalBoxOffset(level int) int {
	switch {
	case level == 1:
		return 0
	case level <= len(p.config.Boxes.Widths):
		return p.config.Boxes.Align.Width + (p.BoxWidth(level)+p.BoxWidth(level-1))/2
	default:
		return p.config.Boxes.Align.Width
	}
}

func (p *Palette) BoxWidth(level int) int {
	if level <= len(p.config.Boxes.Widths) {
		return p.config.Boxes.Widths[level-1]
	}
	return p.config.Boxes.Widths[len(p.config.Boxes.Widths)-1]
}

func (p *Palette) BoxHeight(level int, linecount int) int {
	offset := 0
	if linecount == 0 {
		linecount = 1
	}
	if linecount > 1 {
		offset = (linecount - 1) * p.TextLineOffset()
	}
	return p.VerticalTextAlign()*2 + p.DefaultFont().Height*linecount + offset + 2
}

func (p *Palette) VerticalTextAlign() int {
	return p.config.Fonts.Default.Align.Height
}

func (p *Palette) TextLineOffset() int {
	return p.config.Fonts.Default.Offset
}

func (p *Palette) HorizontalTextAlign() int {
	return p.config.Fonts.Default.Align.Width
}

func (p *Palette) TextLinePoint(level int, lineno int) image.Point {
	offset := 0
	if lineno < 0 {
		lineno = 0
	}
	if lineno > 0 {
		offset = (lineno) * p.TextLineOffset()
	}
	return image.Pt(p.HorizontalTextAlign(), p.VerticalTextAlign()+p.DefaultFont().Height*lineno+offset)
}

func (p *Palette) CursorPoint() image.Point {
	return image.Pt(0, p.TextLineOffset())
}

func (p *Palette) CursorSize() image.Point {
	return image.Pt(p.config.Fonts.Cursor.Width, p.DefaultFont().Height)
}

var once sync.Once
var font *impress.Font

func (p *Palette) DefaultFont() *impress.Font {
	once.Do(func() {
		font = impress.NewFont(p.config.Fonts.Default.Height, p.config.Fonts.Default.Attributes)
	})
	return font
}
