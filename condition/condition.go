package main

import "fmt"

var score int

func main() {
	myMoney := 100
	if myMoney > 100 {
		fmt.Println("คุณสามารถซื้อได้")
	}else{
		fmt.Println("ขออภัย เงินคุณไม่เพียงพอ")
	}
	fmt.Println("Grade Calculator")
	fmt.Scanf("%d",&score)
	fmt.Println("Score = ", score)
	if score >= 80{
		fmt.Println("You got A Grade")
	}else if score >= 70{
		fmt.Println("You got B Grade")
	}else if score >= 60{
		fmt.Println("You got C Grade")
	}else if score >= 50{
		fmt.Println("You got D Grade")
	}else{
		fmt.Println("You got F Grade")
	}
}