package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Glossary:
//   - Condition: a part of a workflow in the form of "#>###:#"
//   - Redirect: a part of a workflow that is either just the name of another workflow
//     or "R" or "A".
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sections := make([][]string, 0)
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
	workflows := make(map[string][]string)
	items := make([]map[string]int, 0)
	// process workflows
	for _, line := range sections[0] {
		parts := strings.Split(line, "{")
		instructions := strings.Split(parts[1][:len(parts[1])-1], ",")
		workflows[parts[0]] = instructions
	}
	// process items
	for _, line := range sections[1] {
		item := make(map[string]int)
		parts := strings.Split(line[1:len(line)-1], ",")
		for _, variable := range parts {
			attribute := strings.Split(variable, "=")
			value, _ := strconv.Atoi(attribute[1])
			item[attribute[0]] = value
		}
		items = append(items, item)
	}
	out := 0
itemProcessor:
	for _, item := range items {
		currentWorkflow := "in"
		for {
			if currentWorkflow == "A" {
				for _, v := range item {
					out += v
				}
			}
			if currentWorkflow == "A" || currentWorkflow == "R" {
				continue itemProcessor
			}
			workflow := workflows[currentWorkflow]
			for _, condition := range workflow {
				// check if it actually is a condition and not a redirect
				conditionParts := strings.Split(condition, ":")
				if len(conditionParts) == 1 {
					currentWorkflow = condition
					break
				}
				// now it is guaranteed that this is a condition
				if strings.Contains(condition, ">") {
					// greater than case
					oprands := strings.Split(conditionParts[0], ">")
					threshold, _ := strconv.Atoi(oprands[1])
					if item[oprands[0]] > threshold {
						currentWorkflow = conditionParts[1]
						break
					}
				} else {
					// less than case
					oprands := strings.Split(conditionParts[0], "<")
					threshold, _ := strconv.Atoi(oprands[1])
					if item[oprands[0]] < threshold {
						currentWorkflow = conditionParts[1]
						break
					}
				}
			}
		}
	}
	fmt.Println(out)
}
