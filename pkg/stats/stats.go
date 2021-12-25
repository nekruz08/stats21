package stats

import "github.com/nekruz08/bank/v2/pkg/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	i:=types.Money(0)
	sum:=types.Money(0)
	for _, payment := range payments {
		if payment.Status!=types.StatusFail{
			sum+=payment.Amount
			i++
		}

	}
	sum=sum/i
	return sum
} 

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum:=types.Money(0)
	for _, payment := range payments {
		if payment.Category==category{
			if payment.Status!=types.StatusFail{
				sum+=payment.Amount
			}
		}
	}
	return sum
}

// FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category==category{
			filtered = append(filtered, payment)
		}
	}
	
	return filtered
}