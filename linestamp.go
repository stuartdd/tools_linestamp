// predate
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	addLineNumber := flag.Bool("n", false, "Prefix lines with a 4 digit line number")
	flag.Parse()

	count := 1
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if *addLineNumber {
			if count < 10 {
				fmt.Printf("000%d:", count)
			} else {
				if count < 100 {
					fmt.Printf("00%d:", count)
				} else {
					if count < 1000 {
						fmt.Printf("0%d:", count)
					} else {
						fmt.Printf("%d:", count)
					}
				}
			}
			count++
		}
		fmt.Print(time.Now().Format("2006.01.02 15:04:05.000 "))
		fmt.Println(line)

	}
}
