package main
import "fmt"
func main() {
	var name string
	var age int
	var salary float32
	var ispass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪资：")
	fmt.Scanln(&salary)
	fmt.Println("请输入是否通过：")
	fmt.Scanln(&ispass)
	fmt.Printf("姓名；%s,年龄；%d,工资：%f,是否通过：%t",name,age,salary,ispass)
}