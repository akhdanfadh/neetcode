func climbStairs(n int) int {
    if n <= 1 { return 1 }
    a, b := 1, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
