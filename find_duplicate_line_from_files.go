package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {

	file_count := len(os.Args[1:])
	if (file_count == 0) {
		fmt.Println("no work to do")
	}

	counts := make(map[string]int)
	for _,file := range os.Args[1:] {
		f,err := os.Open(file)
		if err != nil {
			fmt.Println("open file failed")
			continue
		}
		countLines(f, counts)
		f.Close()
	}
	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n,line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}
