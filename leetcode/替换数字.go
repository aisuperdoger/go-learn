package leetcode



import "fmt"

func main(){
    var strByte []byte
    
    fmt.Scanln(&strByte)
    
    for i := 0; i < len(strByte); i++{
        if strByte[i] <= '9' && strByte[i] >= '0' {
            inserElement := []byte{'n','u','m','b','e','r'}
			// 使用 strByte[i+1:]... 就告诉编译器：“请把这个切片里的每一个元素拿出来，当作独立的参数传给 append”。
            strByte = append(strByte[:i], append(inserElement, strByte[i+1:]...)...)
            i = i + len(inserElement) -1
        }
    }
    
    fmt.Printf(string(strByte))
}