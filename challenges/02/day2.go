package _2

import (
	"adventofcode2024/challenges/util"
)

func redNosedReports(filename string) (int, error) {
	data, err := util.ReadSpaceSeparatedData(filename)
	if err != nil {
		return 0, err
	}

	return checkReports(data)
}

/*
 * checkReports iterates through all reports and counts how many are considered safe.
 */
func checkReports(data [][]int) (int, error) {
	safeReports := 0
	for _, line := range data {
		if isReportSafe(line) {
			safeReports++
		}
	}
	return safeReports, nil
}

/*
 * So, a report only counts as safe if both of the following are true:
 *
 * - The levels are either all increasing or all decreasing.
 * - Any two adjacent levels differ by at least one and at most three.
 */
func isReportSafe(report []int) bool {
	// compare first and last element to determine if the list is increasing or decreasing
	incr := isIncreasing(report)
	safe, place := isSafe(report, incr)
	if safe {
		return true
	}
	// new array with the first element removed
	shortenedReport1 := remove(report, place)

	// new array with the second element removed
	shortenedReport2 := remove(report, place+1)

	safe1, _ := isSafe(shortenedReport1, incr)
	safe2, _ := isSafe(shortenedReport2, incr)
	return safe1 || safe2
}

func remove(report []int, i int) []int {
	shortenedReport := make([]int, len(report)-1)
	copy(shortenedReport, report[:i])
	copy(shortenedReport[i:], report[i+1:])
	return shortenedReport
}

func isIncreasing(report []int) bool {
	inc, dec := 0, 0
	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			inc++
		} else if report[i] > report[i+1] {
			dec++
		}
	}
	return inc > dec
}

func isSafe(report []int, isIncreasing bool) (bool, int) {
	prevLevel := report[0]
	for i, level := range report[1:] {
		if (isIncreasing && prevLevel > level) || (!isIncreasing && prevLevel < level) {
			return false, i
		}
		diff := level - prevLevel
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false, i
		}
		prevLevel = level
	}
	return true, 0
}
