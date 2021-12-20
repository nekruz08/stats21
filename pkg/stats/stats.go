package stats

import "github.com/nekruz08/bank/pkg/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	sum:=types.Money(0)
	for _, payment := range payments {
		sum+=payment.Amount
	}
	sum=sum/types.Money((len(payments)))
	return sum
} 

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum:=types.Money(0)
	for _, payment := range payments {
		if payment.Category==category{
			sum+=payment.Amount
		}
	}
	return sum
}