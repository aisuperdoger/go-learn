package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func scanner_test() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		sc.Scan()
		arr[i],_ = strconv.Atoi(sc.Text())
		if i!= 0 {
			arr[i] +=arr[i-1]
		}
	}

	for {
		var l,r int 
		sc.Scan()
		_,err := fmt.Sscanf(sc.Text(),"%d %d",&l,&r)
		if err !=nil {
			return
		}

		if l >0 {
			fmt.Println(arr[r] - arr[l-1])
		}else{
			fmt.Println(arr[r])
		}
	}
}



func main() {
    var n, m int
    
    reader := bufio.NewReader(os.Stdin)
    
    line, _ := reader.ReadString('\n')
    line = strings.TrimSpace(line)
    params := strings.Split(line, " ")
    
    n, _ = strconv.Atoi(params[0])
    m, _ = strconv.Atoi(params[1])//n和m读取完成
    
    land := make([][]int, n)//土地矩阵初始化
    
    for i := 0; i < n; i++ {
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        values := strings.Split(line, " ")
        land[i] = make([]int, m)
        for j := 0; j < m; j++ {
            value, _ := strconv.Atoi(values[j])
            land[i][j] = value
        }
    }//所有读取完成
    
    //初始化前缀和矩阵
    preMatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		preMatrix[i] = make([]int, m+1)
	}
    
    for a := 1; a < n+1; a++ {
        for b := 1; b < m+1; b++ {
            preMatrix[a][b] = land[a-1][b-1] + preMatrix[a-1][b] + preMatrix[a][b-1] - preMatrix[a-1][b-1]
        }
    }
    
    totalSum := preMatrix[n][m]
    
    minDiff := math.MaxInt32//初始化极大数，用于比较
    
    //按行分割
    for i := 1; i < n; i++ {
        topSum := preMatrix[i][m]
        
        bottomSum := totalSum - topSum
        
        diff := int(math.Abs(float64(topSum - bottomSum)))
        if diff < minDiff {
            minDiff = diff
        }
    }
    
    //按列分割
    for j := 1; j < m; j++ {
        topSum := preMatrix[n][j]
        
        bottomSum := totalSum - topSum
        
        diff := int(math.Abs(float64(topSum - bottomSum)))
        if diff < minDiff {
            minDiff = diff
        }
    }    
    
    fmt.Println(minDiff) 
}
// https://www.programmercarl.com/kamacoder/0058.%E5%8C%BA%E9%97%B4%E5%92%8C.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
// https://www.programmercarl.com/kamacoder/0044.%E5%BC%80%E5%8F%91%E5%95%86%E8%B4%AD%E4%B9%B0%E5%9C%9F%E5%9C%B0.html#%E6%80%9D%E8%B7%AF