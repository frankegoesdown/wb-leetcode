package implement_strstr

func strStr(haystack string, needle string) (res int) {
    if len(needle) == 0 { 
        return 
    }
    for i := 0; i <= len(haystack) - len(needle); i++ {
        if haystack[i:i+len(needle)] == needle { 
            res = i
            return
        }
    }
    res = -1 
    return
}
