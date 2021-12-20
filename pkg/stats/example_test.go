package stats

import (
	"fmt"

	"github.com/nekruz08/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100,
		},
		{
			ID: 2,
			Amount: 100,
		},
		{
			ID: 3,
			Amount: 100,
		},
		{
			ID: 4,
			Amount: 100,
		},
		{
			ID: 5,
			Amount: 100,
		},
	}

	fmt.Println(Avg(payments))
	//Output:
	// 100

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100,
			Category: "travelling",
		},
		{
			ID: 2,
			Amount: 100,
			Category: "donetion",
		},
		{
			ID: 3,
			Amount: 100,
			Category: "donetion",
		},
		{
			ID: 4,
			Amount: 100,
			Category: "donetion",
		},
		{
			ID: 5,
			Amount: 100,
			Category: "donetion",
		},
	}
	fmt.Println(TotalInCategory(payments,"donetion"))
	//Output:
	//400
}