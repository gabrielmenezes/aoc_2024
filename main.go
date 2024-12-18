package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func diff(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func openFile(s string) *os.File {
	dat, err := os.Open(s)
	check(err)
	return dat
}

func dayOne() {

	dat := openFile("input.txt")
	Scanner := bufio.NewScanner(dat)
	Scanner.Split(bufio.ScanWords)

	var leftList []int
	var rightList []int
	var totalDistance int

	check := 0
	for Scanner.Scan() {
		if check == 0 {
			number, err := strconv.Atoi(Scanner.Text())
			if err != nil {
				fmt.Println("Error:", err)
			}
			leftList = append(leftList, number)
			check += 1
		} else {
			number, err := strconv.Atoi(Scanner.Text())
			if err != nil {
				fmt.Println("Error:", err)
			}
			rightList = append(rightList, number)
			check -= 1
		}
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	fmt.Printf("Left %v, Right: %v\n", len(leftList), len(rightList))

	for i := range leftList {
		totalDistance += diff(leftList[i], rightList[i])
	}

	fmt.Println("Total Distance:", totalDistance)

	// SIMILARITY SCORE

	var dictSimilarity map[int]int
	var totalSimilarity int = 0

	dictSimilarity = make(map[int]int)

	for i, v := range leftList {
		founds := 0
		for j := range rightList {
			if rightList[j] == leftList[i] {
				founds += 1
			}
		}
		dictSimilarity[v] = founds
		totalSimilarity += v * founds
	}

	fmt.Println("Total Similarity:", totalSimilarity)
}

func isOk(rep []string) bool {
	isOk := true
	j := 0
	for i := 0; i < len(rep)-1; i++ {
		actualValue, err := strconv.Atoi(rep[i])
		check(err)
		nextValue, err := strconv.Atoi(rep[i+1])
		check(err)
		diff := nextValue - actualValue
		if diff > 0 && (j == 0 || j > 0){
			j = 1
		} else if diff < 0 && (j == 0 || j < 0){
			j = -1
		} else {
      isOk = false
      break
    }
		if !(1 <= abs(diff) && abs(diff) <= 3) {
			isOk = false
			break
		}
	}
	return isOk && (j == 1 || j == -1)
}

func abs(diff int) int {
	if diff < 0 {
		return (diff * -1)
	}
	return diff
}

func consider(n int, report_local []string) {
  reportNew := append([]string{}, report_local[:n]...)
  reportNew = append(reportNew, report_local[n+1:]...)
   // fmt.Println(reportNew)
	if isOk(reportNew) {
		anyOk = true
	}
}

var anyOk = false

func dayTwo() {
	dat := func() *os.File {
		dat, err := os.Open("reports.txt")
		check(err)
		return dat
	}()

	Scanner := bufio.NewScanner(dat)
	Scanner.Split(bufio.ScanLines)

	safeReports := 0
	for Scanner.Scan() {
		report := strings.Split(Scanner.Text(), " ")
		consider(0, report)
		if anyOk {
			anyOk = false
			safeReports++
			continue
		}
		for i := 0; i < (len(report) - 1); i++ {
			actualValue, err := strconv.Atoi(report[i])
			check(err)
			nextValue, err := strconv.Atoi(report[i+1])
			check(err)
			diff := nextValue - actualValue
			if (abs(diff) < 1 || abs(diff) > 3) {
				consider(i, report)
				consider(i+1, report)
				break
			}
			if (i+2 < len(report)) {
				actualValue, err := strconv.Atoi(report[i+1])
				check(err)
				nextValue, err := strconv.Atoi(report[i+2])
				check(err)
				diff2 := nextValue - actualValue
				if (diff > 0) != (diff2 > 0) {
					consider(i, report)
					consider(i+1, report)
					consider(i+2, report)
					break
				}
			}
		}
		if anyOk {
			safeReports++
			anyOk = false
		}
	}

	fmt.Println("Total Safe Reports:", safeReports)
}

func dayTree(){

	dat := openFile("text.txt")
	Scanner := bufio.NewScanner(dat)
	Scanner.Split(bufio.ScanWords)
}

func main() {
}
