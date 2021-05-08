package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fname string
	fmt.Println("Enter name of the file")
	fmt.Scan(&fname)
	f, err := os.Open(fname)
	arr := make([]Name, 0, 10)
	if err == nil {
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			name := scanner.Text()
			fname := strings.Split(name, " ")[0]
			lname := strings.Split(name, " ")[1]
			t := Name{fname, lname}
			arr = append(arr, t)
		}
		for i, v := range arr {
			fmt.Println(strconv.Itoa(i+1) + ". First Name: " + v.fname + ", Last Name: " + v.lname)
		}
		f.Close()
	}
}
