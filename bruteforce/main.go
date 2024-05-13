package main

import (
  "fmt"
)

var Alphabet = []string{"0", "1"}

func bruteforce(word string, limit int) {

  if limit == 0 {
    fmt.Println(word)
  } else {
    for i, _ := range Alphabet {
      //fmt.Println(word + Alphabet[i])
      //time.Sleep(1 * time.Second)
      bruteforce(word+Alphabet[i], limit-1)
    }
  }

}


func main() {
  word := ""
  bruteforce(word, 5)
}
