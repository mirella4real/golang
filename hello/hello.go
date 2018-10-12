package main

import "fmt"

const helloPrefix =  "Hello, "
const helloPrefixSpanish = "Hola, " 
const helloPrefixFrechn = "Bonjour, "

func Hello(name string, language string) string {
    if name == ""{
        name = "World"
    }

    prefix := helloPrefix

    switch language {
        case "Spanish":
            prefix =  helloPrefixSpanish
        case "French":
            prefix =  helloPrefixFrechn
    }

    return prefix + name
}

func main(){
    fmt.Println(Hello("Mirella", ""))
}