package scan_files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ScanCurrentFolder() {
	err1 := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fmt.Println(path)
		}

		return nil
	})
	_, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	if err1 != nil {
		log.Fatal(err1)
	}

	/* for _, f := range files {
		if f.IsDir() {

		}
		fmt.Println(f.Name())

	} */
}
