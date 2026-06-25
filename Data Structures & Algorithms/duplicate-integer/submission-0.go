func hasDuplicate(nums []int) bool {
    note := make(map[int]bool)
    for _, v := range nums {
        if note[v] {
            return true
        }
        note[v] = true
    }
    return false
}
