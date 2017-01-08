package main

import "fmt"
// import "bufio"
import "log"
import "os/exec"

func main() {
    fmt.Printf("hello, world\n")
    fmt.Printf("this is a command: <cmd>\n")

    out, err := exec.Command("whoami").Output()
    if err != nil {
	log.Fatal(err)
    }
    fmt.Printf("I am %s\n", out)

    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Enter some txt :::") 


}
