package main
import(
	"fmt"
	"math"
)
func main(){
	var a,b,c float64
	fmt.Println("请输入a,b,c（用空格分隔）:")
	fmt.Scan(&a, &b, &c)
	d := b*b - 4*a*c
	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println("方程有两个解:", x1, x2)
	} else if d == 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		fmt.Println("x1=x2=", x1)
	} else {
		fmt.Println("方程无解")
	}


}