package stats

import "github.com/bahrom656/bank3/pkg/bank/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	if len(payments) == 0 {
		return payments[0].Amount
	}
	var sum types.Money
	var count int
	for i := 0; i < len(payments); i++ {
		sum = sum + payments[i].Amount
		count++
	}
	return sum / types.Money(count)
}