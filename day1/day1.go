package main

import (
	"fmt"
)

func main(){

	//1st task
	fmt.Println("Hello World")

	//2nd task

	for k:=1;k<=5;k++{
		for j:=1;j<=k;j++{
			fmt.Print("*")
		}
		fmt.Println(" ")
	}

	//3rd task
	i:=1
	for k:=1;k<=5;k++{
		for j:=1;j<=k;j++{
			fmt.Print(i," ")
			i++
		}
		fmt.Println(" ")
	}

    //factorial number
    num:=5
    factorial(num)

    //prime number
    num=5
    primenum(num)


}

func primenum(num int)  {
	flag:=1
	for i:=2;i<num;i++{
		if (num % i ) == 0{
			flag = 0
		}


	}
	if (flag == 1){
		fmt.Println(num ," is a prime NUMBER ")
	}else{
		fmt.Println(num," Number is not prime")
	}
}

func factorial(num int){
	fact:=1
	if (num == 0) {
		fmt.Println("factorial of 0 is 0")
	} else if (num == 1){
		fmt.Println("factorial of 1 is 1")
	}else{
		for i:=1;i<=num;i++{
			fact = fact*i;
		}
		fmt.Println("factorial of ",num," is ",fact)
	}


}

