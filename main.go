package main

import (
	"bufio"
	//"cry/lib/scan_files"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const encryptedMarker = "antiGITHUB.txt"

var exeName string

func main() {
	//scan_files.ScanCurrentFolder()
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeName = filepath.Base(exePath)
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// handle err
	fmt.Println("This script is executed in this directory:")
	fmt.Println(path)

	folderPath := "."
	fmt.Println("Checking if folder is already encrypted...")
	xd, _ := os.Stat(folderPath + "/" + encryptedMarker)
	var isGitAllowed bool
	if xd == nil {
		fmt.Println("Encrypt Mode...")
		isGitAllowed = false //means is encripted, so its already safe to git commit/push
	} else {
		fmt.Println("SAFE\n->Decrypt Mode...")
		isGitAllowed = true
	}
	fmt.Print("Enter Password alphanumeric (32 numbers max): \n")
	reader := bufio.NewReader(os.Stdin)
	password, _ := reader.ReadString('\n')
	password = strings.TrimSuffix(password, "\n")

	fmt.Print("Enter Vectors Numbers 0 to 9 (16 numbers max): \n")
	reader = bufio.NewReader(os.Stdin)
	vectors, _ := reader.ReadString('\n')
	vectors = strings.TrimSuffix(vectors, "\n")

	for i := 0; len(vectors) < 16; i++ {
		vectors += string(vectors[i%len(vectors)])
	}
	for i := 0; len(password) < 32; i++ {
		password += string(password[i%len(password)])
	}

	key := []byte(password)
	iv := []byte(vectors)
	enc_dec_path(folderPath, key, iv, isGitAllowed)
	fmt.Println("finish")
	//if git is allowed, the folder is encrypted
	//otherwise the folder is decypted and git should be disallowed
	if !isGitAllowed {
		err := ioutil.WriteFile(folderPath+"/"+encryptedMarker, []byte("This file is used to check if the curren folder is encrypted, dont delete."), 0644)
		if err != nil {

			panic("Error creating ")
		}
	} else {
		err := os.Remove(folderPath + "/" + encryptedMarker)
		if err != nil {
			panic("Error creating ")
		}
	}

}

func enc_dec_file(filename string, key []byte, iv []byte, isGitAllowed bool) error {
	initial_file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("E1")
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("E2")
		return err
	}
	var final_file []byte
	if isGitAllowed {
		//decrypting...
		final_file = make([]byte, len(initial_file)-aes.BlockSize)
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(final_file, initial_file[aes.BlockSize:])
	} else {
		final_file = make([]byte, aes.BlockSize+len(initial_file))
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(final_file[aes.BlockSize:], initial_file)
	}

	f, err := os.Create(filename)
	if err != nil {

		fmt.Println("Error in 96", filename)

		return err
	}
	defer f.Close()

	_, err = f.Write(final_file)
	if err != nil {
		fmt.Println("error in 104")

		return err
	}

	return nil
}
func isGitAllowed(folderPath string) bool {
	_, err := os.Stat(folderPath + "/" + encryptedMarker)
	return err == nil
}
func enc_dec_path(folderPath string, key []byte, iv []byte, isGitAllowed bool) error {

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("error in !dasdasdasdasr")

		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			if file.Name() == "README.md" || file.Name() == exeName {
				continue
			}
			err := enc_dec_file(folderPath+"/"+file.Name(), key, iv, isGitAllowed)
			if err != nil {
				fmt.Println("error in !isDir")

				return err
			}
		} else {
			if file.Name() != "node_modules" && file.Name() != "public" && file.Name() != ".git" && file.Name() != "enc.git" {
				err := enc_dec_path(folderPath+"/"+file.Name(), key, iv, isGitAllowed)
				if err != nil {
					fmt.Println("error in else of isDirasd")

					return err
				}
			}
			if file.Name() == ".git" && isGitAllowed {
				err := os.Rename(folderPath+"/.git", folderPath+"/enc.git")
				if err != nil {
					fmt.Println("error in renaming .git to enc.git")
					return err
				}
			}
			if file.Name() == "enc.git" && !isGitAllowed {
				err := os.Rename(folderPath+"/enc.git", folderPath+"/.git")
				if err != nil {
					fmt.Println("error in renaming env.git to .git")
					return err
				}
			}
		}
	}

	return nil
}
