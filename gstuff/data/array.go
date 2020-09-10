package data

import (
	"reflect"
)

// GArray ..
type GArray struct{}

// Array ..
var Array GArray

func init() {
	Array = GArray{}
}

// MaxInt ..
func (GArray) MaxInt(arr []int) (max int) {
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return
}

// MaxFloat ..
func (GArray) MaxFloat(arr []float64) (max float64) {
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return
}

// MinInt ..
func (GArray) MinInt(arr []int) (min int) {
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}

// MinFloat ..
func (GArray) MinFloat(arr []float64) (min float64) {
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}

// Index ..
func (GArray) Index(arr interface{}, item interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(item, s.Index(i).Interface()) == true {
				index = i
				break
			}
		}
	}
	return
}
