package main

import "fmt"

func main()  {
	var a int
	var b int
	var c int
	fmt.Scanf("%d %d %d",&a,&b,&c)//输入迷宫行数列数和障碍物数
	var maze [a][b] int
	fmt.Scanf("%d %d",&maze[0][0])
	fmt.Scanf("%d %d",&maze[a][b])//输入起点和终点的坐标,写不下去了，改了好几个版本都莫名其妙的过不了编译，最后还是删掉了


}
