package top

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// copy
func dupSlice(s []int) (d []int) {
	d = make([]int, len(s), cap(s))
	for i, v := range s {
		d[i] = v
	}
	return
}

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

func debugRoundArray(roundArray []int, nRound int) {
	var index, round int
	for i := nRound; i >= 0; i-- {
		len := int(math.Pow(2, float64(i)))
		round++
		fmt.Printf("Round%v: %v\n", round, roundArray[index:index+len])
		index += len
	}
}

func mapToRoundArrayIndex(round, maxRound, index int) int {
	var count int
	for i := 0; i < round; i++ {
		count += int(math.Pow(2, float64(maxRound)))
		maxRound--
	}

	return count + index
}

func calcRoundArrayLength(nRound int) int {
	// var count int
	// for i := nRound; i >= 0; i-- {
	// 	count += int(math.Pow(2, float64(i)))
	// }
	// return count

	return mapToRoundArrayIndex(nRound, nRound, 1)
}

// func getData(data, roundData []int, round, maxRound, index int) int {
// 	return data[roundData[mapToRoundArrayIndex(round, maxRound, index)]]
// }

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
		data[i] = rand.Intn(n)
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

	// 拷贝一份防止修改
	data := make([]int, count)
	copy(data, []int(*d))
	//copy(data, *(*[]int)(d))

	extN := near2Pow(count, true)
	maxRound := int(math.Log2(float64(extN)))

	roundArrayLength := calcRoundArrayLength(maxRound)
	// fmt.Println(count, extN, maxRound, roundArrayLength)

	roundArray := make([]int, roundArrayLength)
	for i := 0; i < count; i++ {
		roundArray[i] = i + 1
	}
	// debugRoundArray(roundArray, maxRound)

	var compareCount int
	for r := 0; r < maxRound; r++ {
		matchCount := int(math.Ceil(float64(count) / math.Pow(2, float64(r))))
		if matchCount%2 != 0 {
			matchCount++
		}
		for i := 0; i < matchCount; i += 2 {
			leftIndex := mapToRoundArrayIndex(r, maxRound, i)
			rightIndex := mapToRoundArrayIndex(r, maxRound, i+1)
			leftDataIndex := roundArray[leftIndex]
			rightDataIndex := roundArray[rightIndex]
			var left, right int
			if leftDataIndex > 0 {
				left = data[leftDataIndex-1]
			}
			if rightDataIndex > 0 {
				right = data[rightDataIndex-1]
			}

			// left := getData(data, roundArray, r, maxRound, i)
			// righ := getData(data, roundArray, r, maxRound, i+1)

			var maxDataIndex int
			if left >= right {
				maxDataIndex = leftDataIndex
			} else {
				maxDataIndex = rightDataIndex
			}
			nextRoundIndex := mapToRoundArrayIndex(r+1, maxRound, i/2)
			roundArray[nextRoundIndex] = maxDataIndex
			compareCount++
		}
	}
	// debugRoundArray(roundArray, maxRound)

	r := make([]int, n)
	for i := 0; i < n; i++ {
		winnerDataIndex := roundArray[roundArrayLength-1] - 1
		r[i] = data[winnerDataIndex]
		// 已经找出的最大值不再参与计算（比较时一定为负）
		data[winnerDataIndex] = 0
		for r := 0; r < maxRound; r++ {
			nextRoundDataIndex := winnerDataIndex / int(math.Pow(2, float64(r)))
			// 确保 nextRoundIndex 为偶数
			var leftI, rightI int
			if nextRoundDataIndex%2 == 0 {
				leftI = nextRoundDataIndex
			} else {
				leftI = nextRoundDataIndex - 1
			}
			rightI = leftI + 1

			leftIndex := mapToRoundArrayIndex(r, maxRound, leftI)
			rightIndex := mapToRoundArrayIndex(r, maxRound, rightI)
			leftDataIndex := roundArray[leftIndex]
			rightDataIndex := roundArray[rightIndex]

			var maxDataIndex int
			var left, right int
			if leftDataIndex > 0 {
				left = data[leftDataIndex-1]
			} else {
				maxDataIndex = rightDataIndex
			}
			if rightDataIndex > 0 {
				right = data[rightDataIndex-1]
			} else {
				maxDataIndex = leftDataIndex
			}

			if maxDataIndex == 0 {
				if left >= right {
					maxDataIndex = leftDataIndex
				} else {
					maxDataIndex = rightDataIndex
				}
			}
			nextRoundIndex := mapToRoundArrayIndex(r+1, maxRound, leftI/2)
			roundArray[nextRoundIndex] = maxDataIndex
			compareCount++
		}
		// debugRoundArray(roundArray, maxRound)
	}

	fmt.Printf("\t共比较次数: %v\n", compareCount)
	return r, nil
}

// ShowTop 取前n个最大值(debug)
func (d *TopData) ShowTop(n int) {
	topSlice, err := d.Top(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\tTop %v: %v\n", n, topSlice)
	}
}
