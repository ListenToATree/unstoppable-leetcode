package array

import "testing"

func Test_countHillValley(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{[]int{2, 4, 1, 1, 6, 5}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHillValley(tt.args.nums); got != tt.want {
				t.Errorf("countHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}
