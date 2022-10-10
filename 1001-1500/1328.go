func breakPalindrome(palindrome string) string {
 
    l:=len(palindrome)
    if l==1{
        return ""
    }
    
    for i:=0;i<l/2;i++{
        if palindrome[i]!='a'{
            return palindrome[:i] + "a" + palindrome[i+1:]
        }
    }
    return palindrome[:l-1] + "b" 
}
