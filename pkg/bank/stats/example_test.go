package stats

import (
	"fmt"
	"github.com/rustamfozilov/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 100,
		},
		{
			Amount: 200,
		},
	}

	fmt.Println(Avg(payments))
	// Output:
	// 150
}