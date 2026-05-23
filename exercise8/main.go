
package main
import (
	"fmt"
)
func swap(a *int,b *int){
	tem:=*a
	*a=*b
	*b=tem
}
func main() {
	var a,b int
	fmt.Println("请输入两个数字(两个数字之间用空格隔开100),交换后输出：")
	fmt.Scanf("%d %d",&a,&b)
	fmt.Println("交换前的两个数字为：",a,b)
	swap(&a,&b)
	fmt.Println("交换后的两个数字为：",a,b)
}
