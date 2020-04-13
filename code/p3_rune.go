package main

import "fmt"

func rune_operation(){
  // byte
  var p7 rune = 66
  p7 -=1;
  if p7 == 'A'{
      fmt.Println("They have the same value")
  }
  fmt.Printf("%c\n",p7)
}
func main() {
   rune_operation()
}