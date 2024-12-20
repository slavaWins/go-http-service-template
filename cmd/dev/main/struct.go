package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {




    projectRoot :=  "../../../"

	err := printDirStructure(projectRoot)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}

func printDirStructure(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}


		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

      if strings.Contains(relPath, ".idea") {
        return nil;
      }


      if strings.Contains(relPath, ".git") {
         			return nil
        }

		if info.IsDir() {
			fmt.Println("├───" + relPath + "/")
		} else {
			fmt.Println("│   ├───" + relPath)
		}
		return nil
	})
	return err
}
