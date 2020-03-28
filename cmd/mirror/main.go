package main

import (
	"log"

	"github.com/kr/pretty"
	"github.com/lapsang-boys/mirror/level"
)

func main() {
	l, err := level.ParseMap("assets/first-level.tmx")
	if err != nil {
		log.Fatalf("Failed to parse map; %+v", err)
	}

	pretty.Println(l)

}
