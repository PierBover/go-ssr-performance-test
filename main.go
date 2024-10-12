package main

import(
	"fmt"
	"net/http"
	"math"
	"context"
	"strconv"
)

type Tile struct {
	x string
	y string
}

const address string = "127.0.0.1:8081"

func main () {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		tiles := getTiles()
		fmt.Printf("len %v", len(tiles))
		component := Layout(tiles)
		component.Render(context.Background(), w)
	})

	fmt.Printf("listening on http://%s", address)

	http.ListenAndServe(address, nil)
}

func getTiles () []Tile {
	const wrapperWidth = 960.0
    const wrapperHeight = 720.0
    const cellSize = 10.0
    const centerX = wrapperWidth / 2.0
    const centerY = wrapperHeight / 2.0

    angle := 0.0
    radius := 0.0

    var tiles []Tile

	for radius < min(wrapperWidth, wrapperHeight) / 2 {

		x := centerX + math.Cos(angle) * radius
		y := centerY + math.Sin(angle) * radius

		if x >= 0 && x <= wrapperWidth - cellSize && y >= 0 && y <= wrapperHeight - cellSize {
			tile := Tile{
				x: strconv.FormatInt(int64(x), 10),
				y: strconv.FormatInt(int64(y), 10),
			}
			tiles = append(tiles, tile)
		}

		angle += 0.2
		radius += cellSize * 0.015
	}

	return tiles
}