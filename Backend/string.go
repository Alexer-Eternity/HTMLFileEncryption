 package main
import(
  "github.com/thanhpk/randstr"
  "fmt"
)
func main() {
    MyString := randstr.String(20)
    fmt.Println(MyString)

}package main

import (
    "encoding/base64"
    "fmt"
)

func main() {

    StringToEncode := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    Encoding := base64.StdEncoding.EncodeToString([]byte(StringToEncode))
    fmt.Println(Encoding)                                        
}
