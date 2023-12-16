package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Box struct {
	labels []string
	lenses map[string]int
}

// adds the lens or replaces the lens if it exists
func (b *Box) Add(key string, val int) {
	_, contained := b.lenses[key]
	if !contained {
		b.labels = append(b.labels, key)
	}
	b.lenses[key] = val
}

func (b *Box) Remove(key string) {
	delete(b.lenses, key)
	for idx, k := range b.labels {
		if k == key {
			b.labels = append(b.labels[:idx], b.labels[idx+1:]...)
		}
	}
}

func hash(in string) int {
	total := 0
	for _, c := range in {
		total += int(c)
		total *= 17
		total %= 256
	}
	return total
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data []string
	for scanner.Scan() {
		in := scanner.Text()
		if len(in) == 0 {
			break
		}
		data = strings.Split(in, ",")
	}
	boxes := make(map[int]*Box)
	// initializing boxes
	for i := 0; i < 256; i++ {
		boxes[i] = &Box{make([]string, 0), make(map[string]int)}
	}
	for _, part := range data {
		if strings.HasSuffix(part, "-") {
			currHash := hash(part[:len(part)-1])
			boxes[currHash].Remove(part[:len(part)-1])
		} else if strings.Contains(part, "=") {
			kv := strings.Split(part, "=")
			currHash := hash(kv[0])
			val, _ := strconv.Atoi(kv[1])
			boxes[currHash].Add(kv[0], val)
		}
	}
	out := 0
	for i, b := range boxes {
		box := i + 1
		for j, s := range b.labels {
			slot := j + 1
			out += box * slot * b.lenses[s]
		}
	}
	fmt.Println(out)
}
