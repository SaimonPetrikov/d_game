package astar

import (
	"fmt"
	_ "image/png"
	"math"
)

const (
	tileSize = 16
)

type Node struct {
	Sx        int
	Sy        int
	Neighbors map[string]string
}

var Graph = make(map[string]Node)

var Costs = make(map[string]int)

func heuristic(a string, b string) float64 {
	currentNodeX := float64(Graph[a].Sx / tileSize)
	currentNodeY := float64(Graph[a].Sy / tileSize)

	finishNodeX := float64(Graph[b].Sx / tileSize)
	finishNodeY := float64(Graph[b].Sy / tileSize)

	return math.Abs(currentNodeX-finishNodeX) + math.Abs(currentNodeY-finishNodeY)
}

func FindPath(goal string, start string) []string {

	frontier := NewMinheap[string](600)

	frontier.Push(0, start)

	path := []string{}

	cameFrom := make(map[string]string)

	costSoFar := make(map[string]int)

	cameFrom[start] = start

	costSoFar[start] = 0

	for !frontier.IsEmpty() {
		current := frontier.Pop()
		 fmt.Println(current)
		path = append(path, current)
		if current == goal {
			break
		}
		fmt.Println(Graph[current].Neighbors)
		for key := range Graph[current].Neighbors {
			newCost := costSoFar[current] + Costs[key]
			// fmt.Println(Costs)
			// fmt.Println(costSoFar)
			costSoFarKey, exist := costSoFar[key]
			if !exist || newCost < costSoFarKey {
				// fmt.Println(current)
				// fmt.Println(newCost)

				costSoFar[key] = newCost
				priority := newCost + int(heuristic(goal, key))
				frontier.Push(priority, key)
				cameFrom[key] = current
			}
		}
	}

	return path
}
