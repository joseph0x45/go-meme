package main

import (
	"bufio"
	"fmt"
	// "go-meme/generator"
	"os"
	"strconv"
)

func main(){
    type Template struct{
        number string
        id string
        name string
        number_of_captions int
    }
    templates := []Template{
        {
            number: "1",
            id: "61579",
            name: "One does not simply",
            number_of_captions: 2,
        },
        {
            number: "2",
            id: "112126428",
            name: "Distracted boyfriend",
            number_of_captions: 2,
        },
        {
            number: "3",
            id: "4087833",
            name: "Waiting skeleton",
            number_of_captions: 2,
        },
    }
    input_reader := bufio.NewScanner(os.Stdin)
    println("Welcome to Gomeme, a goofy meme generator written in the Goofiest language in the world!")
    println("Find below the list of all available templates and chose the one you want\n")
    for _, value := range templates{
        fmt.Printf("%s - %s\n", value.number, value.name)
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
    fmt.Printf("You chose %d", chosen_template)
}
