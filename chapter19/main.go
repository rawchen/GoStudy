package main

import "fmt"

func main() {
	// 数组
	var hens [6]float64
	hens[0] = 1.0
	hens[1] = 5.0
	hens[2] = 2.4
	hens[3] = 40.4
	hens[4] = 3.4
	hens[5] = 3.4

	totalWeight := 0.0
	for i := range hens {
		totalWeight += hens[i]
	}
	// 平均值
	fmt.Printf("%.2f \n", totalWeight/float64(len(hens)))

	// hens的地址为 =0xc00000a420
	// hens[0]地址为=0xc00000a420
	// hens[1]地址为=0xc00000a428
	// hens[2]地址为=0xc00000a430
	fmt.Printf("hens的地址为=%p, hens[0]地址为=%p, hens[1]地址为=%p, hens[2]地址为=%p",
		&hens, &hens[0], &hens[1], &hens[2])
}
