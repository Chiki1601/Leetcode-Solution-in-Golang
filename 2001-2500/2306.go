func distinctNames(ideas []string) (result int64) {
    suffix:=map[string][26]bool{} // keeps a map of suffix to prefix letter
    count:=[26]int{} // keeps a count of each letter frequency in first letter of each idea
    same:=[26][26]int{} //keeps a count of shared prefixes
    val:=0
    s:=""
    for _,idea:=range ideas{
        s=string(idea[1:])
        val=int(idea[0]-'a')
        count[val]++
        if _,ok:=suffix[s];!ok{
            suffix[s]=[26]bool{}
        }
        chars:=suffix[s]
        chars[val]=true
        suffix[s]=chars
    }
    for _,chars:=range suffix{
        for i:=range chars{
            if !chars[i]{ continue }
            for j:=i+1;j<26;j++{
                if !chars[j]{ continue }
                same[i][j]++
            }
        }
    }
    for i:=range count{
        if (count[i] == 0) {continue}
        for j:=i+1;j<26;j++{
            if (count[j] == 0) {continue}
            val = (count[i] - same[i][j]) * (count[j] - same[i][j]);
            result += int64(val)
        }
    }
    return result*2 // to identify symmetrical solutions under j,i conditions
}
