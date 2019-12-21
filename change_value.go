package main

import "fmt"

func main(){
	var menu = map[string]int {"ayam goreng":10000,"nasi goreng":7000,"sate":5000}
	menu["sate"] = 14000
	fmt.Println(menu)
}