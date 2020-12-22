/**
 * Map dispatch demo
 */
package main
import (
    "fmt"
    "errors"
    "strings"
    "bufio"
    "io"
    "os"
)

func taskOne() {
    // todo more
    fmt.Println("To do one")
}

func taskTwo() {
    fmt.Println("To do two")
}

func taskThree() {
    fmt.Println("To do three")
}

func taskFour() {
    fmt.Println("To do four")
}

func dispatch(choice string) error{
    choice = strings.ToLower(choice)
    tasks := map[string]func () {
        "a" : taskOne,
        "b" : taskTwo,
        "c" : taskThree,
        "d" : taskFour,
    }
    fn, ok := tasks[choice]
    if !ok {
        return errors.New("Not valid choice["+choice + "]")
    }
    fn()
    return nil
}

func dialogWithStdin() {
    fmt.Println("=======dialogWithStdin======")
    in := bufio.NewReader(os.Stdin)
    for {
        r, _, err := in.ReadRune() // returns rune, nbytes, error
        if err != nil { // 跳出主goroutine
            fmt.Println("===========\nWelcom use me!")
            os.Exit(1)
        }
        dispatch(string([]rune{r}))
    }
}