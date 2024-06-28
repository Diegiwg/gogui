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

		content += string(f)
	}

	return fmt.Sprintf(data, content)
}

func bundleCss() string {
	return ""
}

func bundle() string {
	html := "<html><head>"

	html += bundleMeta()
	html += bundleJs()
	html += bundleCss()

	html += "</head><body>%s</body></html>"

	return html
}
