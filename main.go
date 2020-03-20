package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	noun := strings.Title(nouns[rand.Intn(len(nouns))])
	adj := strings.Title(adjectives[rand.Intn(len(adjectives))])
	fmt.Println(adj, noun)
}
