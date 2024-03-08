package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    // Create a new reader, assuming input will come from the standard input device
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter URL: ")
    url, err := reader.ReadString('\n') // Read the input until the first newline character
    if err != nil {
        fmt.Println("Error reading URL: ", err)
        return
    }

    // Trim the newline character from the URL, which is read as part of the input
    // This step is important because different operating systems have different newline characters
    url = url[:len(url)-1]

    // Fetch the URL
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error fetching URL: ", err)
        return
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body: ", err)
        return
    }

    // Convert the body to a string (this contains the HTML)
    htmlContent := string(body)
    fmt.Println(htmlContent)
}
