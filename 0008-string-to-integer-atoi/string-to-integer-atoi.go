func myAtoi(str string) (sesult int) {
    str = strings.TrimLeft(str, " ")
    str = strings.TrimSpace(str)
    sign := 1
    
    if len(str) == 0 {
        return 0
    }
    if str[0] == '-' {
        sign = -1
        str = str[1:]
    } else if str[0] == '+' {
        str = str[1:]
    }
    for i := 0; i < len(str); i++ {
        c := str[i]
        if c >= '0' && c <= '9' {
            sesult = sesult * 10 + (int(c - '0') * sign)

            if sesult <= math.MinInt32 {
                return math.MinInt32
            } else if sesult >= math.MaxInt32 {
                return math.MaxInt32
            }
        } else {
            break
        }
    }

    return 
}