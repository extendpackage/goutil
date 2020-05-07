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
	tests := []struct {
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

	for _, tt := range tests {
		err := TransfromObject(tt.args.from, &tt.args.to)
		if (err != nil) != tt.wantErr {
			t.Errorf("[%s] TransfromObject err:%v", tt.name, err)
		}
		if !reflect.DeepEqual(tt.want, tt.args.to) {
			t.Errorf("[%s] TransfromObject(%+v,%+v),result want:%+v,to:%+v", tt.name, tt.args.from, &tt.args.to, tt.want, tt.args.to)
		}
	}

}
