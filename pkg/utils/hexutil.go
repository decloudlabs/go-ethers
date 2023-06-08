package utils

import (
	"strconv"
)

// @title         DecimalToHex
// @description   decimal to hex
// @param         decimal        int64         "decimal"
// @auth          happyboy
func DecimalToHex(decimal int64) string {
	return strconv.FormatInt(decimal, 16)
}

// @title         DecimalToHex
// @description   decimal to hex
// @param         decimal        int64         "decimal"
// @auth          happyboy
func HexToDecimal(decimal string) int64 {
	decimal = decimal[2:len(decimal)]
	r, _ := strconv.ParseInt(decimal, 16, 32)
	return r
}
