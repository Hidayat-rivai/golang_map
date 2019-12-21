package main

import "fmt"

func main(){
	var menu = map[string]int {"ayam goreng":10000,"nasi goreng":7000,"sate":5000}
	menu["lontong"] = 5000
	fmt.Println(menu)
}