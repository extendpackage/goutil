package goutil

import (
	"testing"
)

func TestFloatRound(t *testing.T) {
	type Args struct {
		f float64
		n int
	}
	cases := []struct {
		name string
		args Args
		want float64
	}{
		{
			name: "ok",
			args: Args{
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
