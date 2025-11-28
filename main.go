package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    dict, err := loadDictionary("words_alpha.txt")
    if err != nil {
        panic(err)
    }
    //fmt.Println("Dictionary loaded with", len(dict), "words")

    testStr := "eth"
    var validPerms []string
    for _, p := range permutations(testStr) {
        if dict[p] {
            validPerms = append(validPerms, p)
        }
        //fmt.Println(p, dict[p])
    }
    fmt.Println(validPerms)
}


func loadDictionary(path string) (map[string]bool, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    dict := make(map[string]bool)
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        dict[scanner.Text()] = true
    }

    return dict, scanner.Err()
}

func permutations(s string) []string {
    var res []string
    var helper func([]rune, int)

    helper = func(r []rune, i int) {
        if i == len(r)-1 {
            res = append(res, string(r))
            return
        }
        for j := i; j < len(r); j++ {
            r[i], r[j] = r[j], r[i]
            helper(r, i+1)
            r[i], r[j] = r[j], r[i]
        }
    }

    helper([]rune(s), 0)
    return res
}


