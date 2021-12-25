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

// CategoriesTotal возвращает сумму платежей по каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories:= map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount
	}
	return categories
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

// CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories:= map[types.Category]types.Money{}
	i:=map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount
		i[payment.Category]++
	}

	for key,value:= range categories {
		categories[key]=value/i[key]
	}

	return categories
}

// PeriodsDynamic сравнивает расходы по категориям за два периода
func PeriodsDynamic(
 first map[types.Category]types.Money,second map[types.Category]types.Money,
 )map[types.Category]types.Money{

	list:=map[types.Category]types.Money{} 
	for key:= range first {
		list[key]=0
	}
	
	for key:= range second {
		list[key]=0
	}

	perioDynamic:=map[types.Category]types.Money{} 

	for key:= range list {
		perioDynamic[key]=second[key]-first[key]
	}	
	return perioDynamic
}