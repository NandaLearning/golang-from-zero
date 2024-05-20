package golearn

import "fmt"

func SwitchCase() {
	angka := 1

	switch angka {
	case 1:
		fmt.Println("angka ini satu")
	case 2:
		fmt.Println("ini angka dua")
	default:
		fmt.Println("angka tidak di kenal")
	}

}
