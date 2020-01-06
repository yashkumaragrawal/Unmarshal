Qu. 1312. Minimum Insertion Steps to Make a String Palindrome------

Given a string s. In one step you can insert any character at any index of the string.

Return the minimum number of steps to make s palindrome.

A Palindrome String is one that reads the same backward as well as forward.


Ans --- 

func min(i, j int) int {
    if i > j {
        return j
    }
    return i
}
func minInsertions(s string) int {
    n := len(s)
    a := make([][]int , n)
    for i:=0; i < n; i++ {
        a[i] = make([]int, n)
    }
    
    for i := 1; i < n; i++ {
        h := i
        for l := 0; h < n; l++ {
            if s[l] == s[h] {
                a[l][h] = a[l+1][h-1]
            }else {
                a[l][h] = min(a[l][h-1],a[l+1][h]) + 1
            }
            h = h+1
        }
    } 
    return a[0][n-1]
}
