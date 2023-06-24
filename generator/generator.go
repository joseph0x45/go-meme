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
        println(err)
        return
    }

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        println(err)
        return
    }

    defer response.Body.Close()

    response_body, err := ioutil.ReadAll(response.Body)
    var response_data interface{}
    if err:= json.Unmarshal(response_body, &response_data); err!=nil {
        println(err)
        return
    }
    if meme_data, ok := response_data.(map[string]interface{}); ok {
        fmt.Printf("%+v", meme_data)
        if data, ok := meme_data["data"].(map[string]interface{}); ok {
            if url, ok := data["url"].(string); ok {
                response, err := http.Get(url)
                if err != nil {
                    println(err)
                    return
                }
                defer response.Body.Close()
                
                if response.StatusCode != http.StatusOK {
                    println("Error while downloading image")
                    return
                }
                image, err := ioutil.ReadAll(response.Body)
                if err != nil {
                    println("Error while reading image bytes")
                    return
                }
                err = ioutil.WriteFile( "./random_meme_generated_by_gomeme.jpg" , image, 0644)
                if err != nil {
                    println("Error while saving image")
                    println(err)
                    return
                }
                println("Your meme has been served. Enjoy")
            }
        }else {
            println("Error while extracting data")
            return
        }
    }else {
        println("Error while performing type assertion")
        return
    }
}
