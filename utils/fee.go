package utils

func CalcFee(price float64, fee float64) (float64, float64) {
	feer := (fee / 10) * price
	return price - feer, feer
}
