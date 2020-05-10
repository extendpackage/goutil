package goutil

import (
	"testing"
)

func Test_FloatRound(t *testing.T) {
	type args struct {
		f float64
		n int
	}
	cases := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ok",
			args: args{
				1.2345,
				2,
			},
			want: 1.23,
		},
	}
	for _, oneCase := range cases {
		res := FloatRound(oneCase.args.f, oneCase.args.n)
		if oneCase.want != res {
			t.Errorf("[%s] FloatRound(%+v,%+v), want :%+v,but get %+v", oneCase.name, oneCase.args.f, oneCase.args.n, oneCase.want, res)
		}

	}
}
