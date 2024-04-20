package utils

import (
	"testing"
	"time"
)

func TestIsDateExpired(t *testing.T) {
	type args struct {
		dateString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Date in the past (expired)
		{
			name: "Date in the past",
			args: args{dateString: "2022-01-01"},
			want: true,
		},
		// Date in the future (not expired)
		{
			name: "Date in the future",
			args: args{dateString: "2025-01-01"},
			want: false,
		},
		// Date equal to current date (not expired)
		{
			name: "Date equal to current date",
			args: args{dateString: time.Now().Format("2006-01-02")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDateExpired(tt.args.dateString); got != tt.want {
				t.Errorf("IsDateExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}
