package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {


  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\nInput Text: ")
    text, _ := reader.ReadString('\n')
    //text, _ := reader.ReadString('.') //read until meet .

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "--> ")
	fmt.Fprint(w, text)

	w.Flush()
	}



}
