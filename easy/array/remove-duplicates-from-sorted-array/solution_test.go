package remove_duplicates_from_sorted_array

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "demo-1",
			args: args{nums: []int{1, 1, 2}},
			want: 2,
		},
		{
			name: "demo-2",
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: 5,
		},
		{
			name: "demo-3",
			args: args{nums: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
