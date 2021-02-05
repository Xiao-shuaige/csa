package main

import (
	"fmt"
)

func main()  {
	var a,b int
	var c byte               //声明两个Int类型的变量和一个字符，代表输入计算器中的两个数和运算符
	fmt.Scanf("%d%c%d",&a,&c,&b)   //用scan方式读入
	calculator(a,b,c)//调用函数
}
func calculator( a int, b int ,c byte)  {
	switch c {//判断运算符
	case '+':
		fmt.Println(a+b)
	case '-':
		fmt.Println(a-b)
	case '*':
		fmt.Println(a*b)
	case '/':
		fmt.Println(a/b)
	case '%':
		fmt.Println(a%b)
	default:
		fmt.Println("can not identify the character")//打印不能识别运算的提示语句
	}
}

