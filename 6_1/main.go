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

		p.Children = append(p.Children, c)

		if eof {
			break
		}
	}

	com := systemIndex["COM"]

	var total int
	countOrbits(com, &total, 0)

	fmt.Println(total)
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