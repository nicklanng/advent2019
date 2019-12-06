package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Node struct {
	Name string
	Parent *Node
	Children []*Node
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load input: %v\n", err)
	}

	r := bufio.NewReader(bytes.NewReader(b))

	systemIndex := make(map[string]*Node)

	var eof bool
	for {
		parent, child, err := readLine(r)
		if err == io.EOF {
			eof = true
		} else if err != nil {
			panic(err)
		}

		p := findOrCreateNode(parent, systemIndex)
		c := findOrCreateNode(child, systemIndex)

		c.Parent = p
		p.Children = append(p.Children, c)

		if eof {
			break
		}
	}

	myParent := systemIndex["YOU"].Parent
	sanParent := systemIndex["SAN"].Parent

	visited := make(map[*Node]int)
	var open []*Node
	opencosts := make(map[*Node]int)
	visited[myParent] = 0
	open = append(open, myParent.Parent)
	opencosts[myParent.Parent] = 1
	for _, c := range myParent.Children {
		open = append(open, c)
		opencosts[c] = 1
	}

	for {
		if len(open) == 0 {
			break
		}
		var next *Node
		next, open = open[0], open[1:]
		cost := opencosts[next]
		savedCost, ok := visited[next]
		if !ok {
			visited[next] = cost
			if next.Parent != nil {
				open = append(open, next.Parent)
				opencosts[next.Parent] = cost + 1
			}

			for _, c := range next.Children {
				open = append(open, c)
				opencosts[c] =  cost + 1
			}
		} else {
			if savedCost > cost {
				visited[next] = cost
			}
		}

	}

	fmt.Println(visited[sanParent])
}

func readLine(r *bufio.Reader) (parent, child string, err error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	tokens := strings.Split(strings.TrimSpace(line), ")")
	return tokens[0], tokens[1], nil
}

func findOrCreateNode(name string, index map[string]*Node) *Node {
	p, ok := index[name]
	if !ok {
		p = &Node{
			Name: name,
		}
		index[name] = p
	}
	return p
}

func countOrbits(n *Node, count *int, depth int) {
	*count += depth

	for _, c := range n.Children {
		countOrbits(c, count, depth+1)
	}
}