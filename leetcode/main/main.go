package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	line, _ := sc.ReadString('\n')

	line = strings.TrimSpace(line)
	params := strings.Split(line, " ")
	n, _ := strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[1])

	arr := make([][]int, n)

	for i, _ := range arr {
		arr[i] = make([]int, m)
		line, _ = sc.ReadString('\n')
		line = strings.TrimSpace(line)
		params = strings.Split(line, " ")

		for j, v := range params {
			arr[i][j], _ = strconv.Atoi(v)
		}
	}

	for i, _ := range arr {
		for j, v := range arr[i] {
			if 0 == j && 0 == i {
				arr[i][j] = v
			} else if 0 == i {
				arr[i][j] = arr[i][j-1] + v
			} else if 0 == j {
				arr[i][j] = arr[i-1][j] + v
			} else {
				arr[i][j] = arr[i][j-1] + arr[i-1][j] - arr[i-1][j-1] + v
			}
		}
	}

	ans := math.MaxInt64
	for i := 1; i <= n-1; i++ {
		tmp := int( math.Abs(float64(arr[n-1][m-1] - arr[i-1][m-1]- arr[i-1][m-1])))
		if tmp < ans {
			ans = tmp
		}
	}

	for i := 1; i <= m-1; i++ {
		tmp := int( math.Abs(float64(arr[n-1][m-1] - arr[n-1][i-1]- arr[n-1][i-1])))
		if tmp < ans {
			ans = tmp
		}
	}
	fmt.Println(ans)

}
