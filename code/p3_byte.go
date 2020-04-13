package main

import "fmt"

func byte_operation(){
  // byte
  var p7 byte = 66
  p7 -=1;
  if p7 == 'A'{
      fmt.Println("They have the same value")
  }
  fmt.Printf("%c\n",p7)
}
func main() {
   byte_operation()
}