package main
import "fmt"
//求三个数最大值
func main() { 
	var a int=10
	var b int =20
	var c int =30
	var max =a
	if b>max{
		max=b
	}
	if c>max{
		max=c
	}
	fmt.Println("max=",max)

	 
}