//递归练习
package main
import "fmt"
//计算一个数对应的斐波那契数
func fbs(n int)int{
	if n==1||n==2{
		return 1
	}else{
		return fbs(n-1)+fbs(n-2)
	}

}
//计算函数值
func fx(n int)int{
	if n==1{
		return 3
	}else{
	return 2*fx(n-1)+1	
	}
}
//猴子吃桃问题
func count1(n int)int{
	if n==10{
		return 1
	}else{
		return (count1(n+1)+1)*2
	}
}
func main() {
  var num int
  fmt.Println("请输入数字(计算得到其对应斐波那契数)：")
  fmt.Scanln(&num)
  f:=fbs(num)
  fmt.Println(f)
  fmt.Println("----------------------------")
  var num1 int
  fmt.Println("请输入数字(计算得到其函数值)：")
  fmt.Scanln(&num1)
  f1:=fx(num1)
  fmt.Println(f1)
  fmt.Println("----------------------------")
  var num2 int
  fmt.Println("请输入天数(计算得到其桃子的数量)：")
  fmt.Scanln(&num2)
  f2:=count1(num2)
  fmt.Println(f2)
}