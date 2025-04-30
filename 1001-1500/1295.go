func findNumbers(nums []int) int {
    output :=0 
    for _,v := range nums{
        if len(strconv.Itoa(v))%2==0 {
            output++
        }
    }
    return output 
}
