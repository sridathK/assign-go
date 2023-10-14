package temparature

import (
	"errors"
	"fmt"
)

var ErrTempNotInRange = errors.New("temp not in valid range")

func FtoCConvertor(temp float64) float64 {

	if -459 > temp || temp > 213 {
		err := ErrTempNotInRange
		fmt.Println(err)
		return 0

	}
	return (temp - 32) * 5 / 9

}
