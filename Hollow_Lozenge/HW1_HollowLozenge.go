package main

import (
	utill2 "SW2025_HW1/Hollow_Lozenge/utill"
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
	var downStarLeft = 2
	var downSpaceLeft = n - 1
	var downStarRight = 2
	var downSpaceRight = n - 4
	fmt.Print(" ")
	for i := 0; i < 2*n-2; i++ {
		fmt.Print("*")
	}
	fmt.Print(" ")
	fmt.Println()
	if n%2 == 1 || n%2 == 0 {
		for (spaceLeft+starLeft == n+1 && spaceLeft >= 1 && starLeft >= 1) && (spaceRight+starRight <= n && spaceRight <= n-2 && starRight >= 1) {
			utill2.PrintLeftStar(starLeft)
			utill2.PrintLeftSpace(spaceLeft)
			utill2.PrintRightSpace(starRight, n)
			utill2.PrintRightStar(starRight)
			fmt.Println()
			spaceRight++
			starRight--
			starLeft--
			spaceLeft++
		}
		fmt.Print("*")
		utill2.PrintLeftSpace(2*n - 2)
		fmt.Print("*")
		fmt.Println()
		for (spaceLeft+starLeft == n+1 && spaceLeft >= 1 && starLeft >= 1) && (spaceRight+starRight <= n && spaceRight <= n-2 && starRight >= 1) {
			utill2.PrintLeftStar(starLeft)
			utill2.PrintLeftSpace(spaceLeft)
			utill2.PrintRightSpace(starRight, n)
			utill2.PrintRightStar(starRight)
			fmt.Println()
			spaceRight++
			starRight--
			starLeft--
			spaceLeft++
		}

	}
	for downStarLeft <= n-1 && downSpaceLeft >= 0 {
		utill2.PrintRightStar(downStarLeft)
		utill2.PrintDownLeftSpace(downSpaceLeft)
		utill2.PrintDownRightSpace(downSpaceRight)
		utill2.PrintRightStar(downStarRight)
		fmt.Println()
		downSpaceLeft--
		downStarLeft++
		downStarRight++
		downSpaceRight--
	}
	fmt.Print(" ")
	utill2.PrintRightStar(2*n - 2)
	fmt.Print(" ")
}
