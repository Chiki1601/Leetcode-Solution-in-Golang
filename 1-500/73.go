func setZeroes(matrix [][]int)  {
    m,n:=len(matrix),len(matrix[0])
    frzero:=false
    fczero:=false
    for j:=0;j<n;j++{
        if matrix[0][j]==0{
            frzero=true
            break
        }
    }
    for i:=0;i<m;i++{
        if matrix[i][0]==0{
            fczero=true
            break
        }
    }
    for i:=1;i<m;i++{
        for j:=1;j<n;j++{
            if matrix[i][j]==0{
                matrix[i][0]=0
                matrix[0][j]=0
            }
        }
    }
    for i:=1;i<m;i++{
        for j:=1;j<n;j++{
            if matrix[i][0]==0 || matrix[0][j]==0{
                matrix[i][j]=0
            }
        }
    }
    if frzero{
        for j:=0;j<n;j++{
            matrix[0][j]=0
        }
    }
    if fczero{
        for i:=0;i<m;i++{
            matrix[i][0]=0
        }
    }

}
