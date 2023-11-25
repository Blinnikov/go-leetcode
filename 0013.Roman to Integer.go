func romanToInt(s string) int {
    res := 0
    var prev string

    for _, r := range s {
        ch := string(r)
        
        switch ch {
            case "I": res++
            case "V":  if prev == "I" { res += 3 } else { res += 5 }
            case "X":  if prev == "I" { res += 8 } else { res += 10 }
            case "L":  if prev == "X" { res += 30 } else { res += 50 }
            case "C":  if prev == "X" { res += 80 } else { res += 100 }
            case "D":  if prev == "C" { res += 300 } else { res += 500 }
            case "M":  if prev == "C" { res += 800 } else { res += 1000 }
        }
        
        prev = ch
    }

    return res
}
