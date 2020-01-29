package main

func main() {

}
func amountOfBonus(sales []int) int {
	sumOfBonus := 0
	const minAmountBonus = 10_000
	const bonusPercent = 5
	const hundredPercent = 100
	for _, sale := range sales {
		if sale > minAmountBonus {
			result := (sale - minAmountBonus) * bonusPercent / hundredPercent
			sumOfBonus = sumOfBonus + result
		}
	}
	return sumOfBonus
}

