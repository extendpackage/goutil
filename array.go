package goutil

import (
	"reflect"
	"strconv"
)

// MemSet set `value` to `arr`(slice) with the number of `number`
func AppendServeral(arr interface{}, value interface{}, number int) (newArr interface{}) {
	defer func() {
		if recover() != nil {
			newArr = nil
		}
		return
	}()
	arrValue := reflect.ValueOf(arr)
	valueValue := reflect.ValueOf(value)
	for i := 0; i < number; i++ {
		arrValue = reflect.Append(arrValue, valueValue)
	}
	return arrValue.Interface()
}

//IsInArray judge whether `element` is in `arr`(slice)
func IsInArray(element interface{}, arr interface{}) (exists bool) {
	defer func() {
		if recover() != nil {
			exists = false
		}
		return
	}()
	valueArray := reflect.ValueOf(arr)
	for i := 0; i < valueArray.Len(); i++ {
		if valueArray.Index(i).Interface() == element {
			exists = true
			return
		}
	}
	return
}

// ArrayRemoveRepeat remove repeated element from `arr`
func ArrayRemoveRepeat(arr interface{}) (newArr interface{}) {
	defer func() {
		if recover() != nil {
			newArr = nil
		}
		return
	}()
	arrValue := reflect.ValueOf(arr)
	sliceValue := reflect.MakeSlice(reflect.TypeOf(arr), 0, arrValue.Len())
	for i := 0; i < arrValue.Len(); i++ {
		value := arrValue.Index(i)
		if !IsInArray(value.Interface(), sliceValue.Interface()) {
			sliceValue = reflect.Append(sliceValue, value)
		}
	}
	return sliceValue.Interface()
}

//IntArrayToStringArray transform []int to []string
func IntArrayToStringArray(intArray []int) (strArray []string) {
	for _, intOne := range intArray {
		strArray = append(strArray, strconv.Itoa(intOne))
	}
	return strArray
}

//StringArrayToIntArray transform []string to []int
func StringArrayToIntArray(strArray []string) (intArray []int, err error) {
	for _, oneStr := range strArray {
		oneInt, err := strconv.Atoi(oneStr)
		if err != nil {
			return nil, err
		}
		intArray = append(intArray, oneInt)
	}
	return intArray, nil
}
