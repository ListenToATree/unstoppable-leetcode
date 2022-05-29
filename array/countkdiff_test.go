package array

import "testing"

func Test_countKDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test", args: args{nums: []int{1, 2, 2, 1}, k: 1}, want: 4},
		{name: "Test2", args: args{nums: []int{33, 90, 62, 43, 21, 96, 20, 18, 84, 74, 61, 100, 5, 11, 4, 67, 96, 18, 6, 68, 82, 32, 76, 33, 93, 33, 71, 32, 30, 63, 37, 46, 95, 51, 63, 77, 63, 84, 52, 78, 66, 76, 66, 9, 73, 92, 79, 65, 29, 42}, k: 1}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countKDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
