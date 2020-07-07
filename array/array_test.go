package array

import (
	"reflect"
	"testing"
)

func TestAppendServeral(t *testing.T) {
	type Args struct {
		arr    interface{}
		value  interface{}
		number int
	}
	cases := []struct {
		name string
		args Args
	}{
		{
			name: "int",
			args: Args{
				[]int{},
				2,
				10,
			},
		},
		{
			name: "string",
			args: Args{
				[]string{},
				"test",
				10,
			},
		},
	}
	for _, oneCase := range cases {
		res := AppendServeral(oneCase.args.arr, oneCase.args.value, oneCase.args.number)
		arrValue := reflect.ValueOf(res)
		for i := int(0); i < oneCase.args.number; i++ {
			if arrValue.Index(i).Interface() != oneCase.args.value {
				t.Errorf("[%s] AppendServeral(%+v,%+v,%+v), index i:%+v,current value:%+v,want value:%+v", oneCase.name, oneCase.args.arr, oneCase.args.value, oneCase.args.number, i, res, oneCase.args.value)
			}
		}

	}
}

func TestIsInArray(t *testing.T) {
	type Args struct {
		element interface{}
		arr     interface{}
	}
	cases := []struct {
		name string
		args Args
		want bool
	}{
		{
			name: "ok_true",
			args: Args{
				1,
				[]int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "ok_false",
			args: Args{
				100,
				[]int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "ok_string",
			args: Args{
				"test1",
				[]string{"test1", "test2", "test3"},
			},
			want: true,
		},
	}
	for _, oneCase := range cases {
		res := IsInArray(oneCase.args.element, oneCase.args.arr)
		if res != oneCase.want {
			t.Errorf("[%+v] IsInArray(%+v, %+v) want %+v ,but get %+v", oneCase.name, oneCase.args.element, oneCase.args.arr, oneCase.want, res)
		}
	}
}

func TestArrayRemoveRepeat(t *testing.T) {
	type Args struct {
		arr interface{}
	}
	cases := []struct {
		name string
		args Args
		want interface{}
	}{
		{
			name: "ok_int",
			args: Args{
				[]int{1, 1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "ok_string",
			args: Args{
				[]string{"test1", "test1", "test2", "test3"},
			},
			want: []string{"test1", "test2", "test3"},
		},
	}
	for _, oneCase := range cases {
		res := ArrayRemoveRepeat(oneCase.args.arr)
		if !reflect.DeepEqual(res, oneCase.want) {
			t.Errorf("[%+v] IsInArray(%+v) want %+v ,but get %+v", oneCase.name, oneCase.args.arr, oneCase.want, res)
		}
	}
}

func TestIntArrayToStringArray(t *testing.T) {
	type Args struct {
		arr []int
	}
	cases := []struct {
		name string
		args Args
		want []string
	}{
		{
			name: "ok",
			args: Args{
				[]int{1, 2, 3},
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, oneCase := range cases {
		res := IntArrayToStringArray(oneCase.args.arr)
		if !reflect.DeepEqual(res, oneCase.want) {
			t.Errorf("[%+v] IntArrayToStringArray(%+v) want %+v ,but get %+v", oneCase.name, oneCase.args.arr, oneCase.want, res)
		}
	}
}

func TestStringArrayToIntArray(t *testing.T) {
	type Args struct {
		arr []string
	}
	cases := []struct {
		name    string
		args    Args
		want    []int
		wantErr bool
	}{
		{
			name: "ok",
			args: Args{
				[]string{"1", "2", "3"},
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name: "ok_error",
			args: Args{
				[]string{"1", "2", "3", "3t"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, oneCase := range cases {
		res, err := StringArrayToIntArray(oneCase.args.arr)
		if (err != nil) != oneCase.wantErr {
			t.Errorf("[%+v] StringArrayToIntArray(%+v) wantErr %+v ,but get err: %+v", oneCase.name, oneCase.args.arr, oneCase.wantErr, err)
		}
		if !reflect.DeepEqual(res, oneCase.want) {
			t.Errorf("[%+v] IntArrayToStringArray(%+v) want %+v ,but get %+v", oneCase.name, oneCase.args.arr, oneCase.want, res)
		}
	}
}
