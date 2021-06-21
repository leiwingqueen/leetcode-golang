package dp

import "testing"

func Test_palindromePartition(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{s: "aabbc", k: 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePartition(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("palindromePartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
