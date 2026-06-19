package interest

const (
    NegativeRate float32 = 3.213
    LowRate      float32 = 0.5
    MediumRate   float32 = 1.621
    HighRate     float32 = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return NegativeRate
	case balance < 1000:
		return LowRate
	case balance < 5000:
		return MediumRate
	default:
		return HighRate
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance)
	return balance * float64(rate) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0

	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
	}

	return years
}
