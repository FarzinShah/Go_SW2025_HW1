package utill

import "fmt"

// اگه اول اسم تابعمون رو بزرگ بنویسیم انگار استاتیک و exportable هست

func PrintRightStar(starNums int) {
	for i := 0; i < starNums; i++ {
		fmt.Print("*")
	}
}
func PrintRightSpace(spaceNums int, n int) {
	for i := n; i > spaceNums; i-- {
		fmt.Print(" ")
	}
}
func PrintLeftStar(starNums int) {
	for i := 1; i < starNums; i++ {
		fmt.Print("*")
	}
}
func PrintLeftSpace(spaceNums int) {
	for i := 0; i < spaceNums; i++ {
		fmt.Print(" ")
	}

}
func PrintDownRightSpace(downSpaceRight int) {
	for i := downSpaceRight; i >= 0; i-- {
		fmt.Print(" ")
	}
}
func PrintDownLeftSpace(downSpaceLeft int) {
	for i := downSpaceLeft; i >= 1; i-- {
		fmt.Print(" ")
	}
}
