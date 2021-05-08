package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	int_sli := make([]int, 0, 3)
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Please inupt your numbers:")
	for reader.Scan() {
		line := reader.Text()
		if line != "x" {
			line, _ := strconv.Atoi(line)
			int_sli = append(int_sli, line)
			sort.Ints(int_sli)
			fmt.Println(int_sli)
		} else {
			break
		}

	}
}
