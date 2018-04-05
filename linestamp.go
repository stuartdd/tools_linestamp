// predate
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strings"
	"container/list"
)

const DIV_NANO float64 = 1000000000.0

func main() {
	var testTimeAcc int64 = 0
	var testCount int32 = 0;
	startTime := time.Now().UnixNano()
	testTime := startTime
	lineTime := startTime
	foundTest := false
	foundTestClass := false
	summary := list.New()
	testLine := ""
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		lineTime = time.Now().UnixNano()
		indexTest := strings.Index(line, ":test")
		indexTestClass := strings.Index(line, ":testClasses")
		if indexTestClass > 0 {
			fmt.Print(" CLASS    ")
			testTime = lineTime
			foundTestClass = true
			foundTest = false
		} else {
			if (indexTest > 0) && foundTestClass {
				testCount++
				fmt.Printf(" TEST %3d ", testCount)
				foundTest = true
				testLine = line
				foundTestClass = false
			} else {
				if foundTest {
					elapsed := lineTime - testTime
					testTimeAcc = testTimeAcc + elapsed
					sum := fmt.Sprintf(" %007.3f  ", float64(elapsed) / DIV_NANO)
					summary.PushBack(sum + testLine)
					fmt.Print(sum)
					foundTest = false
					foundTestClass = false
				} else {
					fmt.Print("          ")
				}
			}
		}
		fmt.Print(time.Now().Format("2006.01.02 15:04:05.000 "))
		fmt.Println(line)
	}
	for e := summary.Front(); e != nil; e = e.Next() {
		fmt.Printf("        Test Time:%s\n", e.Value)
	}
	fmt.Printf("Total - Test Time: %007.3f\n", float64(testTimeAcc)  / DIV_NANO)
	fmt.Printf("Total - Test Count:%3d\n", testCount)
	fmt.Printf("Total -  Run Time: %007.3f :All times are in seconds.\n", float64(time.Now().UnixNano() - startTime) / DIV_NANO)
}
