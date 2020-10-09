/** Created By wene<354007048@qq.com> . Date at 2020/6/2 */

package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"lin-cms-gin/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
