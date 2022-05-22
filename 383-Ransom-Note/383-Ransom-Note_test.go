package ransomnote

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Test Case 1
		{
			name: "example 1",
			args: args{"a", "b"},
			want: false,
		},
		// Test Case 2
		{
			name: "example 2",
			args: args{"aa", "ab"},
			want: false,
		},
		// Test Case 3
		{
			name: "example 3",
			args: args{"aa", "aab"},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
