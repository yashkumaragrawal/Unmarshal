Qu. An integer has sequential digits if and only if each digit in the number is one more than the previous digit.

Return a sorted list of all the integers in the range [low, high] inclusive that have sequential digits.

Ans----

func sequentialDigits(low int, high int) []int {
    i := 10
    a := make([]int, 0)
    for i <= high {
        k := 0
        for j:= 1; k < i * 10 && j < 10; j++ {
            l := i
            k = 0
            n := 0
            for l > 0 && j + n < 10 {
                k = k*10 + j + n
                n++
                l = l/10
            }
            if k >= low && k <= high && k > i{
                a = append(a, k)
            }
        }
        i = i * 10
    }
    return a
}
