package main

import (
	"bufio"
	"fmt"
	// "go-meme/generator"
	"os"
	"strconv"
)

func main(){
    type template struct{
        name string
        number_of_captions int
    }
    available_templates := map[string]string{
        "61579":"One does not simply",
        "112126428":"Distracted boyfriend",
        "4087833":"Waiting skeleton",
    }
    input_reader := bufio.NewScanner(os.Stdin)
    println("Welcome to Gomeme, a goofy meme generator written in the Goofiest language in the world!")
    println("Find below the list of all available templates and chose the one you want\n")
    i := 1
    for _, value := range available_templates{
        fmt.Printf("%d - %s\n", i, value)
        i++
    }
    print("Chose your meme template : ")
    input_reader.Scan()
    chosen_template, err := strconv.Atoi(input_reader.Text())
    if err != nil {
        println("You have to enter the number in front of the desired meme template.")
        println("For example, enter 0 for the 'One does not simply' template")
        return
    }
    if chosen_template > len(available_templates)-1 {
        println("Invalid template number.")
        return
    }
    fmt.Printf("You chose %d", chosen_template)
}
