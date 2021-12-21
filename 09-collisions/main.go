package main

import (
	"bufio"
	"fmt"
	"github.com/furstenheim/challenge_encoding"
	"log"
	"os"
)




func main () {
	reader := bufio.NewReader(os.Stdin)
	challenge := &Challenge{}
	unmarshallErr := challenge_encoding.Unmarshall(challenge, reader)
	if unmarshallErr != nil {
		log.Fatal(unmarshallErr)
	}
	log.Println(challenge.NSprites)
	for _, v := range challenge.Sprites {
		log.Println(v.Height, v.Width)
	}
	world := computeWorld(challenge)
	log.Println(challenge.Sprites)
	for i, g := range(challenge.Cases) {
		result := solveCase(g, world)
		logResult(i, result)
	}
}

func logResult(i int, result string) {
	fmt.Printf("Case #%d: %s\n", i + 1, result)
}

type World struct {
	collisions map[SpriteId]map[SpriteId]map[Position]bool
}

type Position struct {
	X, Y int
}

type SpriteId int


func computeWorld (c *Challenge) World {
	collisions := map[SpriteId]map[SpriteId]map[Position]bool{}

	for i, startSprite := range c.Sprites {
		collisionsForSprite := map[SpriteId]map[Position]bool{}

		for j, endSprite := range c.Sprites {
			if j < i {
				continue
			}
			log.Println(i, j)
			collisionsForSpritePair := map[Position]bool{}

			for dx := -endSprite.Width + 1; dx < startSprite.Width; dx++ {
				for dy := -endSprite.Height + 1; dy < startSprite.Height; dy++ {
					computeCollisions(dx, dy, startSprite, endSprite, collisionsForSpritePair)
				}
			}
			collisionsForSprite[SpriteId(j)] = collisionsForSpritePair
		}
		collisions[SpriteId(i)] = collisionsForSprite
	}

	return World{
		collisions: collisions,
	}
}

func computeCollisions (dx, dy int, startSprite, endSprite Sprite, output map[Position]bool) {
	for x := max(0, dx); x < min(startSprite.Width, dx + endSprite.Width); x++ {
		for y := max(0, dy); y < min(startSprite.Height, dy + endSprite.Height); y++ {
			if startSprite.SpriteRows[y][x] == '0' {
				continue
			}
			if endSprite.SpriteRows[-dy + y][-dx + x] == '1' {
				output[Position{
					X: dx,
					Y: dy,
				}] = true
			}
		}
	}
}


func solveCase (c Case, world World) string {
	count := 0
	for i := 0; i < c.NPositions; i++ {
		for j := i + 1; j < c.NPositions; j++ {
			p1 := c.Positions[i]
			p2 := c.Positions[j]
			if p1.SpriteId > p2.SpriteId {
				p1, p2 = p2, p1
			}
			diff := Position{
				X: p2.X - p1.X,
				Y: p2.Y	- p1.Y,
			}
			if _, ok := world.collisions[p1.SpriteId][p2.SpriteId][diff]; ok {
				count++
			}
		}
	}
	return fmt.Sprintf("%d", count)
}

func min (i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max (i, j int) int {
	if i > j {
		return i
	}
	return j
}


type Challenge struct {
	NCases int    `index:"0"`
	NSprites int `index:"1"`
	Sprites []Sprite `index:"2" indexed:"NSprites"`
	Cases []Case `index:"3" indexed:"NCases"`
}
type Case struct {
	NPositions int `index:"0"`
	Positions []CasePosition `index:"1" indexed:"NPositions"`
}

type CasePosition struct {
	SpriteId SpriteId `index:"0" delimiter:" "`
	X int `index:"1" delimiter:" "`
	Y int `index:"2"`
}

type Sprite struct {
	Width int `index:"0" delimiter:" "`
	Height int `index:"1"`
	SpriteRows []SpriteRow `index:"2" indexed:"Height"`
}

type SpriteRow string