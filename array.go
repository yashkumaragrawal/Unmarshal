Qu. XOR Queries of a Subarray----

Given the array arr of positive integers and the array queries where queries[i] = [Li, Ri], for each query i compute the XOR of elements from Li to Ri (that is, arr[Li] xor arr[Li+1] xor ... xor arr[Ri] ). Return an array containing the result for the given queries. 

Ans --- 
func xorQueries(arr []int, q [][]int) []int {
    if len(arr) < 1 {
        return arr
    }
    n := len(arr)
    m := len(q)
    a := make([]int, m)
    b := make([]int, n)
    b[0] = arr[0]
    for i:=1; i<n; i++ {
        b[i] = b[i-1] ^ arr[i]
    }
    for i:=0; i<m; i++ {
        if q[i][0] != 0 {
            a[i] = b[q[i][1]]^b[q[i][0] - 1]
        }else {
            a[i] = b[q[i][1]]
        }
    }
    return a
}
