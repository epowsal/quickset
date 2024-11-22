package quickset

import (
	"cmp"
	"fmt"
	"testing"
)

func check(b bool) {
	if b == false {
		panic("error")
	}
}

func Test1(t *testing.T) {
	qs := New[int]()
	kes2 := qs.Insert(cmp.Compare[int], 3, 5, 52, 1, 9, 44, 22)
	fmt.Println(kes2)
	kes2 = qs.Remove(cmp.Compare[int], 1, 44)
	fmt.Println(kes2)
	ind1 := qs.Index(cmp.Compare[int], 3)
	fmt.Println(ind1)
	check(ind1 == 0)
	ind2 := qs.Index(cmp.Compare[int], 52)
	fmt.Println(ind2)
	check(ind2 == 4)
	ind3 := qs.Index(cmp.Compare[int], 32)
	fmt.Println(ind3)
	check(ind3 == -1)
	ind4 := qs.Index(cmp.Compare[int], 9)
	fmt.Println(ind4)
	check(ind4 == 2)

}
