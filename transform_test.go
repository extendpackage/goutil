package goutil

import (
	"reflect"
	"testing"
)

func Test_TransfromObject(t *testing.T) {
	type entryType struct {
		Int     int     `json:"i"`
		Str     string  `json:"str"`
		Falot64 float64 `json:"f"`
	}
	entry := entryType{
		Int:     1,
		Str:     "str",
		Falot64: 1.1,
	}

	type args struct {
		from entryType
		to   entryType
	}
	cases := []struct {
		name    string
		args    args
		want    entryType
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				entry,
				entryType{},
			},
			want:    entry,
			wantErr: false,
		},
	}

	for _, oneCase := range cases {
		err := TransfromObject(oneCase.args.from, &oneCase.args.to)
		if (err != nil) != oneCase.wantErr {
			t.Errorf("[%s] TransfromObject err:%v", oneCase.name, err)
		}
		if !reflect.DeepEqual(oneCase.want, oneCase.args.to) {
			t.Errorf("[%s] TransfromObject(%+v,%+v),result want:%+v,to:%+v", oneCase.name, oneCase.args.from, &oneCase.args.to, oneCase.want, oneCase.args.to)
		}
	}

}
