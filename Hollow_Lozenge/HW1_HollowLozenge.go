package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	hollowLozenge(n)
	fmt.Println()
	fmt.Println()
	fmt.Println("-----------------------------------------")

	fmt.Printf("That's a %s * %s hollow lozenge! :))", strconv.Itoa(n), strconv.Itoa(n))
}
func hollowLozenge(n int) {
	var spaceLeft = 1
	var starLeft = n
	var starRight = n - 1
	var spaceRight = 1
	var down_starLeft = 2
	var down_spaceLeft = n - 1
	var down_starRight = 2
	var down_spaceRight = n - 4
	fmt.Print(" ")
	for i := 0; i < 2*n-2; i++ {
		fmt.Print("*")
	}

	fmt.Print(" ")
	fmt.Println()
	if n%2 == 1 || n%2 == 0 {
		for (spaceLeft+starLeft == n+1 && spaceLeft >= 1 && starLeft >= 1) && (spaceRight+starRight <= n && spaceRight <= n-2 && starRight >= 1) {
			for i := 1; i < starLeft; i++ {
				fmt.Print("*")
			}
			for i := 0; i < spaceLeft; i++ {
				fmt.Print(" ")
			}
			for i := n; i > starRight; i-- {
				fmt.Print(" ")
			}
			for i := 0; i < starRight; i++ {
				fmt.Print("*")

			}
			fmt.Println()
			spaceRight++
			starRight--
			starLeft--
			spaceLeft++
		}
		fmt.Print("*")
		for i := 0; i < 2*n-2; i++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		fmt.Println()
		for (spaceLeft+starLeft == n+1 && spaceLeft >= 1 && starLeft >= 1) && (spaceRight+starRight <= n && spaceRight <= n-2 && starRight >= 1) {
			for i := 1; i < starLeft; i++ {
				fmt.Print("*")
			}
			for i := 0; i < spaceLeft; i++ {
				fmt.Print(" ")
			}
			for i := n; i > starRight; i-- {
				fmt.Print(" ")
			}
			for i := 0; i < starRight; i++ {
				fmt.Print("*")
			}
			fmt.Println()
			spaceRight++
			starRight--
			starLeft--
			spaceLeft++
		}

	}
	for down_starLeft <= n-1 && down_spaceLeft >= 0 {
		for i := 0; i < down_starLeft; i++ {
			fmt.Print("*")
		}
		for i := down_spaceLeft; i >= 1; i-- {
			fmt.Print(" ")
		}
		for i := down_spaceRight; i >= 0; i-- {
			fmt.Print(" ")
		}
		for i := 0; i < down_starRight; i++ {
			fmt.Print("*")
		}
		fmt.Println()
		down_spaceLeft--
		down_starLeft++
		down_starRight++
		down_spaceRight--
	}
	fmt.Print(" ")
	for i := 0; i < 2*n-2; i++ {
		fmt.Print("*")
	}
	fmt.Print(" ")

}
