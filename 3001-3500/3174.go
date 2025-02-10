package main

import (
    "fmt"
    "strings"
    "unicode"
)

func clearDigits(input string) string {
    clearDigits := []string{}

    for _, current := range input {
        if unicode.IsDigit(current) && len(clearDigits) > 0 {
            clearDigits = clearDigits[:len(clearDigits)-1]
            continue
        }
        clearDigits = append(clearDigits, string(current))
    }

    return strings.Join(clearDigits, "")
}
