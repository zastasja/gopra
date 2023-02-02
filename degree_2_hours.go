//Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
package main
import ("fmt")
 
func main() {
  var degr int
  fmt.Scan(&degr) 
  minute := degr*2
  h := minute/60
  m := minute%60
  fmt.Print("It is ",h," hours ",m," minutes.")
}
