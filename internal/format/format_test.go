package format

import (
	"sort"
	"testing"
)

func Test_quoteRangeTable(t *testing.T) {
	r16 := quoteRangeTable.R16
	for _, r := range r16 {
		if r.Lo > r.Hi {
			t.Errorf("quoteRangeTable has invalid range: %v", r)
		}
	}
	// This test ensures that the quoteRangeTable is sorted.
	isSorted := sort.SliceIsSorted(r16, func(i, j int) bool {
		return r16[i].Lo < r16[j].Lo
	})
	if !isSorted {
		t.Errorf("quoteRangeTable is not sorted")
	}
}

func TestQuoteIfNeed(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty",
			args{
				s: "",
			},
			``,
		},
		{
			"simple",
			args{
				s: "simple",
			},
			`simple`,
		},
		{
			"simple with space",
			args{
				s: "simple with space",
			},
			`"simple with space"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuoteIfNeed(tt.args.s); got != tt.want {
				t.Errorf("QuoteIfNeed() = %q, want %q", got, tt.want)
			}
		})
	}
}
