package main

import "fmt"

func main() {
	/**
		You get error:
	panic: assignment to entry in nil map
	goroutine 1 [running]:
	main.main()
	        G:/HCL/2026/march/mapsapp/main_mapWithNilAssignemnt.go:7 +0x28
	exit status 2
	*/
	var m map[string]int
	m["1"] = 1
	fmt.Println(m)
}
