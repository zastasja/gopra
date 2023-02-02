// Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
package main
import (
  "fmt"
)

func main(){
    var a int
    fmt.Scan(&a)
    
    one:= a/100000
    two:= (a/10000)%10
    three:= (a/1000)%10
    fore:= (a/100)%10
    five:= (a/10)%10
    six:= a%10
    sum_l := one+two+three
    sum_r := fore+five+six
    
    if sum_l == sum_r {
        fmt.Print("YES")
    } else {fmt.Print("NO")}

}
