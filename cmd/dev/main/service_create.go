package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func replaceTextContent(content string, serviceName string ) string{
	content = strings.ReplaceAll(string(content), "template", serviceName)
	content = strings.ReplaceAll(content, "Template", strings.Title(serviceName))

    return content;
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Необходимо указать имя сервиса: go run struct.go <serviceName>")
		return
	}

	serviceName := os.Args[1]

	projectRoot := "../../../"
	domainsPath := filepath.Join(projectRoot, "domains")
	templateService := filepath.Join(domainsPath, "template_service")
	userService := filepath.Join(domainsPath, serviceName+"_service")



    if stat, err := os.Stat(userService); err == nil && stat.IsDir() {
     fmt.Println("Сервис", serviceName, "уже существует!");
      return;
    }


	err := copyDir(templateService, userService)
	if err != nil {
		fmt.Println("Ошибка при копировании папки:", err)
		return
	}

	err = filepath.Walk(userService, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			newContent := replaceTextContent(string(content), serviceName);

			pathPrev :=  path;
			path =  replaceTextContent(path, serviceName);

            if(path != pathPrev){
               os.Remove(pathPrev)
            }

			err = ioutil.WriteFile(path, []byte(newContent), 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Ошибка при замене текста:", err)
		return
	}

	fmt.Println("Сервис", serviceName, "успешно создан!")
}

func copyDir(src, dst string) error {
	err := os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			data, err := ioutil.ReadFile(srcPath)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(dstPath, data, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
