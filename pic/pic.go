package main

import "golang.org/x/tour/pic"

func printPic(dx, dy int) [][]uint8 { // is printing a bunch of characters. Look at it in a smaller window.
	var result = make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		var k uint8 = 1
		result[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			var l uint8 = 1
			result[i][j] = average(k, l)
			l++
		}
		k++
	}
	return result

}

func average(x, y uint8) uint8 {
	return (x + y) / 2.0
}

func multiply(x, y int) int {
	return x * y
}

func power(x, y int) int {
	return x ^ y
}

func main() {
	pic.Show(printPic)
}
