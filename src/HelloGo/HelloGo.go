package main

import (
    "fmt"
    "net/http"
)

type dollars float32

func (d dollars) String() string {
    return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}

func main() {
    account := Account{FirstName:"tao",LastName:"ranqing"}
	fmt.Println(account)
	account.ChangeName("lin","junjie")
	fmt.Println(account)
	employee := Employee{account:account}
	fmt.Println(employee.store)
	employee.AddCredits(100)
	fmt.Println(employee.store)
}

func PrintFizzBuzz(){
	for i := 1 ; i < 101 ; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0 :
			fmt.Println("FizzBuzz")
		case i % 3 == 0 :
			fmt.Println("Fizz")
		case i % 5 == 0 :
			fmt.Println("Buzz")
		default :
			fmt.Println(i)
		}
	}
}

func FibonacciSequence(num int)(res []int) {

	if num < 2 {
		res = make([] int, 0)
		return
	}

	res = make([]int,num)
	res[0] = 1
	res[1] = 1
	for i := 2;i<num;i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return
}

func ParseRomanNum(str string)(res int){

	dic := map[rune]int {
		'M' : 1000,
		'D' : 500,
		'C' : 100,
		'L' : 50,
		'X' : 10,
		'V' : 5,
		'I' : 1,
	}

	numArr := make([]int,len(str) + 1)

	for index,digit := range str{

		if num,exist := dic[digit];exist{

			numArr[index] = num;
		}else {
			res = 0
			return
		}
	}

	for index := 0;index < len(numArr) - 1; index++{

		if(numArr[index] > numArr[index + 1]){

			res += numArr[index]
		}else{
			res -= numArr[index]
		}
	}
	
	return
}