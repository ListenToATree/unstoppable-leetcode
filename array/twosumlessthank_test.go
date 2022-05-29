package array

import "testing"

func Test_twoSumLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test", args: args{nums: []int{34, 23, 1, 24, 75, 33, 54, 8}, k: 60}, want: 58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("twoSumLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
