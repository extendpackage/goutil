package encryption

import (
	"reflect"
	"testing"
)

func TestHmacSha256_Base64(t *testing.T) {
	type Args struct {
		secret  string
		message string
	}
	cases := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "ok",
			args: Args{
				secret:  "test_secret",
				message: "test_message",
			},
			want: "RjKGASqtnQnomajFGUjS5yyIeH0R3x+icaA9Qdiuk7o=",
		},
	}
	for _, oneCase := range cases {
		res := HmacSha256_Base64(oneCase.args.secret, oneCase.args.message)
		if !reflect.DeepEqual(res, oneCase.want) {
			t.Errorf("[%v] HmacSha256_Base64(%v,%v) get %v,but want %v", oneCase.name, oneCase.args.secret, oneCase.args.message, res, oneCase.want)
		}
	}
}
