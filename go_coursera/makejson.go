package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := map[string]interface{}{"name": "", "addr": ""}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("please input a name:\n")
	for scanner.Scan() {
		m["name"] = scanner.Text()
		break
	}
	fmt.Printf("please input an address:\n")
	for scanner.Scan(){
		m["addr"] = scanner.Text()
		break
	}

	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}