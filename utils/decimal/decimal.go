package decimal

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)
/*
	decimal 提供大数字计算能力
*/

// string(float64) 转 M
func ToM(fs string) float64 {
	f, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		return 0
	}
	return Float64DivFmt(f, 1000000, 2)
}

func Float64Div(f float64, divisor int) float64 {
	return f / float64(divisor)
}

// 被除数、除数、保留小数点后位数
func Float64DivFmt(f float64, divisor int, digits int) float64 {
	res := Float64Div(f, divisor)
	fs := fmt.Sprintf("%."+strconv.Itoa(digits)+"f", res)
	resFloat64, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		return 0
	}
	return resFloat64
}

// 计算代币最小单位总额
// amount 数额 precision 精度
func MulStringDecimals(amount string, precision int) string {
	aDecimal, err := decimal.NewFromString(amount)
	if err != nil {
		return ""
	}
	// digit, _ := strconv.Atoi(precision)
	power := decimal.NewFromInt(int64(precision)) // 10 的 N 次方
	ten := decimal.NewFromInt(10)
	res := aDecimal.Mul(ten.Pow(power))
	return res.Round(18).String()
}

// 计算代币除以精度后的总额
func DivStringDecimals(tokenTotalSupply string, precision int) string {
	tokenTotalSupplyDecimal, err := decimal.NewFromString(tokenTotalSupply)
	if err != nil {
		log.Print(err)
	}
	divisor := decimal.NewFromInt(int64(math.Pow(10, float64(precision))))
	tokenTotalSupplyDecimalNumber := tokenTotalSupplyDecimal.Div(divisor)
	return tokenTotalSupplyDecimalNumber.Round(18).String()
}

func GreaterOrEqual(a, b string) bool {
	aDecimal, err := decimal.NewFromString(a)
	if err != nil {
		log.Print(err)
	}
	bDecimal, err := decimal.NewFromString(b)
	if err != nil {
		log.Print(err)
	}
	return aDecimal.GreaterThanOrEqual(bDecimal)
}

// 返回是否超过指定值
func Greater(a, b string) bool {
	aDecimal, err := decimal.NewFromString(a)
	if err != nil {
		log.Print(err)
	}
	bDecimal, err := decimal.NewFromString(b)
	if err != nil {
		log.Print(err)
	}
	return aDecimal.GreaterThan(bDecimal)
}

func LessThan(a, b string) bool {
	aDecimal, err := decimal.NewFromString(a)
	if err != nil {
		log.Print(err)
	}
	bDecimal, err := decimal.NewFromString(b)
	if err != nil {
		log.Print(err)
	}
	return aDecimal.LessThan(bDecimal)
}

func LessThanOrEqual(a, b string) bool {
	aDecimal, err := decimal.NewFromString(a)
	if err != nil {
		log.Print(err)
	}
	bDecimal, err := decimal.NewFromString(b)
	if err != nil {
		log.Print(err)
	}
	return aDecimal.LessThanOrEqual(bDecimal)
}

func Mul(a, b string) string {
	aDecimal, err := decimal.NewFromString(a)
	if err != nil {
		log.Print(err)
	}
	bDecimal, err := decimal.NewFromString(b)
	if err != nil {
		log.Print(err)
	}
	return aDecimal.Mul(bDecimal).Round(18).String()
}

func Sub(a, b string) string {
	aDecimal, err := decimal.NewFromString(a)
	if err != nil {
		log.Print(err)
	}
	bDecimal, err := decimal.NewFromString(b)
	if err != nil {
		log.Print(err)
	}
	return aDecimal.Sub(bDecimal).Round(18).String()
}

func Add(a, b string) string {
	aDecimal, err := decimal.NewFromString(a)
	if err != nil {
		log.Print(err)
	}
	bDecimal, err := decimal.NewFromString(b)
	if err != nil {
		log.Print(err)
	}
	return aDecimal.Add(bDecimal).Round(18).String()
}
