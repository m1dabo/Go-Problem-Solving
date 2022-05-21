// I used the below line to create this file I've used gotests lib
// gotests -all -w -parallel .\234-Palindrome-Linked-List\

package palindromelinkedlist

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Test Case 1
		{
			name: "example 1",
			args: args{[1, 2, 2, 1]},
			want: true,
		},
		// Test Case 2
		{
			name: "example 2",
			args: args{[1, 2]},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
