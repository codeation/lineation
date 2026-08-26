package palette

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"os"
	"path"

	"github.com/codeation/impress"
)

type Palette struct {
	config      *Config
	defaultFont *impress.Font
	boldFont    *impress.Font
}

func NewPalette() (*Palette, error) {
	output, err := readPalette()
	if err != nil {
		if os.IsNotExist(errors.Unwrap(err)) {
			p := &Palette{config: defaultConfig()}
			return p, p.save()
		} else {
			return nil, fmt.Errorf("readPalette: %w", err)
		}
	}
	return output, nil
}

func readPalette() (*Palette, error) {
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

// DefaultAppRect returns default application window size
func (p *Palette) DefaultAppRect() image.Rectangle {
	return image.Rectangle{Max: p.config.Window.Size.Point()}
}

// WindowPadding returns space from application edge to map node
func (p *Palette) WindowPadding() image.Point {
	return p.config.Window.Padding.Point()
}

// CursorSize returns cursor size
func (p *Palette) CursorSize() image.Point {
	return image.Pt(p.config.Window.Cursor, p.config.Fonts.Default.LineHeight)
}

// BoxColumns returns column count to fit nodes
func (p *Palette) BoxColumns() int {
	return len(p.config.Nodes.Widths)
}

// BoxWidth returns width of nodes depends on nesting level
func (p *Palette) BoxWidth(level int) int {
	if level <= len(p.config.Nodes.Widths) {
		return p.config.Nodes.Widths[level-1]
	}
	return p.config.Nodes.Widths[len(p.config.Nodes.Widths)-1]
}

// BoxRowCount returns max node rows count
func (p *Palette) BoxRowCount() int {
	return p.config.Nodes.RowCount
}

// BoxMargin returns space between two map nodes
func (p *Palette) BoxMargin() image.Point {
	return p.config.Nodes.Margin.Point()
}

// BoxPadding returns space from border to text in a map node
func (p *Palette) BoxPadding() image.Point {
	return p.config.Nodes.Padding.Point()
}

// TextLineHeight returns line height for a map node
func (p *Palette) TextLineHeight() int {
	return p.config.Fonts.Default.LineHeight
}

func (p *Palette) DialogLeftColumnWidth() int {
	return p.config.Dialog.Columns.Left
}

func (p *Palette) DialogRightColumnWidth() int {
	return p.config.Dialog.Columns.Right
}

func (p *Palette) DialogPadding() image.Point {
	return p.config.Dialog.Padding.Point()
}

func (p *Palette) DialogElemMargin() image.Point {
	return p.config.Dialog.Elements.Margin.Point()
}

func (p *Palette) DialogElemPadding() image.Point {
	return p.config.Dialog.Elements.Padding.Point()
}

func (p *Palette) DialogButtonWidth() int {
	return p.config.Dialog.Elements.Button
}

// FontsOpen opens fonts
func (p *Palette) FontsOpen(app *impress.Application) {
	p.defaultFont = app.NewFont(p.config.Fonts.Default.Height, p.config.Fonts.Default.Attributes)
	p.boldFont = app.NewFont(p.config.Fonts.Bold.Height, p.config.Fonts.Bold.Attributes)
}

// FontsClose closes fonts
func (p *Palette) FontsClose() {
	p.defaultFont.Close()
	p.boldFont.Close()
}

// DefaultFont returns default font
func (p *Palette) DefaultFont() *impress.Font {
	return p.defaultFont
}

func (p *Palette) BoldFont() *impress.Font {
	return p.boldFont
}
