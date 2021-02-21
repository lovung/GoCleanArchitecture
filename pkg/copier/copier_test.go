package copier

import (
	"testing"

	"github.com/lovung/GoCleanArchitecture/pkg/testhelper"
)

func TestMustCopy(t *testing.T) {
	type a struct {
		A int
	}
	type b struct {
		A int64
	}
	type c struct {
		A string
	}
	type args struct {
		toValue   interface{}
		fromValue interface{}
	}
	testCases := []struct {
		name string
		args args
		ok   bool
	}{
		{
			args: args{
				toValue:   &a{},
				fromValue: &b{A: 1},
			},
			ok: true,
		},
		{
			args: args{
				toValue:   &a{},
				fromValue: &c{A: "1"},
			},
			ok: true,
		},
		{
			args: args{
				toValue:   a{},
				fromValue: &c{A: "1"},
			},
			ok: false,
		},
		{
			args: args{
				toValue:   &a{},
				fromValue: nil,
			},
			ok: true,
		},
	}
	for _, tt := range testCases {
		if tt.ok {
			t.Run(tt.name, func(t *testing.T) {
				MustCopy(tt.args.toValue, tt.args.fromValue)
			})
		} else {
			testhelper.ShouldPanic(t, func() {
				MustCopy(tt.args.toValue, tt.args.fromValue)
			})
		}
	}
}
