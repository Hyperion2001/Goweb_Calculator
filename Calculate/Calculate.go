package Calculate

import (
	"fmt"
	"goweb计算器/stack"
	"strconv"
)

func Calculate(expression string) (result float64) {
	var stack1 stack.Stack1
	var stack2 stack.Stack2
	var length = len(expression)
	var g, j, b0 int
	var a, b float64
	var i = 0
	var temp string
	var operate rune
	stack1.Top = -1 //将栈初始化
	stack2.Top = -1 //将栈初始化
	for true {
		temp = ""
		g = 0
		if stack.IsSingleDigit(rune(expression[i])) { //确保读入要运算的了全部的数字字符
			j = 0
			g = 1
			for stack.IsSingleDigit(rune(expression[i])) {
				temp += string(expression[i])
				j++
				i++
				if i >= length {
					break
				}
			}
			if i < length && (expression[i] != '-' && expression[i] != '+' && expression[i] != '*' && expression[i] != '/' && expression[i] != '^' && expression[i] != ')' && expression[i] != '%') {
				fmt.Printf("输入有误1\n") //运算符后面不是数字，错误
				return 0
			}
			//temp[j] = '\0'
			b0, _ = strconv.Atoi(temp) //将字符串所存储对应的整形数据赋值给b0
			b = float64(b0)            //转化为浮点形式b
			fmt.Printf("%f\n", b)
			stack.Push(&stack1, b)
			fmt.Printf("栈顶元素：%f\n", stack1.N[stack1.Top])
			if i >= length {
				break
			}
		} else {
			if expression[i] == '-' || expression[i] == '+' || expression[i] == '*' || expression[i] == '/' || expression[i] == '^' || expression[i] == '(' || expression[i] == ')' || expression[i] == '%' {
				g = 1
				//if (expression[i] == '%' && i == 1) || (stack.IsSingleDigit(rune(expression[i + 1]))) {
				//	fmt.Printf("输入有误2\n") //百分号放在最前面或百分号后面是数字，错误
				//	return 0
				//}
				if expression[i] == '(' && i == length {
					fmt.Printf("输入有误3\n") //前括号放在最后面，错误
					return 0
				}
				if expression[i] == '-' || expression[i] == '+' || expression[i] == '*' || expression[i] == '/' || expression[i] == '^' {
					if i == length {
						fmt.Printf("输入有误4\n") //运算符放在表达式最后，错误
						return 0
					}
					if (!stack.IsSingleDigit(rune(expression[i+1]))) && (!stack.Isalpha(rune(expression[i+1]))) && expression[i+1] != '(' { //“+、-、*、/、^”后面跟的不是数字或者变量，错误
						fmt.Printf("输入有误5\n")
						return 0
					}
				}
				if expression[i] == '-' && (i == 0 || expression[i-1] == '(') { //处理表达式开头的负号
					stack.Push(&stack1, 0)
					stack.Push2(&stack2, rune(expression[i]))
					i++
				}
				if expression[i] == '%' { //处理百分号
					const percent = 0.0100000000000000000000000
					const multiple = '*'
					stack.Pop(&stack1, &a)
					stack.Count(percent, multiple, a, &stack1)
					if i == length-1 {
						break
					} else {
						i++
					}
				} else {
					fmt.Print(stack2.Top)
					fmt.Printf("\n")

					if stack.Isstack2empty(stack2) || (stack.Inprior(stack2.N[stack2.Top]) < stack.Outprior(rune(expression[i]))) {
						stack.Push2(&stack2, rune(expression[i]))
						fmt.Printf("%d\n", stack2.Top)
						i++
					} else if stack.Inprior(stack2.N[stack2.Top]) == stack.Outprior(rune(expression[i])) {
						i++
						stack2.Top--
					} else if stack.Inprior(stack2.N[stack2.Top]) > stack.Outprior(rune(expression[i])) {
						stack.Pop(&stack1, &a)
						stack.Pop(&stack1, &b)
						stack.Pop2(&stack2, &operate)
						stack.Count(b, operate, a, &stack1)
					}
					if i >= length {
						break
					}
				}
			}
		}
		if g == 0 { //g = 0表示该字符不符合要求
			fmt.Printf("输入有误6\n")
			return 0
		}
	}
	for stack2.Top != -1 {
		stack.Pop(&stack1, &a)
		stack.Pop(&stack1, &b)
		stack.Pop2(&stack2, &operate)
		if operate == '(' || operate == ')' { //括号多余的情况
			fmt.Printf("输入有误7\n")
			return 0
		}
		fmt.Printf("a:%f  b:%f\n", a, b)
		//fmt.Printf("Count之前:%d\n",stack1.Top)
		stack.Count(b, operate, a, &stack1)
		//fmt.Printf("Count之后:%d\n",stack1.Top)
	}
	fmt.Printf("%.20f\n", stack1.N[stack1.Top])
	return stack1.N[stack1.Top]
}
