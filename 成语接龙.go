package main

import "fmt"

func main()  {
	var words [5] string//声明一个5*5的数组，用来储存20条内不超过20个字母的单词
	words [4]="nil"
	fmt.Println("请输入4串单词")
	for  i:=0;i<4;i++ {
		fmt.Scanf("%s",&words[i])///输入单词
	}
	var firstLetter[5] byte
	for i:=0;i<5;i++{
		fmt.Scanf("%c",&firstLetter[i])//输入五个单词的首字母
	}
	var letter byte
	fmt.Println("请输入首字母")//输入一个字母
	fmt.Scanf(" %c",&letter)
    for i:=0;i<4;i++{
    	if firstLetter[i]==letter {//比较之后进行接龙
    		for j:=i;j<4;j++{
    			fmt.Printf("%s",words[j])
			}
		}
	}
}
