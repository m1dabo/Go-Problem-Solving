package RomanToInteger

import "testing"

func Test_romanToInt(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Test Case 1
		{
			name: "example 1",
			args: args{"III"},
			want: 3,
		},
		// Test Case 2
		{
			name: "example 2",
			args: args{"LVIII"},
			want: 58,
		},
		// Test Case 3
		{
			name: "example 3",
			args: args{"MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
