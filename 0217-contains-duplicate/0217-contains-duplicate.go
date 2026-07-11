func containsDuplicate(nums []int) bool {
    numbersMap := make(map[int]bool)
    for _, value := range nums {
        if (numbersMap[value]) {
            return true
        }
        numbersMap[value] = true
    }
    return false
}