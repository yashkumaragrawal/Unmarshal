Qu.  Decrypt String from Alphabet to Integer Mapping--

 Given a string s formed by digits ('0' - '9') and '#' . We want to map s to English lowercase characters as follows:

    Characters ('a' to 'i') are represented by ('1' to '9') respectively.
    Characters ('j' to 'z') are represented by ('10#' to '26#') respectively. 

Return the string formed after mapping.

It's guaranteed that a unique mapping will always exist.

Ans:-----
 
 func isdigit(b byte) bool {
    if b >= 48 && b <= 57 {
        return true
    }
    return false
}
func freqAlphabets(s string) string {
    str := ""
    m := make(map[string]string)
    m["1"] = "a"
    m["2"] = "b"
    m["3"] = "c"
    m["4"] = "d"
    m["5"] = "e"
    m["6"] = "f"
    m["7"] = "g"
    m["8"] = "h"
    m["9"] = "i"
    m["10"] = "j"
    m["11"] = "k"
    m["12"] = "l"
    m["13"] = "m"
    m["14"] = "n"
    m["15"] = "o"
    m["16"] = "p"
    m["17"] = "q"
    m["18"] = "r"
    m["19"] = "s"
    m["20"] = "t"
    m["21"] = "u"
    m["22"] = "v"
    m["23"] = "w"
    m["24"] = "x"
    m["25"] = "y"
    m["26"] = "z"
    for i := 0; i< len(s); i++ {
        if isdigit(s[i]) == true {
            j := i+2
            if j < len(s) && isdigit(s[j]) == false {
                str = str + m[s[i:i+2]]
                i = i+2
            }else {
                str = str + m[s[i:i+1]]
            }
        }
    }
    return str
}
