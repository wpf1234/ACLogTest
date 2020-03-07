package util

import "testing"

func TestInetNtoA(t *testing.T) {
	type args struct {
		ip int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{2886734290},
			want: "172.16.17.210",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InetNtoA(tt.args.ip); got != tt.want {
				t.Errorf("InetNtoA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInetAtoN(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.

		{
			name: "",
			args: args{"172.16.17.210"},
			want: 2886734290,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InetAtoN(tt.args.ip); got != tt.want {
				t.Errorf("InetAtoN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
