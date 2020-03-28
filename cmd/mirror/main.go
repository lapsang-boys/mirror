package main

import (
	"image"
	"image/draw"
	"log"

	"github.com/lapsang-boys/mirror/level"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/pkg/errors"
)

func main() {
	lvl, err := level.ParseMap("assets/first-level.tmx")
	if err != nil {
		log.Fatalf("Failed to parse map; %+v", err)
	}

	err = render(lvl)
	if err != nil {
		log.Fatalf("Failed to render map; %+v", err)
	}
}

func render(lvl *level.Level) error {
	bounds := image.Rect(0, 0, lvl.MapWidth(), lvl.MapHeight())
	screen := image.NewRGBA(bounds)

	for x := 0; x < lvl.Width; x++ {
		for y := 0; y < lvl.Height; y++ {
			tileID := lvl.Tiles[x][y]
			tileImg := lvl.Tileset.Tiles[tileID]
			tileRect := lvl.RectAtTile(x, y)
			draw.Draw(screen, tileRect, tileImg, tileImg.Bounds().Min, draw.Src)
		}
	}

	err := imgutil.WriteFile("slask/test.png", screen)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
