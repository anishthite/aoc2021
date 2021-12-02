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

    x, aim, y := 0, 0, 0

    for scanner.Scan() {
        res := strings.Split(scanner.Text(), " ")
        fmt.Println(res)
        switch res[0] {
        case "down":
            if n, err := strconv.Atoi(res[1]); err == nil {
                aim += n
            }
        case "up":
            if n, err := strconv.Atoi(res[1]); err == nil {
                aim -= n
            }
        case "forward":
            if n, err := strconv.Atoi(res[1]); err == nil {
                x += n
                y += (n * aim)
            }
        case "backward":
            if n, err := strconv.Atoi(res[1]); err == nil {
                x -= n
            }

        }
    }
    fmt.Println(x * y)

}
