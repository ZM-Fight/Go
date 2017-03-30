package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

//全局变量
var (
	numberArr     [4][4]int  //游戏棋盘
	numberArrbool [4][4]bool //表示棋盘是否有数字，true表示有数字，false表示无数字
)

//在棋盘上生成随机数字
func randomNum() {
	randomMax := 0 //记录棋盘空格数目
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if !numberArrbool[i][j] {
				randomMax++
			}
		}
	}
	if randomMax == 0 {
		exit()
	} else {
		newNum := rand.Intn(randomMax) + 1 //新数字产生的位置
		randomMax = 0

		//生成新数字，更新棋盘
	loop:
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if !numberArrbool[i][j] {
					randomMax++
					if newNum == randomMax {
						numberArrbool[i][j] = true
						numberArr[i][j] = 2
						break loop
					}
				}
			}
		}
	}
}

//打印操作结果
func printnum() {
	fmt.Println("- - - - - -")
	for i := 0; i < 4; i++ {
		fmt.Print("- ")
		for j := 0; j < 4; j++ {
			fmt.Print(numberArr[i][j], " ")
		}
		fmt.Println("-")
	}
	fmt.Println("- - - - - -")
}

//计算游戏分数：数组最大值
func mygrade() int {
	grade := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if numberArr[i][j] > grade {
				grade = numberArr[i][j]
			}
		}
	}
	return grade
}

//退出程序
func exit() {
	fmt.Println("游戏结束: ", mygrade(), "分， 你这智商也就告别自行车了！")
	os.Exit(1)
}

func main() {

	//在游戏初始化
	fmt.Println("2048小游戏:\na -- 左移动\nd -- 右移动\nw -- 上移动\ns -- 下移动\nq -- 退出游戏")
	randomNum()
	randomNum()
	fmt.Println("\n\n游戏开始：")
	printnum()

	//游戏过程
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter Command->")
		rawline, _, _ := r.ReadLine()
		direction := string(rawline)

		//处理移动操作
		switch direction {
		case "a", "d", "w", "s": //左,右,上,下
			move(direction)
			randomNum()
			printnum()
		case "q": //退出
			exit()
		default:
			fmt.Println("未知操作，注意游戏说明！")
		}
	}
}

func move(direction string) {
	for index := 0; index < 4; index++ {
		//－>右：每一行 移动第3列,移动第2列,移动第1列
		//－>左：每一行 移动第2列,移动第3列,移动第4列
		//－>下：每一列 移动第3行,移动第2行,移动第1行
		//－>上：每一列 移动第2行,移动第3行,移动第4行
		moveFirst(direction, index)
		moveSecond(direction, index)
		moveThird(direction, index)
	}
}

func moveFirst(direction string, index int) {
	var fromRow, fromColumn, toRow, toColumn int
	switch direction {
	case "a": //左
		fromRow, toRow = index, index
		fromColumn, toColumn = 2-1, 1-1
	case "d": //右
		fromRow, toRow = index, index
		fromColumn, toColumn = 3-1, 4-1
	case "w": //上
		fromColumn, toColumn = index, index
		fromRow, toRow = 2-1, 1-1
	case "s": //下
		fromColumn, toColumn = index, index
		fromRow, toRow = 3-1, 4-1
	}
	switch {
	case numberArr[fromRow][fromColumn] == 0: //位置3空，不动；
	case numberArr[fromRow][fromColumn] != 0 && numberArr[toRow][toColumn] == 0: //位置3不空，位置4空；位置3的数移动到位置4
		numberArr[toRow][toColumn] = numberArr[fromRow][fromColumn]
		numberArr[fromRow][fromColumn] = 0
		numberArrbool[toRow][toColumn] = true
		numberArrbool[fromRow][fromColumn] = false
	case numberArr[fromRow][fromColumn] != numberArr[toRow][toColumn]: //位置3不空，位置4不空；不相同，不动；
	case numberArr[fromRow][fromColumn] == numberArr[toRow][toColumn]: //位置3不空，位置4不空；数字相同，位置3的数移动到位置4，加倍；
		numberArr[toRow][toColumn] *= 2
		numberArr[fromRow][fromColumn] = 0
		numberArrbool[fromRow][fromColumn] = false
	}
}

func moveSecond(direction string, index int) {
	var fromRow, fromColumn, toRow, toColumn int
	switch direction {
	case "a": //左
		fromRow, toRow = index, index
		fromColumn, toColumn = 3-1, 2-1
	case "d": //右
		fromRow, toRow = index, index
		fromColumn, toColumn = 2-1, 3-1
	case "w": //上
		fromColumn, toColumn = index, index
		fromRow, toRow = 3-1, 2-1
	case "s": //下
		fromColumn, toColumn = index, index
		fromRow, toRow = 2-1, 3-1
	}
	switch {
	case numberArr[fromRow][fromColumn] == 0: //位置2空，不动；
	case numberArr[fromRow][fromColumn] != 0 && numberArr[toRow][toColumn] == 0: //位置2不空，位置3空；位置3的数移动到位置4
		numberArr[toRow][toColumn] = numberArr[fromRow][fromColumn]
		numberArr[fromRow][fromColumn] = 0
		numberArrbool[toRow][toColumn] = true
		numberArrbool[fromRow][fromColumn] = false
		moveFirst(direction, index)
	case numberArr[fromRow][fromColumn] != numberArr[toRow][toColumn]: //位置2不空，位置3不空；不相同，不动；
	case numberArr[fromRow][fromColumn] == numberArr[toRow][toColumn]: //位置2不空，位置3不空；数字相同，位置2的数移动到位置3，加倍；
		numberArr[toRow][toColumn] *= 2
		numberArr[fromRow][fromColumn] = 0
		numberArrbool[fromRow][fromColumn] = false
		moveFirst(direction, index)
	}
}

func moveThird(direction string, index int) {
	var fromRow, fromColumn, toRow, toColumn int
	switch direction {
	case "a": //左
		fromRow, toRow = index, index
		fromColumn, toColumn = 4-1, 3-1
	case "d": //右
		fromRow, toRow = index, index
		fromColumn, toColumn = 1-1, 2-1
	case "w": //上
		fromColumn, toColumn = index, index
		fromRow, toRow = 4-1, 3-1
	case "s": //下
		fromColumn, toColumn = index, index
		fromRow, toRow = 1-1, 2-1
	}
	switch {
	case numberArr[fromRow][fromColumn] == 0: //位置1空，不动；
	case numberArr[fromRow][fromColumn] != 0 && numberArr[toRow][toColumn] == 0: //位置1不空，位置2空；位置1的数移动到位置2
		numberArr[toRow][toColumn] = numberArr[fromRow][fromColumn]
		numberArr[fromRow][fromColumn] = 0
		numberArrbool[toRow][toColumn] = true
		numberArrbool[fromRow][fromColumn] = false
		moveSecond(direction, index)
	case numberArr[fromRow][fromColumn] != numberArr[toRow][toColumn]: //位置1不空，位置2不空；不相同，不动；
	case numberArr[fromRow][fromColumn] == numberArr[toRow][toColumn]: //位置1不空，位置2不空；数字相同，位置1的数移动到位置2，加倍；
		numberArr[toRow][toColumn] *= 2
		numberArr[fromRow][fromColumn] = 0
		numberArrbool[fromRow][fromColumn] = false
		moveSecond(direction, index)
	}
}
