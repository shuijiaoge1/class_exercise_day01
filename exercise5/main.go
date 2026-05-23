package main
import "fmt"
var total int
func main() { 
//打印金字塔	
fmt.Println("请输入金字塔的层数：")
fmt.Scanln(&total)
for i:=1;i<=total;i++{ 
	for k:=1;k<=total-i;k++{
		fmt.Printf(" ")
	}
	for j:=1;j<=2*i-1;j++{
		fmt.Print("*")
	}
	fmt.Println()
}
//打印空心金字塔
println("-------------------------------------")
for i:=1;i<=total;i++{ 
	for k:=1;k<=total-i;k++{
		fmt.Printf(" ")
	}
	for j:=1;j<=2*i-1;j++{
		if j==1||j==2*i-1||i==total{
			fmt.Print("*")
		}else{
			fmt.Print(" ")
		}
		
	}
	fmt.Println()
}

}