package _5

import (
	"adventofcode2024/challenges/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Update struct {
	pages []int
}

type Graph struct {
	vertices map[int][]int
}

func newGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

func (g *Graph) addEdge(u, v int) {
	g.vertices[u] = append(g.vertices[u], v)
}

func PrintQueue(filenameOrdering, filenameUpdates string) (int, int, error) {
	// read files
	ordering, err := util.ReadLines(filenameOrdering)
	if err != nil {
		return 0, 0, err
	}
	updateStrings, err := util.ReadLines(filenameUpdates)
	if err != nil {
		return 0, 0, err
	}

	// parse ordering
	graph, err := parseOrdering(ordering)
	if err != nil {
		return 0, 0, err
	}
	log.Printf("Parsed Ordering-Rules: %v", graph)

	// parse updateStrings
	updates, err := parseUpdates(updateStrings)
	if err != nil {
		return 0, 0, err
	}

	// filter out the updates that are incorrect
	validUpdate, invalidUpdates := filterUpdates(updates, graph)

	log.Printf("Updates: %v", validUpdate)

	// add the middle-numbers of the updates
	sum := 0
	for _, update := range validUpdate {
		wantPageAtIndex := len(update.pages) / 2
		sum += update.pages[wantPageAtIndex]
	}

	// fix ordering of invalid updates
	invalidUpdates = fixOrdering(invalidUpdates, graph)
	log.Printf("Invalid Updates ordered: %v", invalidUpdates)
	sumInvalid := 0
	for _, update := range invalidUpdates {
		wantPageAtIndex := len(update.pages) / 2
		sumInvalid += update.pages[wantPageAtIndex]
	}

	return sum, sumInvalid, nil
}

// fixOrdering iterates through all updates and fixes the ordering for each
func fixOrdering(invalidUpdates []Update, graph *Graph) []Update {
	// topologically sort the rules
	for i, update := range invalidUpdates {
		update = fixPageOrdering(update, graph)

		invalidUpdates[i] = update
	}

	return invalidUpdates
}

// fixPageOrdering takes an update and fixes the ordering of the pages
func fixPageOrdering(update Update, graph *Graph) Update {
	result := Update{pages: make([]int, len(update.pages))}
	// iterate through the pages and copy them into place
	for _, page := range update.pages {
		// for all consecutive pages we need to go through the already copied pages and check if the current page has
		// to be inserted before the inserted page
		for j, insertedPage := range result.pages {
			// if the inserted page number is 0, we can insert the page at this place
			if insertedPage == 0 {
				result.pages[j] = page
				break
			}
			// check if the page has to be inserted before the existing page
			if laterNeighbours, ok := graph.vertices[page]; ok {
				if contains(laterNeighbours, insertedPage) {
					//log.Printf("Page [%d] has to be inserted before page [%d]", page, insertedPage)
					// shift all pages from that index to the right and insert the page at the current index
					for k := len(result.pages) - 1; k > j; k-- {
						result.pages[k] = result.pages[k-1]
					}
					result.pages[j] = page
					break
				}
			}
			// the page is not mentioned in the ordering rules, so we can insert it at the end
		}
	}

	return result
}

// filterUpdates splits the updates into valid and invalid updates
func filterUpdates(updates []Update, orderingRules *Graph) ([]Update, []Update) {
	// iterate through all updates
	corruptedUpdates := make([]int, 0)
	for i, update := range updates {
		isCorrupted, corruptedPageIndex := isCorruptedUpdate(update, orderingRules, i)
		if isCorrupted {
			log.Printf("Update [%d] is corrupted at page [%d]", i, update.pages[corruptedPageIndex])
			// remove the update
			corruptedUpdates = append(corruptedUpdates, i)
		}
	}

	invalidUpdates := make([]Update, 0)
	validUpdates := make([]Update, 0)
	for i, update := range updates {
		// only add the update if it is not in the corruptedUpdates
		if contains(corruptedUpdates, i) {
			invalidUpdates = append(invalidUpdates, update)
		} else {
			validUpdates = append(validUpdates, update)
		}
	}
	return validUpdates, invalidUpdates
}

// isCorruptedUpdate checks if an update is valid according to the orderingRules
// returns true if the update is corrupted
func isCorruptedUpdate(update Update, orderingRules *Graph, i int) (bool, int) {
	isCorrupted := false
	corruptedPageIndex := -1
	// iterate through all pages
pageLoop:
	for pageIndex, page := range update.pages {
		// check if the page is in the orderingRules
		if rule, ok := orderingRules.vertices[page]; ok {
			// page is in the orderingRules
			//log.Printf("Update [%d] - Page [%d] is in the orderingRules - rule: %v", i, page.number, rule)
			// check the index of the page and compare it with the indices of all of the other pages in the rule.
			// the page must have a lower index than all of the other pages in the rule
			for _, otherPage := range rule {
				// find the index of the other page in the "update"
				otherPageIndex := -1
				for j, pageInUpdate := range update.pages {
					if pageInUpdate == otherPage {
						otherPageIndex = j
						break
					}
				}
				// check if the otherPageIndex is before the current pageIndex
				if otherPageIndex >= 0 && otherPageIndex < pageIndex {
					// the other page is before the current page
					//log.Printf("Update [%d] - Page [%d] is before Page [%d], corrupted Update", i, page,
					//    otherPage)
					isCorrupted = true
					corruptedPageIndex = pageIndex
					break pageLoop
				}
			}
		}
	}
	return isCorrupted, corruptedPageIndex
}

// contains checks if an integer is in an array
func contains(array []int, i int) bool {
	for _, value := range array {
		if value == i {
			return true
		}
	}
	return false
}

// parseUpdates parses the input strings into a 2D slice of integers
func parseUpdates(input []string) ([]Update, error) {
	result := make([]Update, 0)
	for _, line := range input {
		// split lines
		numberStrings := strings.Split(line, ",")
		update := Update{pages: make([]int, 0)}
		for _, numberString := range numberStrings {
			// parse to integers
			number, err := strconv.Atoi(numberString)
			if err != nil {
				return result, fmt.Errorf("invalid number: %s - %v", line, err)
			}
			update.pages = append(update.pages, number)
		}
		result = append(result, update)
	}

	return result, nil
}

// parseOrdering parses the input strings into a map of integers where the key is the first number and the
// value is a slice of all integers that must be printed later
func parseOrdering(ordering []string) (*Graph, error) {
	graph := newGraph()
	for _, line := range ordering {
		// split lines
		var a, b int
		numberStrings := strings.Split(line, "|")
		if len(numberStrings) != 2 {
			return graph, fmt.Errorf("invalid line: %s", line)
		}
		// parse to integers
		a, err := strconv.Atoi(numberStrings[0])
		if err != nil {
			return graph, fmt.Errorf("invalid first number: %s - %v", line, err)
		}
		b, err = strconv.Atoi(numberStrings[1])
		if err != nil {
			return graph, fmt.Errorf("invalid second number: %s - %v", line, err)
		}
		// check if a is already in the map
		graph.addEdge(a, b)
	}

	return graph, nil
}
