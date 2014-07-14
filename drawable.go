package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// A drawable directory under res
type DrawableDirectory struct {
	Name      string
	Path      string
	Drawables []*Drawable
}

func NewDrawableDirectory(path string) (*DrawableDirectory, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Get all drawables in directory
	drawables := []*Drawable{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, drawableInfo := range files {
		drawable, err := NewDrawable(filepath.Join(path, drawableInfo.Name()))
		// Ignore files which produce errors
		if err == nil {
			drawables = append(drawables, drawable)
		}
	}

	return &DrawableDirectory{
		fileInfo.Name(),
		path,
		drawables,
	}, nil
}

type Filter func(Drawable) bool

func (d *DrawableDirectory) FilteredDrawables(f Filter) []*Drawable {
	filtered := []*Drawable{}
	for _, drawable := range d.Drawables {
		if f(*drawable) {
			filtered = append(filtered, drawable)
		}
	}

	return filtered
}

func (d *DrawableDirectory) HasDrawable(name string) bool {
	for _, drawable := range d.Drawables {
		if drawable.Name == name {
			return true
		}
	}

	return false
}

func (d *DrawableDirectory) Drawable(name string) *Drawable {
	for _, drawable := range d.Drawables {
		if drawable.Name == name {
			return drawable
		}
	}

	return nil
}

type DrawableType string

const (
	// .9.png
	NinePatch DrawableType = "ninepatch"
	// png, jpg or gif
	Bitmap = "bitmap"
	// .xml
	Xml     = "xml"
	Unknown = "unknown"
)

// A drawable file
type Drawable struct {
	Name string
	Path string
	Type DrawableType
}

// Create Drawable from file at path
func NewDrawable(path string) (*Drawable, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &Drawable{
		fileInfo.Name(),
		path,
		drawableType(fileInfo.Name()),
	}, nil
}

// Get type from file extension of name
func drawableType(name string) DrawableType {
	if strings.HasSuffix(name, ".9.png") {
		return NinePatch
	} else if strings.HasSuffix(name, ".png") || strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".gif") {
		return Bitmap
	} else if strings.HasSuffix(name, ".xml") {
		return Xml
	} else {
		return Unknown
	}
}
