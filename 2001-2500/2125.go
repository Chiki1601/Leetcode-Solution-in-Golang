func numberOfBeams(bank []string) int {
    res,k,p :=0,0,0
    for _,row := range bank{
        if !strings.Contains(row, "1"){
            continue 
        }
        for _,char:= range row{
            if char=='1'{
                k+=1
            }
        } 
        res+=k*p
        p=k
        k=0
    }
    return res
}
