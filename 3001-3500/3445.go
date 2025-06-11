
func maxDifference(s string, k int) int {
	freqPx := [5][]int{{0}, {0}, {0}, {0}, {0}}
	for i := range s {
		for j := range freqPx {
			freqPx[j] = append(freqPx[j], 0)
			freqPx[j][len(freqPx[j])-1] = freqPx[j][len(freqPx[j])-2]
		}
		freqPx[s[i]-48][len(freqPx[s[i]-48])-1]++
	}

	rs := -math.MaxInt32
	fn := func(i0, i1 int) int {
		mx := -math.MaxInt32
		ar01 := [4][][2]int{}
		for i := k; i < len(freqPx[0]); i++ {
			
			if freqPx[i0][i]%2 == 0 && freqPx[i1][i]%2 == 0 {
				ar01[0] = append(ar01[0], [2]int{freqPx[i0][i] - freqPx[i1][i], i})
			}
			
			if freqPx[i0][i]%2 == 0 && freqPx[i1][i]%2 == 1 {
				ar01[1] = append(ar01[1], [2]int{freqPx[i0][i] - freqPx[i1][i], i})
			}
			
			if freqPx[i0][i]%2 == 1 && freqPx[i1][i]%2 == 0 {
				ar01[2] = append(ar01[2], [2]int{freqPx[i0][i] - freqPx[i1][i], i})
			}
			
			if freqPx[i0][i]%2 == 1 && freqPx[i1][i]%2 == 1 {
				ar01[3] = append(ar01[3], [2]int{freqPx[i0][i] - freqPx[i1][i], i})
			}
		}
		for q := range ar01 {
			sort.Slice(ar01[q], func(i, j int) bool {
				return ar01[q][i][0] < ar01[q][j][0]
			})
		}
		for i := 0; i < len(freqPx[0])-k; i++ {
			ar01 = removeArr(i, k, i0, i1, ar01, freqPx)
			
			if freqPx[i0][i]%2 == 0 && freqPx[i1][i]%2 == 0 {
				
				if len(ar01[1]) > 0 {
					mi := -ar01[1][0][0] + freqPx[i0][i] - freqPx[i1][i]
					mx = max(mx, mi)
				}
				
				if len(ar01[2]) > 0 {
					mi := ar01[2][len(ar01[2])-1][0] - freqPx[i0][i] + freqPx[i1][i]
					mx = max(mx, mi)
				}
			}
			
			if freqPx[i0][i]%2 == 0 && freqPx[i1][i]%2 == 1 {
				
				if len(ar01[0]) > 0 {
					mi := -ar01[0][0][0] + freqPx[i0][i] - freqPx[i1][i]
					mx = max(mx, mi)
				}
				
				if len(ar01[3]) > 0 {
					mi := ar01[3][len(ar01[3])-1][0] - freqPx[i0][i] + freqPx[i1][i]
					mx = max(mx, mi)
				}
			}
			
			if freqPx[i0][i]%2 == 1 && freqPx[i1][i]%2 == 0 {
				
				if len(ar01[0]) > 0 {
					mi := ar01[0][len(ar01[0])-1][0] - freqPx[i0][i] + freqPx[i1][i]
					mx = max(mx, mi)
				}
				
				if len(ar01[3]) > 0 {
					mi := -ar01[3][0][0] + freqPx[i0][i] - freqPx[i1][i]
					mx = max(mx, mi)
				}
			}
			
			if freqPx[i0][i]%2 == 1 && freqPx[i1][i]%2 == 1 {
				
				if len(ar01[1]) > 0 {
					mi := ar01[1][len(ar01[1])-1][0] - freqPx[i0][i] + freqPx[i1][i]
					mx = max(mx, mi)
				}
				
				if len(ar01[2]) > 0 {
					mi := -ar01[2][0][0] + freqPx[i0][i] - freqPx[i1][i]
					mx = max(mx, mi)
				}
			}
		}
		return mx
	}
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			rs = max(rs, fn(i, j))
		}
	}
	return rs
}
func removeArr(start, k, i0, i1 int, ar01 [4][][2]int, freqPx [5][]int) [4][][2]int {
	for i := range ar01 {
		for len(ar01[i]) > 0 && (ar01[i][0][1]-start < k || freqPx[i0][start] == freqPx[i0][ar01[i][0][1]] || freqPx[i1][start] == freqPx[i1][ar01[i][0][1]]) {
			ar01[i] = ar01[i][1:]
		}
		for len(ar01[i]) > 0 && (ar01[i][len(ar01[i])-1][1]-start < k || freqPx[i0][start] == freqPx[i0][ar01[i][len(ar01[i])-1][1]] || freqPx[i1][start] == freqPx[i1][ar01[i][len(ar01[i])-1][1]]) {
			ar01[i] = ar01[i][:len(ar01[i])-1]
		}
	}
	return ar01
}
