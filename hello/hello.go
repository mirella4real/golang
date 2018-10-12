package main

import "fmt"

const helloPrefixEnglish =  "Hello, "
const helloPrefixSpanish = "Hola, " 
const helloPrefixFrench = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
    if name == ""{
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { 

    switch language {
        case spanish:
            prefix =  helloPrefixSpanish
        case french:
            prefix =  helloPrefixFrench
        default:
            prefix = helloPrefixEnglish
    }

    return prefix
}

func main(){
    fmt.Println(Hello("Mirella", ""))
}