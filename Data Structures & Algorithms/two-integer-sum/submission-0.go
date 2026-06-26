func twoSum(nums []int, target int) []int {
    prevNote := make(map[int]int)

    for i, n := range nums {
        diff := target - n
        if j, found := prevNote[diff]; found {
            return []int{j, i}
        }
        prevNote[n] = i
    }
    return []int{}
}