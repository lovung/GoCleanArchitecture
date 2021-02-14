package hasher

import "testing"

func Test_bcryptHashAndCheckPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{password: "testpassword"}},
		{args: args{password: "!@#$%^&*()_+"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := bcryptHashPassword(tt.args.password)
			if ok, _ := bcryptCheckPasswordHash(tt.args.password, got); !ok {
				t.Errorf("Got error Hash and Check password hash = %v, password %v", got, tt.args.password)
			}
		})
	}
}

func Benchmark_bcryptHashAndCheckPassword(b *testing.B) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{password: "testpassword"}},
		{args: args{password: "!@#$%^&*()_+"}},
	}
	for _, tt := range tests {
		got, _ := bcryptHashPassword(tt.args.password)
		if ok, _ := bcryptCheckPasswordHash(tt.args.password, got); !ok {
			b.Errorf("Got error Hash and Check password hash = %v, password %v", got, tt.args.password)
		}
	}
}

func Benchmark_bcryptHashPassword(b *testing.B) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{password: "testpassword"}},
		{args: args{password: "!@#$%^&*()_+"}},
	}
	for _, tt := range tests {
		bcryptHashPassword(tt.args.password)
	}
}
