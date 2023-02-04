type Data struct {
	Arr [26]int
	Len int
}

func (d *Data) Update(add, del byte) {
	d.Arr[add-'a']++
	d.Arr[del-'a']--
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	arrS1 := new(Data)
	arrS2 := new(Data)
	arrS1.Len = len(s1)
	arrS2.Len = len(s2)
	for i := 0; i < len(s1); i++ {
		arrS1.Arr[s1[i]-'a']++
		arrS2.Arr[s2[i]-'a']++
	}

	for i := 0; i < arrS2.Len-arrS1.Len; i++ {
		if matchs(arrS1.Arr, arrS2.Arr) {
			return true
		}
		arrS2.Update(s2[i+arrS1.Len], s2[i])
	}
	return matchs(arrS1.Arr, arrS2.Arr)
}

func matchs(a1, a2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
