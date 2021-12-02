package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)



func main() {

    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    x, y := 0, 0

    for scanner.Scan() {
        res := strings.Split(scanner.Text(), " ")
        fmt.Println(res)
        switch res[0] {
        case "down":
            if n, err := strconv.Atoi(res[1]); err == nil {
                y += n
            }
        case "up":
            if n, err := strconv.Atoi(res[1]); err == nil {
                y -= n
            }
        case "forward":
            if n, err := strconv.Atoi(res[1]); err == nil {
                x += n
            }
        case "backward":
            if n, err := strconv.Atoi(res[1]); err == nil {
                x -= n
            }

        }
    }
    fmt.Println(x * y)

}
