package money_change

type changeSet struct {
	one    uint
	two    uint
	five   uint
	ten    uint
	twenty uint
	fifty  uint
}

// The problem: You get a fixed money value like 321
// and you need to give the best change possible.
// So, big coins first and so on..
// 321 == 6x 50-coin, 1x 20-coin, 1x 1-coin
// Time complexity: O(n)
// Space complexity: O(n)
func giveChange(input uint) (chs changeSet) {
	for input != 0 {
		switch {
		case input >= 50:
			chs.fifty++
			input -= 50
		case input >= 20:
			chs.twenty++
			input -= 20
		case input >= 10:
			chs.ten++
			input -= 10
		case input >= 5:
			chs.five++
			input -= 5
		case input >= 2:
			chs.two++
			input -= 2
		case input >= 1:
			chs.one++
			input--
		default:
			continue
		}
	}
	return chs
}
