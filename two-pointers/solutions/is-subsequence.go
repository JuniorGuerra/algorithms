package solutions

import "fmt"

type mval struct {
	p     int
	count int
}

func IsSubsequence(s string, t string) bool {
	var m = map[int32]mval{}

	for p, val := range t {

		if mv, ok := m[val]; ok {

			m[val] = mval{
				p:     p,
				count: mv.count + 1,
			}
			continue
		}
		m[val] = mval{
			p:     p,
			count: 1,
		}
	}
	var actualPosition int = 0
	for _, val := range s {
		if p, ok := m[val]; ok {
			//fmt.Println(string(val))
			//fmt.Println(p.count)
			m[val] = mval{
				p:     p.p,
				count: p.count - 1,
			}
			fmt.Println(p.p, actualPosition)
			if p.p < actualPosition {
				return false
			}
			actualPosition = p.p
			if m[val].count == 0 {
				delete(m, val)
			}
			//
			continue
		}
		return false
	}
	return true
}
