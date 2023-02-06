func isPossibleToCutPath(grid [][]int) bool {
    n := len(grid)+2
    m := len(grid[0])+2

    dp1 := [][]int{}
    for i:=0;i<n;i++{
        dp1 = append(dp1,make([]int,m))
    }
    dp1[1][1] = 1

    for i:=1;i<n-1;i++{
        for j:=1;j<m-1;j++{
            if i == 1 && j == 1{
                continue
            }
            if grid[i-1][j-1] == 0{
                continue
            }
            dp1[i][j] = dp1[i-1][j]+dp1[i][j-1]
        }
    }
        
    if dp1[n-2][m-2] == 0{
        return true
    }
                
    dp2 := [][]int{}
    for i:=0;i<n;i++{
        dp2 = append(dp2,make([]int,m))
    }
    dp2[n-2][m-2] = 1
    for i:=n-2;i>0;i--{
        for j:=m-2;j>0;j--{
            if i == n-2 && j == m-2{
                continue
            } 
            if grid[i-1][j-1] == 0{
                continue
            }
            dp2[i][j] = dp2[i+1][j]+dp2[i][j+1]
        }
    }            
    
    target := dp1[n-2][m-2]

    for i:=1;i<n-1;i++{
        for j:=1;j<m-1;j++{
            if i == 1 && j == 1{
                continue
            }
            if i == n-2 && j == m-2{
                continue
            }
            if grid[i-1][j-1] == 0{
                continue
            }
            if dp1[i][j] * dp2[i][j] == target{
                return true
            }
        }
    }
    return false
}
