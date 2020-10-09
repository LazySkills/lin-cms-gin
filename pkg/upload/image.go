/** Created By wene<354007048@qq.com> . Date at 2020/6/20 */
package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/models"
	"lin-cms-gin/pkg/e"
	"lin-cms-gin/pkg/file"
	"lin-cms-gin/pkg/lin"
	"lin-cms-gin/pkg/logging"
	"lin-cms-gin/pkg/setting"
	"lin-cms-gin/pkg/tools"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type Image struct {
	Url string `json:"path"`
	Id int `json:"id"`
	Key string `json:"key"`
}

func UploadFile(c *gin.Context)  {
	var appG = lin.Gin{C: c}
	code := e.SUCCESS
	var images []Image

	CheckImageFormSize(c)

	// 获取表单
	form := appG.C.Request.MultipartForm
	for fileFormName, _ := range form.File {
		files := form.File[fileFormName]
		for fileItem, _ := range files {
			fileObj, err := files[fileItem].Open()
			defer fileObj.Close()
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			}
			imageName := GetImageName(files[fileItem].Filename)
			fileMd5 := tools.EncodeMD5(imageName)
			linFile := models.GetLinFileByFileMd5(fileMd5)
			if linFile.ID > 0 {
				images = append(images, Image{
					Id:  linFile.ID,
					Key: fileFormName,
					Url: GetImageFullUrl(linFile.Name),
				})
			}else {
				fullPath := GetImageFullPath()
				savePath := GetImagePath()
				src := fullPath + imageName
				if ! CheckImageExt(imageName) || ! CheckImageSize(fileObj) {
					code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
				} else {
					err := CheckImage(fullPath)
					if err != nil {
						logging.Warn(err)
						code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
					} else if err := c.SaveUploadedFile(files[fileItem], src); err != nil {
						logging.Warn(err)
						code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
					} else {
						size,_:= file.GetSize(fileObj)
						ext := file.GetExt(imageName)
						linFileId := models.AddLinFile(
							imageName,
							savePath + imageName,
							"LOCAL",
							ext,
							size,
							fileMd5,
						)

						images = append(images, Image{
							Id:  linFileId,
							Key: fileFormName,
							Url: GetImageFullUrl(imageName),
						})
					}
				}
			}
		}
	}
	if code == 200 {
		appG.Response(code,images)
	}else {
		appG.ResponseError(http.StatusBadRequest,code,images)
	}
}

func GetImageFullUrl(name string) string {
	return setting.FileSetting.FileDomain  + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = tools.EncodeMD5(fileName)

	return fileName + ext
}

func GetImagePath() string {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	return setting.FileSetting.FileSavePath + year + "/" + month + "/" + day + "/"
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.FileSetting.FileInclude {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageFormSize(c *gin.Context)  {
	var appG = lin.Gin{C: c}
	err := c.Request.ParseMultipartForm(setting.FileSetting.FileMultipartMaxSize)
	if err != nil {
		logging.Warn(err)
		appG.ResponseError(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT,nil)
	}
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.FileSetting.FileMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission PermissionMate denied src: %s", src)
	}

	return nil
}