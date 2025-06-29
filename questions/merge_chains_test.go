package questions

import (
	"fmt"
	"testing"
	"time"
)

// 7 , 8/1 , 9/4, 2 , 3/5, 6
func TestMergeChains(t *testing.T) {
	c1 := gen([]int{1, 2, 3}, 100*time.Millisecond)
	c2 := gen([]int{4, 5, 6}, 150*time.Millisecond)
	c3 := gen([]int{7, 8, 9}, 50*time.Millisecond)
	merged := merge(c1, c2, c3)
	for i := range merged {
		fmt.Println(i)
	}
}
