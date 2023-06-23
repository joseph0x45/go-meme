package generator

import (
	"encoding/json"
	"fmt"
	"go-meme/shared"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Generate( template shared.Template, text0 string, text1 string ){
    println("Cooking......")
    uri := "https://api.imgflip.com/caption_image"
    form_data := url.Values{}
    form_data.Set("template_id", template.Id)
    form_data.Set("username", "pigeondev")
    form_data.Set("password", "N;QgK~TwDM&jgJ4")
    form_data.Set("text0", text0)
    form_data.Set("text1", text1)

    request, err := http.NewRequest(
        "POST",
        uri,
        strings.NewReader(form_data.Encode()),
    )
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    if err != nil {
        println("Something went wrong when generating your meme.")
        println("Please retry or report this bug on https://github.com/TheWisePigeon/go-meme")
        return
    }

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        println("Something went wrong when generating your meme.")
        println("Please retry or report this bug on https://github.com/TheWisePigeon/go-meme")
        return
    }

    defer response.Body.Close()

    response_body, err := ioutil.ReadAll(response.Body)
    var response_data map[string]interface{}
    if err:= json.Unmarshal(response_body, &response_data); err!=nil {
        println("Something went wrong when generating your meme.")
        println("Please retry or report this bug on https://github.com/TheWisePigeon/go-meme")
        return
    }
    fmt.Printf("%+v\n", response_data)
}
