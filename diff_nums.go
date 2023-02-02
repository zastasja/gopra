//По данному трехзначному числу определите, все ли его цифры различны.
package main
import (
  "fmt"
)

func main(){
    var num int
    fmt.Scan(&num)
  
    one:= num/100
    two:= (num/10)%10
    three:= num%10
    
    if one == two || two == three || three == one{
        fmt.Print("NO")
    } else{fmt.Print("YES")}
}
