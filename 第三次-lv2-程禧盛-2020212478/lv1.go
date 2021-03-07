package main



import (
"fmt"
"time"
"math"
)

var flag int

func primeNumber(a float64){
	var i float64
	var j float64
	for j=2;j <= a;j++{
		flag=1
		for i=2;i<math.Sqrt(j);i++{
			if int(j)%int(i)==0{
				flag=0}
		}
		if flag==1{
			fmt.Printf("%d is a primenumber\n",int(j))
		}
	}
}
func main()  {
	var a=50000.0
	go primeNumber(a)
	time.Sleep(1 *time.Second)
}