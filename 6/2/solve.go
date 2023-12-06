package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data []int // first row = duration, 2nd row = distance threshold
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		line = strings.ReplaceAll(line, " ", "")
		val, _ := strconv.Atoi(strings.Split(line, ":")[1])
		data = append(data, val)
	}
	dur := data[0]
	dist := data[1]
	l := 0
	for {
		if l*(dur-l) > dist {
			break
		}
		l++
	}
	fmt.Println(dur - 2*l + 1)
}
