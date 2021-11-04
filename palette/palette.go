package palette

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/codeation/impress"
)

type Palette struct {
	config *Config
}

type ColorType impress.Color

var (
	black  = impress.NewColor(0, 0, 0)
	white  = impress.NewColor(255, 255, 255)
	gray   = impress.NewColor(192, 192, 192)
	silver = impress.NewColor(239, 239, 239)
	lite   = impress.NewColor(224, 224, 224)
	red    = impress.NewColor(255, 0, 0)
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

func (p *Palette) DefaultAppRect() impress.Rect {
	return impress.NewRect(0, 0, p.config.Window.Size.Width, p.config.Window.Size.Height)
}

func (p *Palette) Color(c ColorType) impress.Color {
	return impress.Color(c)
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

func (p *Palette) HorizontalBoxOffset(level int) int {
	switch level {
	case 1:
		return 0
	case 2:
		return p.config.Boxes.Align.Width + (p.BoxWidth(level)+p.BoxWidth(level-1))/2
	case 3:
		return p.config.Boxes.Align.Width + (p.BoxWidth(level)+p.BoxWidth(level-1))/2
	default:
		return p.config.Boxes.Align.Width
	}
}

func (p *Palette) BoxWidth(level int) int {
	switch level {
	case 1:
		return p.config.Boxes.Widths[0]
	case 2:
		return p.config.Boxes.Widths[1]
	default:
		return p.config.Boxes.Widths[2]
	}
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

func (p *Palette) TextLinePoint(level int, lineno int) impress.Point {
	offset := 0
	if lineno < 0 {
		lineno = 0
	}
	if lineno > 0 {
		offset = (lineno) * p.TextLineOffset()
	}
	return impress.NewPoint(p.HorizontalTextAlign(), p.VerticalTextAlign()+p.DefaultFont().Height*lineno+offset)
}

func (p *Palette) CursorPoint() impress.Point {
	return impress.NewPoint(0, p.TextLineOffset())
}

func (p *Palette) CursorSize() impress.Size {
	return impress.NewSize(p.config.Fonts.Cursor.Width, p.DefaultFont().Height)
}

var once sync.Once
var font *impress.Font

func (p *Palette) DefaultFont() *impress.Font {
	once.Do(func() {
		var err error
		font, err = impress.NewFont(p.config.Fonts.Default.Family, p.config.Fonts.Default.Height)
		if err != nil {
			log.Fatalf("NewFont: %v", err)
		}
	})
	return font
}
