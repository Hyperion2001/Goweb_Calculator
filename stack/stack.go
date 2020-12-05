package stack

import (
	"fmt"
)

type  Stack1 struct{
	N   [1000] float64 //将操作数录入数组
	Top int           //栈顶数用top表示
}
type Stack2 struct{
	N   [1000] rune  //将操作数录入数组
	Top int			 //栈顶数用top表示
}
func isstack1empty(x Stack1) bool{//判断stack1是否为空
	if x.Top == -1{
		return true
	}else{
		return false
	}
}
func Isstack2empty(x Stack2) bool{ //判断stack2是否为空
	if x.Top == -1{
		return true
	}else{
		return false
	}
}
func Push(x *Stack1 , e float64){ //将整数e入栈
	if x.Top == 999 {
		fmt.Println("栈已满")
		return
	}else{
		x.Top++
		x.N[x.Top] = e
		return
	}
}
func Push2(x *Stack2 , e rune){ //将字符e入栈
	if x.Top == 999 {
		fmt.Println("栈已满")
		return
	}else{
		x.Top += 1
		x.N[x.Top] = e
		return
	}
}
func Pop(x *Stack1 , e *float64){ //将栈1顶元素出栈，存到e中
	if x.Top == -1{
		fmt.Printf("栈已空！！\n")
	}else{
		*e = x.N[x.Top]
		x.Top--
	}
}

func Pop2(x *Stack2, e *rune){ //将栈2顶元素出栈，存到e中
	if x.Top == -1{
		fmt.Printf("栈已空！！\n")
	}else{
		*e = x.N[x.Top]
		x.Top--
	}
}
func Inprior(e rune)(x int){ //运算符e在栈内的优先级别
	if e == '+' || e == '-' {
		return 2
	}
	if e == '*' || e == '/' {
		return 4
	}
	if e == '^' {
		return 5
	}
	if e == '(' {
		return 0
	}
	if e == ')' {
		return 7
	}else{
		return -1
	}
}
func Outprior(e rune)(x int){ //运算符e在栈外的优先级别
	if e == '+' || e == '-' {
		return 1
	}
	if e == '*' || e == '/' {
		return 3
	}
	if e == '^' {
		return 6
	}
	if e == '(' {
		return 7
	}
	if e == ')' {
		return 0
	}else{
		return -1
	}
}
func Count( op1 float64 , operand rune , op2 float64 , x *Stack1){ //进行计算并将结果入栈
	var result float64
	if operand == '+' {
		result = op1 + op2
	}
	if operand == '-' {
		result = op1 - op2
	}
	if operand == '*' {
		result = op1 * op2
	}
	if operand == '/' {
		result = op1 / op2
	}
	if operand == '^' {
		result = float64(exponent(int(op1) , int(op2)))
	}
	Push(x , result)
}

func IsSingleDigit(data rune) bool {//判断字符是否为数字
	digit := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, item := range digit {
		if data == item {
			return true
		}
	}
	return false
}

func Isalpha(data rune) bool {
	alpha := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j' , 'k' , 'l' , 'm' , 'n' , 'o' , 'p' , 'q' , 'r' , 's' , 't' , 'u' , 'v' , 'w' , 'x' , 'y' , 'z',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J' , 'K' , 'L' , 'M' , 'N' , 'O' , 'P' , 'Q' , 'R' , 'S' , 'T' , 'U' , 'V' , 'W' , 'X' , 'Y' , 'Z'}
	for _, item := range alpha {
		if data == item {
			return true
		}
	}
	return false
}
func exponent (a, n int) int  {
	result := 1
	for i := n ; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}