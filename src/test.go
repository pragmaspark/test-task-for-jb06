package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "encoding/json"
)

type JsonStruct struct {
    Action string  `json:"action"`
    Task   string  `json:"task,omitempty"`
    Text   string  `json:"text,omitempty"`
}


func ParsingStrings (lines []string) []JsonStruct {
    result := []JsonStruct{}

    for i := 0; i < len(lines); i++ {
        a := JsonStruct{}
        err := json.Unmarshal([]byte(lines[i]),&a)
        if err != nil {
            log.Fatal("Parsing error")
        }
        result = append(result, a)
    }
    return result
}


func PrintAction (jsonLines []JsonStruct) {
    for i := 0; i < len(jsonLines); i++ {
        //ParsingJsonFields(a[i])
        switch a := jsonLines[i]; a.Action {
        case "start":
            fmt.Println(">", a.Task)
        case "message":
            fmt.Println("#", a.Text)
        default:
            fmt.Println(`Wrong json "action" field :`, a.Action)
        }
    }
}


// собираем данные с аргументов
func ParseCommand () []string {
    lines := []string{}
    for i := 1; i < len(os.Args); i++ {
        lines = append(lines, os.Args[i])
    }
    return lines
}


// собираем данные с консоли до тех пор пока не встретим символ "q"
func ParseCommandLine () []string {
    lines := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        text := scanner.Text()
        if scanner.Text() == "q" {
            break
        }
        lines = append(lines, text)
    }
    return lines
}

func main () {
    lines := []string{}

    argLines := ParseCommand()
    for i := 0; i < len(argLines); i++ {
        lines = append(lines, argLines[i])
    }

    cmdLines := ParseCommandLine()
    for i := 0; i < len(cmdLines); i++ {
        lines = append(lines, cmdLines[i])
    }

//     lines := []string{
//       `{ "action": "start", "task": "abc"}`,
//       `{"action":"message", "text": "123"}`}

    jsonLines := ParsingStrings(lines)
    PrintAction(jsonLines)
}
