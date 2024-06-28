package gogui

import (
	"fmt"
	"os"
	"path/filepath"
)

func bundleMeta() string {
	return "<title>gogui</title>"
}

func bundleJs() string {
	data := "<script type=\"module\" defer>%s</script>"

	dirPath := filepath.Join("gogui", "js")
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		println(err.Error())
		return ""
	}

	content := ""
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, entry.Name())
		f, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		content += string(f) + "\n"
	}

	return fmt.Sprintf(data, content)
}

func bundleCss() string {
	data := "<style>%s</style>"

	dirPath := filepath.Join("gogui", "css")
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		println(err.Error())
		return ""
	}

	content := ""
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, entry.Name())
		f, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		content += string(f) + "\n"
	}

	return fmt.Sprintf(data, content)
}

func bundle() string {
	html := "<html><head>"

	html += bundleMeta()
	html += bundleJs()
	html += bundleCss()

	html += "</head><body>"

	return html
}
