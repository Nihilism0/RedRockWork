package main

import "fmt"

// calculate函数通过cmd值进行switch分支运算 返回结果
func calculate(a int, cmd byte, b int) int {
	var wow int
	switch cmd {
	case '+':
		wow = a + b
	case '-':
		wow = -+b
	case '*':
		wow = a * b
	case '/':
		wow = a / b
	}
	return wow
}
func main() {
	var a, b int
	var cmd byte
	var num int
	var tool byte
	numbers := make([]int, 0, 20) //分配内存给一个切片,用来储存每次计算后的num值
	//循环进行
	for {
		fmt.Scanf("%d%c%d", &a, &cmd, &b)
		fmt.Scanf("%c", &tool)
		num = calculate(a, cmd, b)
		numbers = append(numbers, num)
		for _, k := range numbers {
			fmt.Printf("%d ", k)
		}
		fmt.Println() //愚蠢的进行两次换行
		fmt.Println()
	}
}
