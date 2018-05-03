package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	start := time.Now()
	search = google.First(
		google.Search("replica1"),
		google.Search("replica2"),
	)
	result := search("golang")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)

}
