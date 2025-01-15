package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机数生成器
/*
游戏规则：
1. 点击A，A有几率获取0-2个灯，B有几率熄灭1个灯
2. 点击B，B有几率获取0-2个灯，A有几率熄灭1个灯
3. 当a和b都获取7盏灯时，游戏结束
4. 当A或B获取灯的数量大于7时，则A或B全部熄灭，游戏结束
*/

var A = 0
var B = 0

func main() {
	// 初始化灯
	for i := 0; i < 10000000; i++ {
		user := getRandom()
		// 判断A和B的灯数量
		// 当A为7时，则只会选B，反之亦然
		if A == 7 {
			user = 1
		}
		if B == 7 {
			user = 0
		}
		if A == 7 && B == 7 {
			fmt.Println(A, B, i)
			A = 0
			B = 0
		}
		if user == 0 {
			selectA()
		}
		if user == 1 {
			selectB()
		}
	}
}

// 随机选择A或B
func getRandom() int {
	rand.NewSource(time.Now().UnixNano())
	randomIndex := rand.Intn(2) // 生成0到1之间的随机数
	// 定义一个映射函数，将随机数映射为A或B
	return randomIndex
}

// 选择A
func selectA() {
	number := getLight()
	if number == -1 {
		B = B - 1
		if B < 0 {
			B = 0
		}
	} else {
		A = A + number
		if A > 7 {
			A = 0
		}
	}
}

// 选择B
func selectB() {
	number := getLight()
	if number == -1 {
		B = B - 1
		if B < 0 {
			B = 0
		}
	} else {
		B = B + number
		if B > 7 {
			B = 0
		}
	}
}

// 点灯
func getLight() int {

	// 定义一个权重数组，表示A和B的点灯概率
	weights := []int{2, 1, 3, 4}
	randomIndex := weightedRandom(weights)
	// 假设我们有一个与权重对应的值数组 values，可以根据 randomIndex 来获取值
	values := []int{-1, 0, 1, 2}
	// 这里我们仅打印随机索引

	return values[randomIndex]
}

// 权重选择
func weightedRandom(weights []int) int {
	// 计算权重总和
	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}

	// 选取一个随机数
	randomValue := rand.Intn(totalWeight) + 1 //1-6

	// 遍历权重数组，根据权重分布来确定返回的索引
	weightSum := 0
	for index, weight := range weights {
		weightSum += weight //weightSum = 0+3+1+2   第一个元素权重为3，第二个为1，第三个为2。weightSum = 3,4,6
		if randomValue <= weightSum {
			return index
		}
	}

	// 如果因为某些原因没有返回，则此处返回一个错误值（理论上不应该发生）
	return -1
}
