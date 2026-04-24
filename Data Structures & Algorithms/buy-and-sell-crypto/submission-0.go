import "slices"

func maxProfit(prices []int) int {
    // wrong intuition: searching for max and min wont solve because order is
    //   important as this is by date. min's index could be later after max's
    // naive approach: O(n^2), set one at the beginning, iterate over the rest
    //   to found the value that can gain profit (value should be larger)
    
    // this is naive approach
    max := 0
    for i, buy := range prices {
        if i == len(prices)-1 { break }
        sell := slices.Max(prices[i+1:])
        profit := sell - buy
        if profit > max { max = profit }
    }
    return max
}
