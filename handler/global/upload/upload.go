package upload

import (
	"alpha/handler"
	"alpha/repositories/util/dir"
	"fmt"

	"bytes"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
	"os"
	"path"
	"time"
)

func Upload(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(2000); err != nil {
		handler.SendBadResponse(c, err, nil)
		return
	}
	formData := c.Request.MultipartForm
	files := formData.File["file"]
	fmt.Println(files)
	fileNames := make([]string, len(files))
	for i, v := range files {
		file, err := v.Open()
		defer file.Close()
		if err != nil {
			handler.SendBadResponse(c, err, nil)
			return
		}
		var currentPath bytes.Buffer

		currentTime := time.Now().Format("2006/01")

		currentPath.WriteString("uploads/")
		currentPath.WriteString(currentTime)

		if _, err := dir.IsDir(currentPath.String(), true); err != nil {
			handler.SendBadResponse(c, err, nil)
			return
		}

		cUuid := uuid.NewV1()
		if err != nil {
			handler.SendBadResponse(c, err, nil)
			return
		}
		currentPath.WriteString("/")
		currentPath.WriteString(cUuid.String())
		currentPath.WriteString(path.Ext(v.Filename))
		name := currentPath.String()
		out, err := os.Create(name)
		defer out.Close()

		if err != nil {
			handler.SendBadResponse(c, err, nil)
			return
		}
		if _, err := io.Copy(out, file); err != nil {
			handler.SendBadResponse(c, err, nil)
			return
		}
		fileNames[i] = name
	}

	handler.SendResponse(c, nil, fileNames)
}
