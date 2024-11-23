// quickset project quickset.go by suirosu exgaya epowsal wlb iwlb@outlook.com exgaya@gmail.com 20241122;
package quickset

import "fmt"

type QuickSet[T any] struct {
	It []T
}

func New[T any]() *QuickSet[T] {
	qs := &QuickSet[T]{}
	return qs
}

func NewN[T any](cmpf func(a T, b T) int, values ...T) *QuickSet[T] {
	qs := &QuickSet[T]{}
	qs.Insert(cmpf, values...)
	return qs
}

func (qs *QuickSet[T]) Insert(cmpf func(a T, b T) int, values ...T) (ot []T) {
	for _, value := range values {
		if len(qs.It) == 0 {
			qs.It = []T{value}
			continue
		}
		start := 0
		end := len(qs.It) - 1
		cur := end / 2
		for true {
			cmpr := cmpf(value, qs.It[cur])
			if cmpr < 0 {
				end = cur
				cur2 := int(float64(start+end) / 2)
				if cur2 == cur {
					if cur2-1 >= start {
						cur = cur2 - 1
						end = cur
					} else {
						qs.It = append(qs.It[:cur*1+1], qs.It[cur*1:]...)
						qs.It[cur*1] = value
						//return qs.It
						goto insend
					}
				} else {
					cur = cur2
				}
			} else if cmpr > 0 {
				start = cur
				cur2 := int(float64(start+end) / 2)
				if cur2 == cur {
					if cur2+1 <= end {
						cur = cur2 + 1
						start = cur
					} else {
						//fmt.Println(qs.It, cur, value)
						if (cur+1)*1+1 <= len(qs.It) {
							qs.It = append(qs.It[:(cur+1)*1+1], qs.It[(cur+1)*1:]...)
							qs.It[(cur+1)*1] = value
						} else {
							qs.It = append(qs.It, value)
						}
						//return qs.It
						goto insend
					}
				} else {
					cur = cur2
				}
			} else {
				//return qs.It
				goto insend
			}
		}
	insend:
	}
	return qs.It
}

func (qs *QuickSet[T]) Remove(cmpf func(a T, b T) int, values ...T) (ot []T) {
	for _, value := range values {
		if len(qs.It) == 0 {
			break
		}
		start := 0
		end := len(qs.It) - 1
		cur := end / 2
		for cur < len(qs.It) {
			cmpr := cmpf(value, qs.It[cur])
			if cmpr < 0 {
				end = cur
				cur2 := int(float64(start+end) / 2)
				if cur2 == cur {
					if cur2-1 >= start {
						cur = cur2 - 1
						end = cur
					} else {
						fmt.Println("<", qs.It)
						goto rmend
					}
				} else {
					cur = cur2
				}
			} else if cmpr > 0 {
				start = cur
				cur2 := int(float64(start+end) / 2)
				if cur2 == cur {
					if cur2+1 <= end {
						cur = cur2 + 1
						start = cur
					} else {
						fmt.Println(">", qs.It)
						goto rmend
					}
				} else {
					cur = cur2
				}
			} else {
				qs.It = append(qs.It[:cur], qs.It[cur+1:]...)
				goto rmend
			}
		}
	rmend:
	}
	return qs.It
}

func (qs *QuickSet[T]) Index(cmpf func(a T, b T) int, value T) (index int) {
	index = -1
	start := 0
	end := len(qs.It) - 1
	cur := end / 2
	for cur < len(qs.It) {
		cmpr := cmpf(value, qs.It[cur])
		if cmpr < 0 {
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
		} else if cmpr > 0 {
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
