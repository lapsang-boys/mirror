package level

import (
	"image"
	"path/filepath"

	"github.com/lafriks/go-tiled"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/pkg/errors"
)

type Tileset struct {
	Tiles      []image.Image
	TileWidth  int
	TileHeight int
}

// CreateTileset parses a tiled.Tileset spritesheet.
func CreateTileset(tileset *tiled.Tileset) (*Tileset, error) {
	ts := &Tileset{
		TileWidth:  tileset.TileWidth,
		TileHeight: tileset.TileHeight,
	}

	imagePath := filepath.Join("assets", tileset.Image.Source)

	img, err := imgutil.ReadFile(imagePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	subImgaer, ok := img.(imgutil.SubImager)
	if !ok {
		return nil, errors.Errorf("image %T is not subImager", img)
	}

	xTiles := img.Bounds().Dx() / tileset.TileWidth
	yTiles := img.Bounds().Dy() / tileset.TileHeight

	for yIndex := 0; yIndex < yTiles; yIndex++ {
		for xIndex := 0; xIndex < xTiles; xIndex++ {
			xCord := xIndex * tileset.TileWidth
			yCord := yIndex * tileset.TileHeight
			rect := image.Rect(xCord, yCord, xCord+tileset.TileWidth, yCord+tileset.TileHeight)
			subImg := subImgaer.SubImage(rect)
			ts.Tiles = append(ts.Tiles, subImg)
		}
	}
	return ts, nil
}
