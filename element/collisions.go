package element

import (
	"math"
)

type Circle struct {
	X, Y, Radius float64
}

func collides(c1, c2 Circle) bool {
	dist := math.Sqrt(math.Pow(c2.X-c1.X, 2) +
		math.Pow(c2.Y-c1.Y, 2))
	if dist < (c1.Radius + c2.Radius) {
		return true
	} else {
		return false
	}

}

func CheckCollisions() error {
	var err error
	for i := 0; i < len(Elements)-1; i++ {
		for j := i + 1; j < len(Elements); j++ {
			if Elements[i].Name != "player" && Elements[j].Name != "player" {
				for _, c1 := range Elements[i].Collisions {
					for _, c2 := range Elements[j].Collisions {
						if collides(c1, c2) && Elements[i].Active && Elements[j].Active {
							err = Elements[j].Collision(Elements[i])
							if err != nil {
								return err
							}
							err = Elements[i].Collision(Elements[j])
							if err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}
	return nil
}
