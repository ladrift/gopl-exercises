// Exercise 3.13:
// Write const declarations for KB, MB, up through YB as compactly as you can.
package main

import "fmt"

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Println("KB:\t", KB)
	fmt.Println("MB:\t", MB)
	fmt.Println("GB:\t", GB)
	fmt.Println("TB:\t", TB)
	fmt.Println("PB:\t", PB)
	fmt.Println("EB:\t", EB)
	fmt.Println("YB/ZB:\t", YB/ZB)
	//fmt.Println(ZB, YB) // Warning: constant ... overflows int
}
