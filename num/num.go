package num

import (
	"github.com/shopspring/decimal"
	"math"
)

func Num2() float64 {
	TmpMoney, _ := decimal.NewFromFloat(35.69).Sub(decimal.NewFromFloat(26.73)).Float64() //8.958
	retMoney, _ := decimal.NewFromFloat(TmpMoney).Round(2).Float64()
	return retMoney
}
func Num6() float64 {
	TmpMoney, _ := decimal.NewFromFloat(35.695418).Sub(decimal.NewFromFloat(26.737497)).Float64() //8.958
	retMoney, _ := decimal.NewFromFloat(TmpMoney).Round(6).Float64()
	return retMoney
}
func Big() int { //结果为3
	num := 50
	size := 24
	total := int(math.Ceil(float64(num) / float64(size)))
	return total
}
func Small() int { //结果为2
	num := 50
	size := 24
	total := int(math.Floor(float64(num) / float64(size)))
	return total
}
func SiSheWuRu() int { //2
	var a, b float64
	a = 5
	b = 3
	return round(a / b)
}
func round(x float64) int {
	return int(math.Floor(x + 0.5))
}
