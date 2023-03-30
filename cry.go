package main

import (
	"bufio"
    "fmt"
    "crypto/aes"
    "crypto/cipher"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    fmt.Print("Enter filename: ")
    reader := bufio.NewReader(os.Stdin)
    filename, _ := reader.ReadString('\n') 
    filename = strings.TrimSuffix(filename, "\n")    
    fmt.Println(filename)

    fmt.Print("e=encrypt or d=decrypt:")
    reader = bufio.NewReader(os.Stdin)
    mode,_:= reader.ReadString('\n') 
    mode = strings.TrimSuffix(mode, "\n")    
    fmt.Println(mode)


    fmt.Print("Enter Password (32 numbers): \n")
    reader = bufio.NewReader(os.Stdin)
    password,_:= reader.ReadString('\n') 
    password = strings.TrimSuffix(password, "\n")    
    fmt.Println(password)

    fmt.Print("Enter Vectors (16 numbers): \n")
    reader = bufio.NewReader(os.Stdin)
    vectors,_:= reader.ReadString('\n') 
    vectors = strings.TrimSuffix(vectors, "\n")    
    fmt.Println(vectors)


for i := 0; len(vectors) < 16; i++ {
    vectors += string(vectors[i%len(vectors)])
}
for i := 0; len(password) < 32; i++ {
    password += string(password[i%len(password)])
}
    fmt.Print("Filename: ",filename," Mode: ",mode)
    fmt.Println()
    fmt.Print(password," Length: ", len(password))
    fmt.Println()
    fmt.Print(vectors," Length: ", len(vectors))
    fmt.Println()
    key := []byte(password)
    iv := []byte(vectors)

    if mode[0]=='e' {

    fmt.Println("Encrypting file...")
        err := encryptFile(filename, key, iv)
    if err != nil {
        panic(err)
    }
    fmt.Println("File encrypted successfully.")

    }else{

    fmt.Println("Decrypting file...")
        err := decryptFile(filename, key, iv)
    if err != nil {
        panic(err)
    }
    fmt.Println("File decrypted successfully.")
    }
}

func encryptFile(filename string, key []byte, iv []byte) error {
    plaintext, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return err
    }

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

    f, err := os.Create(filename + ".enc")
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Write(ciphertext)
    if err != nil {
        return err
    }

    return nil
}

func decryptFile(filename string, key []byte, iv []byte) error {
    ciphertext, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return err
    }

    plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])

    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Write(plaintext)
    if err != nil {
        return err
    }

    return nil
}
