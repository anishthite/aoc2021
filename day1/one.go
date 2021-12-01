package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main(){

    file, err := os.Open("one.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    var l []int
    ans := 0

    for scanner.Scan() {
        i, _ := strconv.Atoi(scanner.Text())
        l = append(l, i)
    }
    for i := 1; i < len(l); i++ {
        if l[i] > l[i-1] {
            ans  += 1
        }
    }
    fmt.Println("answer: ", ans)
}
