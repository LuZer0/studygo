package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	fname_list := make([]string, 0)
	lname_list := make([]string, 0)
	dat, err := os.Open("test.txt")
	if err != nil{
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer dat.Close()
	br := bufio.NewReader(dat)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF{
			break
		}
		fname := strings.Split(string(a), " ")[0]
		lname := strings.Split(string(a), " ")[1]
		fname_list = append(fname_list, fname)
		lname_list = append(lname_list, lname)
	}
	i := 0
	for i < len(fname_list){
		fmt.Printf("first name: %v, last name: %v\n", fname_list[i],lname_list[i])
		//fmt.Println(lname_list[i])
		i++
	}

}
