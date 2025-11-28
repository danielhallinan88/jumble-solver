package main

import (
    "bufio"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

func main() {
    dict, err := loadDictionary("words_alpha.txt")
    if err != nil {
        panic(err)
    }

    port := "8082"
    if len(os.Args) > 1 {
        port = os.Args[1]
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
        path := strings.Trim(r.URL.Path, "/")
        if path == "" {
            fmt.Fprintln(w, "Hello! Use /<jumble> to solve a word.")
            return
        }

        w.WriteHeader(http.StatusOK)
        foundStrings := solveJumble(path, dict)
        log.Printf("INFO: Found permutations for %s: %s", path, foundStrings)
        fmt.Fprintln(w, "Found permutations for", path, ": ", foundStrings)
    })

    fmt.Printf("Server listening on http://localhost:%s\n", port)
    http.ListenAndServe(":"+port, nil)
}

func solveJumble(jumble string, dict map[string]bool) ([]string) {
    var validPerms []string
    var resultSet = make(map[string]bool)

    for _, p := range permutations(jumble) {
        if dict[p] {
            resultSet[p] = true
        }
    }

    for word := range resultSet {
        validPerms = append(validPerms, word)
    }
    return validPerms
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


