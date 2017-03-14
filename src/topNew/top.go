package top

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
	"unsafe"
)

var _i int

const INTSIZE = uint(unsafe.Sizeof(_i) * 8)
const MINVALUE = -(1<<(INTSIZE-1) - 1)

func init() {
	fmt.Printf("%v %T %v %T\n", INTSIZE, INTSIZE, MINVALUE, MINVALUE)
}

// MINVALUE 数据集中的最小值 (32位)
// const MINVALUE = -(1<<((32-1) - 1))

func calcRoundInfo(count int) (int, int) {
	var maxRound, length int
	for ; count >= 1; count /= 2 {
		maxRound++
		if count > 2 && count%2 != 0 {
			// 补齐为偶数
			count++
		}
		length += count
	}

	return length, maxRound
}

// 根据数据个数count, 生成数组化的计算树
func createRounds(count int) [][]int {
	roundArrayLength, maxRound := calcRoundInfo(count)
	roundArray := make([]int, roundArrayLength)
	for i := 0; i < count; i++ {
		roundArray[i] = i + 1
	}
	rounds := make([][]int, maxRound)
	var start int
	for r := 0; r < maxRound; r++ {
		matchCount := int(math.Ceil(float64(count) / math.Pow(2, float64(r))))
		if matchCount > 2 && matchCount%2 != 0 {
			// 补齐为偶数
			matchCount++
		}
		end := start + matchCount
		rounds[r] = roundArray[start:end:end]
		start += matchCount
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
		leftDataIndex := rounds[r][i]
		rightDataIndex := rounds[r][i+1]
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
		matchCount := len(rounds[r-1])
		for i := 0; i < matchCount; i += 2 {
			rounds[r][i/2] = calcMaxDataIndex(r-1, i)
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
			rounds[r][dataIndex/2] = calcMaxDataIndex(r-1, dataIndex)
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
