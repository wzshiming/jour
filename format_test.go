package jour

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/wzshiming/jour/internal/format"
)

func Test_formatLog(t *testing.T) {
	type args struct {
		msg       string
		attrsStr  string
		level     Level
		termWidth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "msg",
			args: args{
				msg: "msg",
			},
			want: "msg\n",
		},
		{
			name: "msg with attrs",
			args: args{
				msg:      "msg",
				attrsStr: `a=b`,
			},
			want: "msg a=b\n",
		},
		{
			name: "msg with attrs and level",
			args: args{
				msg:      "msg",
				attrsStr: `a=b`,
				level:    LevelDebug,
			},
			want: "\x1b[0;36mDEBUG\x1b[0m msg a=b\n",
		},
		{
			name: "msg with attrs and termWidth",
			args: args{
				msg:       "msg",
				attrsStr:  `a=b`,
				termWidth: 20,
			},
			want: "msg              a=b\n",
		},
		{
			name: "msg with attrs and level and termWidth",
			args: args{
				msg:       "msg",
				attrsStr:  `a=b`,
				level:     LevelDebug,
				termWidth: 20,
			},
			want: "\x1b[0;36mDEBUG\x1b[0m msg        a=b\n",
		},
		{
			name: "msg with attrs and 5 termWidth",
			args: args{
				msg:       "msg",
				attrsStr:  `a=b`,
				termWidth: 5,
			},
			want: "msg a=b\n",
		},
		{
			name: "msg with attrs and 6 termWidth",
			args: args{
				msg:       "msg",
				attrsStr:  `a=b`,
				termWidth: 6,
			},
			want: "msg a=b\n",
		},
		{
			name: "msg with attrs and 7 termWidth",
			args: args{
				msg:       "msg",
				attrsStr:  `a=b`,
				termWidth: 7,
			},
			want: "msg a=b\n",
		},
		{
			name: "msg with attrs and 8 termWidth",
			args: args{
				msg:       "msg",
				attrsStr:  `a=b`,
				termWidth: 8,
			},
			want: "msg  a=b\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatLog(tt.args.msg, tt.args.attrsStr, tt.args.level, tt.args.termWidth); got != tt.want {
				t.Errorf("formatLog() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_formatValue(t *testing.T) {
	type args struct {
		val slog.Value
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "format for error",
			args: args{
				val: slog.AnyValue(errors.New("unknown command \"subcommand\" for \"jour\"")),
			},
			want: format.QuoteIfNeed(errors.New("unknown command \"subcommand\" for \"jour\"").Error()),
		},
		{
			name: "format for string",
			args: args{
				val: slog.AnyValue("unknown command \"subcommand\" for \"jour\""),
			},
			want: format.QuoteIfNeed("unknown command \"subcommand\" for \"jour\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatValue(tt.args.val); got != tt.want {
				t.Errorf("formatValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
