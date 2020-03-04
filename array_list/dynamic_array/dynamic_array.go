package dynamic_array

import "fmt"

type array struct {
	arr []interface{}
	len uint
	cap uint
}

func NewArray(cap uint) *array {
	return &array{arr: make([]interface{}, 10), len: 0, cap: cap}
}

func (arr *array) Push(data interface{}) {
	if arr.len >= arr.cap {
		t := make([]interface{}, arr.cap*2)
		for k, v := range arr.arr {
			t[k] = v
		}
		arr.arr = t
		arr.cap *= 2
	}
	arr.arr[arr.len] = data
	arr.len++
}

func (arr *array) Get(idx uint) interface{} {
	if idx >= arr.len {
		return nil
	}
	return arr.arr[idx]
}

func (arr *array) Delete(idx uint) interface{} {
	if idx >= arr.len {
		return nil
	}

	if arr.len <= arr.cap/3 {
		t := make([]interface{}, arr.cap/2)
		for i := uint(0); i < arr.len; i++ {
			t[i] = arr.arr[i]
		}
		arr.arr = t
		arr.cap /= 2
	}

	del := arr.arr[idx]
	arr.arr[idx] = arr.arr[arr.len-1]
	arr.len--
	return del
}

func (arr *array) Len() uint {
	return arr.len
}

func (arr *array) Cap() uint {
	return arr.cap
}

func (arr *array) Print() {
	var format string
	for i := uint(0); i < arr.len; i++ {
		format += fmt.Sprintf("|%+v", arr.arr[i])
	}
	fmt.Printf("len: %d, cap: %d\n", arr.Len(), arr.Cap())
	fmt.Println(format)
}
