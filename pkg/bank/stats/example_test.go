package stats

import (
	"fmt"
	"github.com/bahrom656/bank3/pkg/bank/types"
)

func ExampleAvg() {
	cards := []types.Payment{
		{
			ID:       1,
			Amount:   20_000_00,
			Category: "car",
		},
		{
			ID:       2,
			Amount:   10_000_00,
			Category: "phone",
		},
	}
	fmt.Println(Avg(cards))
	//Output:
	//1500000
}

func ExampleTotalInCategory() {
	cards := []types.Payment{
		{
			ID:       1,
			Amount:   20_000_00,
			Category: "car",
		},
		{
			ID:       2,
			Amount:   10_000_00,
			Category: "phone",
		},
		{
			ID:       3,
			Amount:   20_000_00,
			Category: "car",
		},
	}
	fmt.Println(TotalInCategory(cards, "car"))
	//Output:
	//4000000
}