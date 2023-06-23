package main

import (
	"bufio"
	"fmt"
	"go-meme/generator"
	"os"
	"strconv"
)

func main(){
    type Template struct{
        id string
        name string
        number_of_captions int
    }
    templates := []Template{
        {
            id: "61579",
            name: "One does not simply",
            number_of_captions: 2,
        },
        {
            id: "112126428",
            name: "Distracted boyfriend",
            number_of_captions: 2,
        },
        {
            id: "4087833",
            name: "Waiting skeleton",
            number_of_captions: 2,
        },
    }
    input_reader := bufio.NewScanner(os.Stdin)
    println("Welcome to Gomeme, a goofy meme generator written in the Goofiest language in the world!")
    println("Find below the list of all available templates and chose the one you want\n")
    i := 1
    for _, value := range templates{
        fmt.Printf("%d - %s\n", i, value.name)
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
    if chosen_template > len(templates) || chosen_template <= 0 {
        println("Invalid template number.")
        return
    }
    generator.Generate()
}
