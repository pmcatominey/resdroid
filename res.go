package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// A res directory
type ResDirectory struct {
	path                string
	DrawableDirectories []*DrawableDirectory
}

func NewResDirectory(path string) (*ResDirectory, error) {
	// Check path exists
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	r := &ResDirectory{
		path,
		[]*DrawableDirectory{},
	}
	err = r.buildTree()

	return r, err
}

func (r *ResDirectory) buildTree() error {
	resFiles, err := ioutil.ReadDir(r.path)
	if err != nil {
		return err
	}

	for _, file := range resFiles {
		// Skip non directories
		if !file.IsDir() {
			continue
		}
		path := filepath.Join(r.path, file.Name())

		if strings.HasPrefix(file.Name(), "drawable") {
			// Drawable directory
			if dir, err := NewDrawableDirectory(path); err == nil {
				r.DrawableDirectories = append(r.DrawableDirectories, dir)
			}
		}
	}

	return nil
}
