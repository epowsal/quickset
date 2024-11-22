// quickset project quickset.go
package quickset

import "fmt"

type QuickSet[T any] struct {
	it []T
}

func New[T any]() *QuickSet[T] {
	qs := &QuickSet[T]{}
	return qs
}

func (qs *QuickSet[T]) Insert(cmpf func(a T, b T) int, values ...T) (ot []T) {
	for _, value := range values {
		if len(qs.it) == 0 {
			qs.it = []T{value}
			continue
		}
		start := 0
		end := len(qs.it) - 1
		cur := end / 2
		for true {
			if cmpf(value, qs.it[cur]) < 0 {
				end = cur
				cur2 := int(float64(start+end) / 2)
				if cur2 == cur {
					if cur2-1 >= start {
						cur = cur2 - 1
						end = cur
					} else {
						qs.it = append(qs.it[:cur*1+1], qs.it[cur*1:]...)
						qs.it[cur*1] = value
						//return qs.it
						goto insend
					}
				} else {
					cur = cur2
				}
			} else if cmpf(value, qs.it[cur]) > 0 {
				start = cur
				cur2 := int(float64(start+end) / 2)
				if cur2 == cur {
					if cur2+1 <= end {
						cur = cur2 + 1
						start = cur
					} else {
						//fmt.Println(qs.it, cur, value)
						if (cur+1)*1+1 <= len(qs.it) {
							qs.it = append(qs.it[:(cur+1)*1+1], qs.it[(cur+1)*1:]...)
							qs.it[(cur+1)*1] = value
						} else {
							qs.it = append(qs.it, value)
						}
						//return qs.it
						goto insend
					}
				} else {
					cur = cur2
				}
			} else {
				//return qs.it
				goto insend
			}
		}
	insend:
	}
	return qs.it
}

func (qs *QuickSet[T]) Remove(cmpf func(a T, b T) int, values ...T) (ot []T) {
	for _, value := range values {
		if len(qs.it) == 0 {
			break
		}
		start := 0
		end := len(qs.it) - 1
		cur := end / 2
		for true {
			if cmpf(value, qs.it[cur]) < 0 {
				end = cur
				cur2 := int(float64(start+end) / 2)
				if cur2 == cur {
					if cur2-1 >= start {
						cur = cur2 - 1
						end = cur
					} else {
						fmt.Println("<", qs.it)
						goto rmend
					}
				} else {
					cur = cur2
				}
			} else if cmpf(value, qs.it[cur]) > 0 {
				start = cur
				cur2 := int(float64(start+end) / 2)
				if cur2 == cur {
					if cur2+1 <= end {
						cur = cur2 + 1
						start = cur
					} else {
						fmt.Println(">", qs.it)
						goto rmend
					}
				} else {
					cur = cur2
				}
			} else {
				qs.it = append(qs.it[:cur], qs.it[cur+1:]...)
				goto rmend
			}
		}
	rmend:
	}
	return qs.it
}

func (qs *QuickSet[T]) Index(cmpf func(a T, b T) int, value T) (index int) {
	index = -1
	start := 0
	end := len(qs.it) - 1
	cur := end / 2
	for true {
		if cmpf(value, qs.it[cur]) < 0 {
			end = cur
			cur2 := int(float64(start+end) / 2)
			if cur2 == cur {
				if cur2-1 >= start {
					cur = cur2 - 1
					end = cur
				} else {
					return -1
				}
			} else {
				cur = cur2
			}
		} else if cmpf(value, qs.it[cur]) > 0 {
			start = cur
			cur2 := int(float64(start+end) / 2)
			if cur2 == cur {
				if cur2+1 <= end {
					cur = cur2 + 1
					start = cur
				} else {
					return -1
				}
			} else {
				cur = cur2
			}
		} else {
			return cur
		}
	}
	return index
}
