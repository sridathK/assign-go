package shape

import (
	"assign1/assign8/model"
	"fmt"
)

func AreaOfCircle(a model.Circle) {
	area := (22 / 7.0) * (a.Radius * a.Radius)
	fmt.Println(area)
}
