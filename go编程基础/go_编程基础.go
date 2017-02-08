// go 编程基础
/*:多行注释  */
/*
go   注释方法：
// : 单行注释
go 的内置关键字（25个均为小写）
break default func interface select case defer go map struct chan else goto package switch
const fallthrough  if range type continue for import return var

go 程序的一般结构：
go 程序是通过package 来组织的（与python类似）
只有package 名称为main 的包，可以包含main 函数
一个可执行程序有且仅有一个main包
通过import 关键字来导入其他非main 包
通过const 关键字来进行常量的定义
通过在函数体外部使用 var 关键字来进行全局变量的声明与赋值
通过type 关键字来进行结构struct 或接口（interface）的声明
通过func 关键字来进行函数的声明

 */


package main       //  当前程序的包名
import "fmt"       //导入fmt包
const PI  = 3.14    //常量的定义
var name = "gopher"    //全局变量的声明与赋值
type newtype  int       //一般类型的声明
type gopher struct {}       //结构的声明
type golang interface {}     //接口的声明
//由main 函数作为程序入口点启动
func main() {
	fmt.Println("hello go ")

}


//go 可见性规则：
/*go 语言中，使用大小写来决定该常量，变量，类型，接口，结构，或函数是否可以被外部包所调用：
根据约定，函数名首字母小写即为private
函数名首字母大写即为 public

*/

/*
指针
go 虽然保留了指针，但与其他编程语言不同的是，在go当中，不支持指针运算以及"->"运算符，而直接采用"." 选择符来操作指针目标对象的成员
操作符 "& " 取变量地址，使用"*" 通过指针间接访问目标对象
默认值为 nil ，而非NULL

递增递减语句
在go 中，++与-- 是作为语句而并不是作为表达式
 */


