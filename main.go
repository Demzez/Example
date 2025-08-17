package main

import "fmt"

func calculation(start *int) int {
	for *start <= 100 {
		fmt.Println(*start)
		*start++
	}
	return *start
}
func UpdateV0_2() string {
	fmt.Println("successfully updated v0.2!")
	return "End/"
}

func main() {
	var start int
	for {
		fmt.Println("Input your start: ")
		_, err := fmt.Scanln(&start)
		if err == nil {
			break
		}
		fmt.Println("Некоректный ввод.")
	}
	calculation(&start)
	fmt.Println(fmt.Sprintf("Your finish value: %d", start))
	fmt.Println(UpdateV0_2())
}
