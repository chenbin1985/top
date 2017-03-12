package top

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// MINVALUE 数据集中的最小值 (32位)
const MINVALUE = -(1<<32 - 1)

// 计算离n最近2的x次方值
func near2Pow(n int, greater bool) int {
	l := math.Log2(float64(n))
	i := math.Floor(l)
	if l == i {
		return n
	}
	if greater {
		return int(math.Pow(2, i+1))
	}
	return int(math.Pow(2, i))
}

func calcRoundArrayLength(nRound int) int {
	var count int
	for i := nRound - 1; i >= 0; i-- {
		count += int(math.Pow(2, float64(i)))
	}
	return count
}

func calcRoundSlice(round, maxRound int, roundArray []int) []int {
	var start, end int
	for i := 0; i < round; i++ {
		maxRound--
		start += int(math.Pow(2, float64(maxRound)))
	}
	maxRound--
	end = start + int(math.Pow(2, float64(maxRound)))
	return roundArray[start:end:end]
}

// 根据数据个数count, 生成数组化的计算树
func createRounds(count int) [][]int {
	extCount := near2Pow(count, true)
	maxRound := int(math.Log2(float64(extCount))) + 1

	roundArrayLength := calcRoundArrayLength(maxRound)
	roundArray := make([]int, roundArrayLength)
	for i := 0; i < count; i++ {
		roundArray[i] = i + 1
	}
	rounds := make([][]int, maxRound)
	for r := 0; r < maxRound; r++ {
		rounds[r] = calcRoundSlice(r, maxRound, roundArray)
	}
	return rounds
}

func getData(data []int, i int) int {
	if i > 0 {
		return data[i-1]
	}

	return MINVALUE
}

// TopData 可取得前n个最大值的slice
type TopData []int

// New 新建
func New(n int) *TopData {
	d := TopData(make([]int, n))
	return &d
}

// Rand 随机填充
func (d *TopData) Rand(n int) *TopData {
	rand.Seed(time.Now().Unix())
	data := []int(*d)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(n) + 1
	}
	return d
}

// Seq 顺序填充
func (d *TopData) Seq() *TopData {
	data := []int(*d)
	for i := 0; i < len(data); i++ {
		data[i] = i + 1
	}
	return d
}

// Top 取前n个最大值
func (d *TopData) Top(n int) ([]int, error) {
	count := len(*d)
	if count < 2*n {
		return nil, errors.New("Top N is too big")
	}

	// 拷贝一份防止修改原数据
	data := make([]int, count)
	copy(data, []int(*d)) // copy(data, *(*[]int)(d))

	rounds := createRounds(count)
	maxRound := len(rounds)

	// calcMaxDataIndex 计算第r轮的第i与第i+1项的大小，返回数据索引
	calcMaxDataIndex := func(r, i int) int {
		leftDataIndex := rounds[r-1][i]
		rightDataIndex := rounds[r-1][i+1]
		left := getData(data, leftDataIndex)
		right := getData(data, rightDataIndex)

		var maxDataIndex int
		if left >= right {
			maxDataIndex = leftDataIndex
		} else {
			maxDataIndex = rightDataIndex
		}

		return maxDataIndex
	}

	var compareCount int
	// 初始化计算树
	for r := 1; r < maxRound; r++ {
		matchCount := int(math.Ceil(float64(count) / math.Pow(2, float64(r-1))))
		if matchCount%2 != 0 {
			// 补齐为偶数
			matchCount++
		}
		for i := 0; i < matchCount; i += 2 {
			rounds[r][i/2] = calcMaxDataIndex(r, i)
			compareCount++
		}
	}

	top := make([]int, n)
	for i := 0; i < n; i++ {
		winnerDataIndex := rounds[maxRound-1][0] - 1
		top[i] = data[winnerDataIndex]
		// 已经找出的最大值不再参与计算（比较时一定为负）
		data[winnerDataIndex] = MINVALUE
		for r := 1; r < maxRound; r++ {
			dataIndex := winnerDataIndex / int(math.Pow(2, float64(r-1)))
			if dataIndex%2 != 0 {
				// 确保 dataIndex 为偶数
				dataIndex = dataIndex - 1
			}
			rounds[r][dataIndex/2] = calcMaxDataIndex(r, dataIndex)
			compareCount++
		}
	}

	fmt.Printf("\tTopNew 共比较次数: %v\n", compareCount)
	return top, nil
}

// ShowTop 取前n个最大值(debug)
func (d *TopData) ShowTop(n int) {
	topSlice, err := d.Top(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\tTopNew %v: %v\n", n, topSlice)
	}
}
