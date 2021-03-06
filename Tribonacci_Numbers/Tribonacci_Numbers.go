/*
TRIBONACCI NUMBERS

It is a series of numbers where each term is the sum of preceeding 3
numbers of the sequence.

T(n) = T(n-1) + T(n-2) + T(n-3)
*/

package main

import "fmt"

func tribonacci(n int) int{
    x := n - 3
    if x < 0 {
        return 0
    }
    if x == 0 {
        return 1
    }
    return (tribonacci(n - 1) + tribonacci(n - 2) + tribonacci(n - 3))
}

func main() {
    var n int
    fmt.Print("Enter N: ")
    fmt.Scan(&n)
    fmt.Print("The tribonacci number at position ", n, " is: ", tribonacci(n))

}

/*
INPUT : n = 6
OUTPUT: The tribonacci number at position 6 is: 4
VERIFICATION: The tribonacci series : 0 0 1 1 2 4
*/
