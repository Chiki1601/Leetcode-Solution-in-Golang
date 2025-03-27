func minimumIndex(nums []int) int {
    findDominantElement := func(arr []int) int {
        candidate, count := -1, 0

        // Boyer-Moore Majority Vote algorithm
        for _, num := range arr {
            if count == 0 {
                candidate = num
                count = 1
            } else if num == candidate {
                count++
            } else {
                count--
            }
        }

        totalCount := 0
        for _, num := range arr {
            if num == candidate {
                totalCount++
            }
        }

        if totalCount > len(arr)/2 {
            return candidate
        }
        return -1
    }

    dominant := findDominantElement(nums)
    if dominant == -1 {
        return -1
    }

    leftCount, totalDominantCount := 0, 0
    for _, num := range nums {
        if num == dominant {
            totalDominantCount++
        }
    }

    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == dominant {
            leftCount++
        }

        leftSubarrayCount := leftCount
        rightSubarrayCount := totalDominantCount - leftCount

        if leftSubarrayCount > (i+1)/2 && rightSubarrayCount > (len(nums)-i-1)/2 {
            return i
        }
    }

    return -1
}
