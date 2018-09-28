package add

import (
    "fmt"
)

func Adder(num ...float64) float64 {
    var total float64 = 0
    for i := 0; i<len(num); i++ {
        total += num[i]
    }
    fmt.Println(total)
    return total
}

