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
    ans, sum, prevsum := 0, 0, 0

    for scanner.Scan() {
        i, _ := strconv.Atoi(scanner.Text())
        l = append(l, i)
    }
    for i := 0; i < len(l)-2; i++ {
        sum = l[i] + l[i+1] + l[i+2]
        
        if sum > prevsum {
            ans  += 1
        }
        prevsum = sum
    }
    fmt.Println("answer: ", ans)
}
