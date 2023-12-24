package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	min int
	max int
}

type Range map[string]Interval

type Item struct {
	r Range
	w string
}

type Condition struct {
	rating    string
	greater   bool
	threshold int
	next      string // if true, redirect here
}

type Redirect struct {
	next string
}

type Work interface {
	IsCondition() bool
	GetCondition() Condition
	GetNext() string
}

func (i Interval) GetPossibilities() int {
	return i.max - i.min + 1
}

func (c Condition) IsCondition() bool {
	return true
}

func (c Condition) GetCondition() Condition {
	return c
}

func (c Condition) GetNext() string {
	return c.next
}

func (r Redirect) IsCondition() bool {
	return false
}

func (r Redirect) GetCondition() Condition {
	return Condition{}
}

func (r Redirect) GetNext() string {
	return r.next
}

func (r Range) Copy() Range {
	intervals := make(map[string]Interval)
	for k, v := range r {
		intervals[k] = v
	}
	return intervals
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sections := make([][]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		section := make([]string, 0)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				break
			}
			section = append(section, line)
		}
		if len(section) == 0 {
			break
		}
		sections = append(sections, section)
	}
	workflows := make(map[string][]Work)
	for _, workflow := range sections[0] {
		components := strings.Split(workflow, "{")
		flow := strings.Split(components[1][:len(components[1])-1], ",")
		steps := make([]Work, 0, len(flow))
		for _, work := range flow {
			conditionParts := strings.Split(work, ":")
			if len(conditionParts) == 1 {
				steps = append(steps, Redirect{work})
				continue
			}
			greater := false
			oprands := strings.Split(conditionParts[0], ">")
			if len(oprands) > 1 {
				greater = true
				oprands = strings.Split(conditionParts[0], ">")
			} else {
				oprands = strings.Split(conditionParts[0], "<")
			}
			threshold, _ := strconv.Atoi(oprands[1])
			rating := oprands[0]
			steps = append(steps, Condition{rating, greater, threshold, conditionParts[1]})
		}
		workflows[components[0]] = steps
	}
	possibleRanges := make([]Range, 0)
	start := Interval{1, 4000}
	initialRange := Range{"x": start, "m": start, "a": start, "s": start}
	rangeQueue := make([]Item, 0)
	rangeQueue = append(rangeQueue, Item{initialRange, "in"})
	for len(rangeQueue) > 0 {
		top := rangeQueue[0]
		rangeQueue = rangeQueue[1:]
		if top.w == "A" {
			possibleRanges = append(possibleRanges, top.r)
			continue
		} else if top.w == "R" {
			continue
		}
		topRange := top.r
		workflow := workflows[top.w]
		for _, work := range workflow {
			// handles redirects
			if !work.IsCondition() {
				rangeQueue = append(rangeQueue, Item{topRange, work.GetNext()})
				break
			}
			// handles conditions
			condition := work.GetCondition()
			upper := topRange.Copy()
			lower := topRange.Copy()
			// If the condition is ">", then the split point is after the threshold.
			if condition.greater {
				newMin := max(condition.threshold+1, upper[condition.rating].min)
				upper[condition.rating] = Interval{newMin, upper[condition.rating].max}
				newMax := min(condition.threshold, lower[condition.rating].max)
				lower[condition.rating] = Interval{lower[condition.rating].min, newMax}
				// top continues as the else case, in this case lower
				topRange = lower
				// part that satisfies the condition is enqueued
				rangeQueue = append(rangeQueue, Item{upper, condition.next})
			} else {

				newMin := max(condition.threshold, upper[condition.rating].min)
				upper[condition.rating] = Interval{newMin, upper[condition.rating].max}
				newMax := min(condition.threshold-1, lower[condition.rating].max)
				lower[condition.rating] = Interval{lower[condition.rating].min, newMax}
				// top continues as the else case, in this case upper
				topRange = upper
				// part that satisfies condition is enqueued
				rangeQueue = append(rangeQueue, Item{lower, condition.next})
			}
		}
	}
	out := 0
rangeProcessor:
	for _, r := range possibleRanges {
		for _, v := range r {
			// filter out impossible ranges
			if v.min > v.max {
				continue rangeProcessor
			}
		}
		possibilities := 1
		for _, interval := range r {
			possibilities *= interval.GetPossibilities()
		}
		out += possibilities
	}
	fmt.Println(out)
}
