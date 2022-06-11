package dp

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		nums        []int
		multipliers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test", args: args{nums: []int{1, 2, 3}, multipliers: []int{3, 2, 1}}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.nums, tt.args.multipliers); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
