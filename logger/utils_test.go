package logger

import (
	"testing"

	"github.com/wzshiming/jour"
)

func TestParseLevel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantL   jour.Level
		wantErr bool
	}{
		{
			name: "debug",
			args: args{
				s: "debug",
			},
			wantL: jour.LevelDebug,
		},
		{
			name: "info",
			args: args{
				s: "info",
			},
			wantL: jour.LevelInfo,
		},
		{
			name: "-4",
			args: args{
				s: "-4",
			},
			wantL: jour.LevelDebug,
		},
		{
			name: "0",
			args: args{
				s: "0",
			},
			wantL: jour.LevelInfo,
		},
		{
			name: "4",
			args: args{
				s: "4",
			},
			wantL: jour.LevelWarn,
		},
		{
			name: "8",
			args: args{
				s: "8",
			},
			wantL: jour.LevelError,
		},
		{
			name: "info+1",
			args: args{
				s: "info+1",
			},
			wantL: jour.LevelInfo + 1,
		},
		{
			name: "info-1",
			args: args{
				s: "info-1",
			},
			wantL: jour.LevelInfo - 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, err := ParseLevel(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotL != tt.wantL {
				t.Errorf("ParseLevel() gotL = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
