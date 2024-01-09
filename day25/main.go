package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 25

func main() {
	part := flag.Int("part", 0, "Which parts to try?")
	try := flag.Int("try", 0, "Whether to try the real input?")
	flag.Parse()

	fmt.Println("Running day", Day, "( part:", *part, ", try:", *try, ")")

	exampleOne, _ := utils.ReadFileToArray(Day, "example1", false)
	input, _ := utils.ReadFileToArray(Day, "input", false)

	if *part == 0 || *part == 1 {
		solutionOneExample := SolvePartOne(exampleOne)
		fmt.Println("Solution to part 1 (example):", solutionOneExample)

		if *try == 1 {
			SolutionOneInput := SolvePartOne(input)
			fmt.Println("Solution to part 1 (input):", SolutionOneInput)
		}
	}
}

func dictAppend(dict map[string][]string, key string, val string) {
	if list, ok := dict[key]; ok {
		list = append(list, val)
		dict[key] = list
	} else {
		dict[key] = []string{val}
	}
}

func parseData(data []string) (map[string][]string, []string) {
	graph := make(map[string][]string)
	for _, line := range data {
		splitLine := strings.Split(line, ": ")
		key := splitLine[0]
		targets := strings.Split(splitLine[1], " ")

		for _, val := range targets {
			dictAppend(graph, key, val)
			dictAppend(graph, val, key)
		}
	}

	var components []string
	for key := range graph {
		components = append(components, key)
	}

	return graph, components
}

func getPath(graph map[string][]string, a string, b string, path []string) []string {
	// find any path between a and b inside the graph (return first found)
	if len(path) > 100 {
		return []string{} // skip long paths for speed
	}

	current := a
	if current == b {
		return path // have path
	}

	// randomly choose next node from targets
	idxs := rand.Perm(len(graph[current]))

	for _, idx := range idxs {
		next := graph[current][idx]

		if slices.Contains(path, next) {
			continue // skip if already in path
		}

		newPath := append(path, next)
		result := getPath(graph, next, b, newPath)

		if len(result) > 0 {
			return result // have path
		}
	}

	return []string{} // empty path if nothing exists
}

type Edge struct {
	a string
	b string
}

func newEdge(a string, b string) Edge {
	var e Edge
	if a < b {
		e = Edge{a: a, b: b}
	} else {
		e = Edge{a: b, b: a}
	}
	return e
}

func SolvePartOne(data []string) int {
	graph, components := parseData(data)

	edgeCounts := make(map[Edge]int)

	// initialise map of edges, to count how often they're used
	for k, v := range graph {
		for _, target := range v {
			edgeCounts[newEdge(k, target)] = 0
		}
	}

	// find paths between 2 random nodes, keep count of edges from paths
	for i := 0; i < 500; i++ {
		a := components[rand.Intn(len(components))]
		b := components[rand.Intn(len(components))]

		if a == b {
			continue
		}

		path := getPath(graph, a, b, []string{a})

		for j := 0; j < len(path)-1; j++ {
			edgeCounts[newEdge(path[j], path[j+1])]++
		}

	}

	// find most commonly traversed edges
	var edges []Edge
	for k := range edgeCounts {
		edges = append(edges, k)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edgeCounts[edges[j]] < edgeCounts[edges[i]]
	})
	ignore := edges[:3]

	// find elements in group
	start := components[0] // any node to start
	var group []string

	// declare queue
	process := utils.Queue[string]{}
	process.Push(start)

	for {
		if process.Len() == 0 {
			break // have all nodes
		}

		node := process.Pop()
		if slices.Contains(group, node) {
			continue
		}
		group = append(group, node)

		for _, next := range graph[node] {
			e := newEdge(node, next)
			if slices.Contains(ignore, e) {
				continue // skip if edge is ignored
			}
			process.Push(next)
		}

	}

	return len(group) * (len(components) - len(group))
}
