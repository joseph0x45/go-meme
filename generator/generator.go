package generator

import "net/http"

func Generate(){
    url := "https://api.imgflip.com/caption_image"
    data := map[string]interface{}{
        "template_id":"",
        "username":"",
        "password":"",
        "text0":"one does not simply",
        "text1":"learn Rust's basics in one night",
    }
    response, err := http.Post(
        url,
        "application/json",
        nil,
    )

}
