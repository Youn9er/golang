package main

import "fmt"

var aa = 11
var ss = "string"	//可定义全局变量
var bb = true

//或者

var (	//省略多余的var关键字 使用()代替
	aaa = 11
	sss = "Young"
	bbb= false
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)	//println 打印变量名；printf 打印变量名和变量类型； q%  单引号围绕的字符字面值，由Go语法安全地转义 可以打印出引号
}

func variableInitialValue() {
	var a int = 5	//赋初值
	var b, c int = 6, 7		//变量一旦被定义就一定要被使用
	var s string = "abc"
	fmt.Println(a, s, b, c)
}

func variableTypeDeduction() {
	var a, b, c, s = 1, 2, true, "abc"
	fmt.Println(a, b, c, s)
}

func variableSorter() {
	a, b, c, s := 6, 7, false, "def"	// := 定义变量  :=只能在函数内使用

	b = 10	//继续赋值
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()		//0 ""
	variableInitialValue()	//5 abc 6 7
	variableTypeDeduction()	//1 2 true abc 可以省略变量类型
	variableSorter()	//6 7 false def 可以省略var关键字

	fmt.Println(aaa, sss, bbb)	//11 Young false定义包package内变量（非全局变量）
}
