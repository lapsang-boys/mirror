package level

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"

	"github.com/lafriks/go-tiled"
	"github.com/pkg/errors"
)

// Level is a map level.
type Level struct {
	Width   int
	Height  int
	Tiles   [][]int
	Tileset *Tileset
}

// NewLevel returns a new Level of the specified dimensions.
func NewLevel(width int, height int) *Level {
	var l *Level
	l = &Level{
		Width:  width,
		Height: height,
		Tiles:  makeTiles(width, height),
	}
	return l
}

// ParseMap parses the given tiled tmx-map and returns the corresponding level.
func ParseMap(tmxPath string) (*Level, error) {
	tmx, err := tiled.LoadFromFile(tmxPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	width := tmx.Width
	height := tmx.Height
	level := NewLevel(width, height)
	level.Tiles = parseTiles(level, tmx.Layers[0], width, height)
	ts, err := parseTileset(tmx.Tilesets[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	level.Tileset = ts

	return level, nil
}

func parseTiles(level *Level, layer *tiled.Layer, width int, height int) [][]int {
	tiles := makeTiles(width, height)

	for i := range layer.Tiles {
		x := i % width
		y := i / width
		tiles[x][y] = int(layer.Tiles[i].ID)
	}
	return tiles
}

func makeTiles(width int, height int) [][]int {
	tiles := make([][]int, width)
	for i := range tiles {
		tiles[i] = make([]int, height)
	}
	return tiles
}

func parseTileset(tileset *tiled.Tileset) (*Tileset, error) {
	tilesetSource := filepath.Join("assets", tileset.Source)
	data, err := ioutil.ReadFile(tilesetSource)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	output := &tiled.Tileset{}
	err = xml.Unmarshal(data, output)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ts, err := CreateTileset(output)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ts, nil
}
