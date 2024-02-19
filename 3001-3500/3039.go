func lastNonEmptyString(s string) string {
    letters := make(map[rune][]int)
    max_counter := 0
    for i,letter := range(s){
        if _,ok := letters[letter]; !ok{
            letters[letter] = []int{i,0}   
        }
        letters[letter][0] = i
        letters[letter][1]++
        max_counter = max(max_counter,letters[letter][1])   
    }

    ans := []int{}

    for _,value := range(letters){
        idx := value[0]
        count := value[1]
        if count == max_counter{
            ans = append(ans,idx)     
        } 
    }

    sort.Ints(ans)
    answer := []byte{}
    for _,idx := range(ans){
        answer = append(answer,s[idx])
    }
    
    return string(answer)
}
