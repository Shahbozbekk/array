// Напишите программу, принимающая на вход число
// N(N≥4), N целых чисел-элементов среза.
// На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
// Sample Input:
// 5
// 41 -231 24 49 6
// Sample Output:
// 49

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a >= 4 {
		b := make([]int, a)

		for i := 0; i < a; i++ {
			fmt.Scan(&b[i])

		}

		fmt.Println(b[3])
	}
}
