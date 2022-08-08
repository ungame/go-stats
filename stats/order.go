package stats

type Order uint8

const (
	ORDER_BY_ASC  Order = 0x1
	ORDER_BY_DESC       = 1 << ORDER_BY_ASC
)

var (
	OrderBy = func(o Order, n []float64) func(i, j int) bool {
		return func(i, j int) bool {
			if o == ORDER_BY_ASC {
				return n[i] < n[j]
			}
			return n[i] > n[j]
		}
	}
)
