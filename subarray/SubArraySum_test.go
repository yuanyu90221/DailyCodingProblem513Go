package subarrary

import (
	"reflect"
	"testing"
)

func TestSubArraySum(t *testing.T) {
	type args struct {
		nums []int
		K    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1, 2, 3, 4, 5]",
			args: args{nums: []int{1, 2, 3, 4, 5}, K: 9},
			want: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArraySum(tt.args.nums, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
