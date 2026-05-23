package main
import "fmt"
func main() {
	score := 0
	avg_sum:=0.0
for j:=0;j<3;j++{
	avg:=0.0
	sum:=0
  for i:=0;i<5;i++{
	
	fmt.Printf("请输入第%d班第%d个学生的成绩：",j+1,i+1)
	fmt.Scanln(&score)
	sum+=score
  }
  avg=float64(sum)/5
  fmt.Printf("第%d班平均分是：%v\n",j+1,avg)
  avg_sum+=avg
}
fmt.Println("总平均分是：",avg_sum/3)
}