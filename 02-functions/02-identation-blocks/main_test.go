package main

import "testing"

func TestIsValidIp(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "should return true",
			ip:   "127.0.0.1",
			want: true,
		},
		{
			name: "should return true",
			ip:   "192.168.0.1",
			want: true,
		},
		{
			name: "should return true",
			ip:   "192.168.152.3",
			want: true,
		},
		{
			name: "should return false when it is empty",
			ip:   "",
			want: false,
		},
		{
			name: "should return false when it has invalid chars",
			ip:   "abc.def.ghi.jkl",
			want: false,
		},
		{
			name: "should return false when it is not an octet",
			ip:   "12.34.56",
			want: false,
		},
		{
			name: "should return false when it has an invalid start",
			ip:   "027.0.0.1",
			want: false,
		},
		{
			name: "should return false when has an invalid range",
			ip:   "900.900.900.900",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidIp(tt.ip); got != tt.want {
				t.Errorf("IsValidIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
