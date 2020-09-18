/** Created By wene<354007048@qq.com> . Date at 2020/7/5 */
package admin

import (
	"github.com/gin-gonic/gin"
	"lin-cms-gin/internal/pkg/lin"
	"lin-cms-gin/internal/pkg/upload"
)

type LinFilePermission struct {
	lin.Permission
}

func (p *LinFilePermission) AuthMapping()  {
	p.Mapping("UploadImage","POST","上传文件","文件管理",0)
}

// @Summary 上传文件
// @Tags 文件管理
// @Produce  json
// @Accept multipart/form-data
// @Param file formData file true "文件资源,可以写多个"
// @Param file1 formData file false "文件资源1,可以写多个"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload [post]
func UploadImage(c *gin.Context) {
	upload.UploadFile(c)
}