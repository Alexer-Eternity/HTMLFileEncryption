package main
import (
 "crypto/aes"
 "crypto/cipher"
 "encoding/base64"
 "fmt"
"log"
"os"
//"./string"
//"./number"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"
func Encode(b []byte) string {
 return base64.StdEncoding.EncodeToString(b)
}


func Encrypt(text, MySecret string) (string, error) {
 block, err := aes.NewCipher([]byte(MySecret))
 if err != nil {
  return "", err
 }
 plainText := []byte(text)
 cfb := cipher.NewCFBEncrypter(block, bytes)
 cipherText := make([]byte, len(plainText))
 cfb.XORKeyStream(cipherText, plainText)
 return Encode(cipherText), nil
}


func Writetofile(content string) string{

	f, err := os.Create("data.txt")
	 if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(content)

    if err2 != nil {
        log.Fatal(err2)
    }
return "data.txt"



}
func main() {
    StringToEncrypt := "???"

    encText, err := Encrypt(StringToEncrypt, MySecret)
    if err != nil {
     fmt.Println("error encrypting your classified text: ", err)
    }
    fmt.Println(Writetofile(encText))

}

