package main

import "testing"

func TestHello(t *testing.T){

    assertCorrectMessage := func(t *testing.T, got, want string){
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T){
        got := Hello("Mirella", "")
        want := "Hello, Mirella"

        assertCorrectMessage(t, got, want)
    })
    
    t.Run("say 'Hello, World' when an empty strin is supplied", func(t *testing.T){
        got := Hello("", "")
        want := "Hello, World"

        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T){
        got := Hello("Mirella", "Spanish")
        want := "Hola, Mirella"

        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T){
        got := Hello("Mirella", "French")
        want := "Bonjour, Mirella"

        assertCorrectMessage(t, got, want)
    })
    
}