package main

func doesAliceWin(s string) bool {
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c=='a' || c=='e' || c=='i' || c=='o' || c=='u' {
            return true
        }
    }
    return false
}
