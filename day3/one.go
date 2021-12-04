package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    
    numlist := make([]string, 0)
    for scanner.Scan() {
        numlist = append(numlist, scanner.Text());
    }

    startswith := ""
    cont := true
    var num_sel int
    for i := 0; cont ; i++ {
        fmt.Println(i) 
        diff := pick_index(numlist, i, startswith)
        //positive more one, neg zeros 
        if diff < 0 {
            startswith += "1"
        } else {
            startswith += "0"
        }
        num_sel = num_selected(numlist, startswith)
        
        if num_sel == 1 || i == 11 {
            fmt.Println(startswith)
            for _, num := range numlist {
                if strings.HasPrefix(num, startswith) {
                    fmt.Println(num)

                }
            }
            break
        }
    }
}

func num_selected(nums []string, startswith string) (int) {
    num_selected := 0
    chosen := make([]string, 0)
    for _, num := range nums {
        if strings.HasPrefix(num, startswith) || startswith == "" {
            num_selected += 1
            chosen = append(chosen, num)
        }
    }
    return num_selected
}

func pick_index(nums[] string, i int, prefix string) int{
    ones, zeros := 0, 0
    for _, num := range nums {
        if strings.HasPrefix(num, prefix) {
            if num[i] == '1' {
                ones += 1
            } 
            if num[i] == '0'{
                zeros += 1
            }
        }
    }
    return ones - zeros
} 












func give_nums(nums []string, i int, startswith string) (bool, string) {
    s := 0
    num_selected := 0 
    for _, num := range nums {
        if startswith == "" || strings.HasPrefix(num, startswith) {
            num_selected += 1
            if num[i] == '0' {
                s -= 1
            } else {
                s += 1
            }
        }
    }

    if num_selected == 1  || i == 11{
        for _, num := range nums {
            if strings.HasPrefix(num, startswith) {
                fmt.Println(num)
            }
        }
    }
 
    if s >= 0 {
        startswith += "1";
    } else {
        startswith += "0";
    }
    return num_selected > 1, startswith
}

func give_numsmin(nums []string, i int, startswith string) (bool, string) {
    s := 0
    num_selected := 0 
    for _, num := range nums {
        if startswith == "" || strings.HasPrefix(num, startswith) {
            num_selected += 1
            if num[i] == '0' {
                s -= 1
            } else {
                s += 1
            }
        }
    }

    if num_selected == 1 {
        for _, num := range nums {
            if strings.HasPrefix(num, startswith) {
                fmt.Println(num)
            }
        }
    }

    if s < 0 {
        startswith += "1";
    } else {
        startswith += "0";
    }
    return num_selected > 1, startswith
}
