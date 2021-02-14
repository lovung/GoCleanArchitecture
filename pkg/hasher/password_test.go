package hasher

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
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
			got, _ := HashPassword(tt.args.password)
			if ok, _ := CheckPasswordHash(tt.args.password, got); !ok {
				t.Errorf("Got error Hash and Check password hash = %v, password %v", got, tt.args.password)
			}
		})
	}
}

func BenchmarkHashAndCheckPassword(b *testing.B) {
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
		got, _ := HashPassword(tt.args.password)
		if ok, _ := CheckPasswordHash(tt.args.password, got); !ok {
			b.Errorf("Got error Hash and Check password hash = %v, password %v", got, tt.args.password)
		}
	}
}

func BenchmarkHashPassword(b *testing.B) {
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
		HashPassword(tt.args.password)
	}
}
