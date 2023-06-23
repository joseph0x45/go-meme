package main

import (
	"bufio"
	"fmt"
	"go-meme/generator"
	"go-meme/shared"
	"os"
	"strconv"
)

func main(){
    templates := []shared.Template{
        {
            Id: "61579",
            Name: "One does not simply",
        },
        {
            Id: "112126428",
            Name: "Distracted boyfriend",
        },
        {
            Id: "4087833",
            Name: "Waiting skeleton",
        },
    }
    input_reader := bufio.NewScanner(os.Stdin)
    println("Welcome to Gomeme, a goofy meme generator written in the Goofiest language in the world!")
    println("Find below the list of all available templates and chose the one you want\n")
    i := 0
    for _, value := range templates{
        fmt.Printf("%d - %s\n", i, value.Name)
        i++
    }
    print("Chose your meme template : ")
    input_reader.Scan()
    chosen_template, err := strconv.Atoi(input_reader.Text())
    if err != nil {
        println("You have to enter the number in front of the desired meme template.")
        println("For example, enter 1 for the 'One does not simply' template")
        return
    }
    if chosen_template > len(templates)-1 || chosen_template < 0 {
        println("Invalid template number.")
        return
    }
    fmt.Printf("You chose the %s\n", templates[chosen_template].Name)
    print("Enter the top text: ")
    input_reader.Scan()
    text0 := input_reader.Text()
    print("Enter the bottom text: ")
    input_reader.Scan()
    text1 := input_reader.Text()
    generator.Generate(templates[chosen_template], text0, text1)
}
