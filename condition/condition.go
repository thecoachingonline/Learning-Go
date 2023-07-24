package main

import "fmt"

func main() {
	myMoney := 100
	if myMoney > 100 {
		fmt.Println("คุณสามารถซื้อได้")
	}else{
		fmt.Println("ขออภัย เงินคุณไม่เพียงพอ")
	}
}