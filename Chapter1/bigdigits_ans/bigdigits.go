package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    if len(os.Args) == 1 || os.Args[1] == "-h" ||
        os.Args[1] == "--help" ||
        (len(os.Args) == 2 && (os.Args[1] == "-b" ||
            os.Args[1] == "--bar")) {
        fmt.Printf("usage: %s [-b|--bar] <whole-number>\n"+
            "-b --bar  draw an underbar and an overbar\n",
            filepath.Base(os.Args[0]))
        os.Exit(1)
    }

    bar := false
    var stringOfDigits string
    if os.Args[1] == "-b" || os.Args[1] == "--bar" {
        bar = true
        stringOfDigits = os.Args[2]
    } else {
        stringOfDigits = os.Args[1]
    }
    for row := range bigDigits[0] {
        line := ""
        for column := range stringOfDigits {
            digit := stringOfDigits[column] - '0'
            if 0 <= digit && digit <= 9 {
                line += bigDigits[digit][row]
                if column+1 < len(stringOfDigits) {
                    line += "  "
                }
            } else {
                log.Fatal("invalid whole number")
            }
        }
        if bar && row == 0 {
            fmt.Println(strings.Repeat("*", len(line)))
        }
        fmt.Println(line)
        if bar && row+1 == len(bigDigits[0]) {
            fmt.Println(strings.Repeat("*", len(line)))
        }
    }
}

var bigDigits = [][]string{
    {"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ",
        "  000  "},
    {" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
    {" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
    {" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
    {"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
        "   4  "},
    {"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
    {" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
    {"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
    {" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
    {" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}