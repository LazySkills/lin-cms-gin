/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	v1 "lin-cms-gin/internal/validator/v1"
	"lin-cms-gin/internal/middleware/permission"
	"lin-cms-gin/internal/models"
	"lin-cms-gin/pkg/e"
	"lin-cms-gin/pkg/lin"
	"lin-cms-gin/pkg/setting"
	"lin-cms-gin/pkg/tools"
	"net/http"
)

type BookPermission struct {
	lin.Permission
}

func (p *BookPermission) AuthMapping()  {
	p.Mapping("GetBooksInfo","GET","获取图书信息","图书",0)
	p.Mapping("GetBooks","GET","获取图书列表","图书",1)
	p.Mapping("AddBook","POST","新增图书","图书",0)
	p.Mapping("EditBook","PUT","修改图书","图书",0)
	p.Mapping("DeleteBook","DELETE","删除图书","图书",0)
}

// @Summary 获取图书信息
// @Tags 图书
// @Produce  json
// @Param id query int true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /v1/book/:id [get]
func GetBooksInfo(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
	)

	if !permission.GroupRequired(c.Request.Method,"GetBooksInfo") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		c.Abort()
		return
	}

	id := com.StrTo(c.Param("id")).MustInt()

	data := models.ExistBookByID(id)
	if 	data.ID < 1 {
		appG.ResponseError(http.StatusBadRequest, e.ERROR_NOT_EXIST_BOOK,nil)
	}

	appG.Response(http.StatusOK,data)
}

// @Summary 获取图书列表
// @Tags 图书
// @Produce  json
// @Param title query string false "title"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /v1/book [get]
func GetBooks(c *gin.Context) {
	var appG = lin.Gin{C: c}

	if !permission.GroupRequired(appG.C.Request.Method,"GetBooks") {
		appG.ResponseError(http.StatusForbidden, e.AUTH_FAIL,nil)
		appG.C.Abort()
		return
	}

	title := appG.C.Query("title")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if title != "" {
		maps["title"] = title
	}

	var state int = -1
	if arg := appG.C.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data["lists"] = models.GetBook(tools.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetBookTotal(maps)

	appG.Response(http.StatusOK,data)
}


// @Summary 新增图书
// @Tags 图书
// @Produce  json
// @Param title query string true "图书名称"
// @Param author query string true "图书作者"
// @Param summary query string true "图书简介"
// @Param image query string true "图书图片"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /v1/books [post]
func AddBook(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		req = &v1.AddBookValidator{}
	)

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusBadRequest, e.INVALID_PARAMS,err)
		return
	}

	errCode := e.ERROR_EXIST_BOOK
	if ! models.ExistBookByName(req.Title) {
		errCode = e.SUCCESS
		models.AddBook(req.Title, req.Author, req.Summary, req.Image)
	}

	if errCode != 200 {
		appG.ResponseError(http.StatusBadRequest,errCode,nil)
	}else {
		appG.Response(http.StatusOK, nil)
	}
}


// @Summary 修改图书
// @Tags 图书
// @Produce  json
// @Param id path int true "ID"
// @Param title query string true "图书名称"
// @Param author query string true "图书作者"
// @Param summary query string true "图书简介"
// @Param image query string false "图书图片"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /v1/books/{id} [put]
func EditBook(c *gin.Context) {
	var (
		appG = lin.Gin{C: c}
		errCode = e.SUCCESS
		req = &v1.UpdateBookValidator{}
	)

	if err := lin.Validator(appG.C,req); err != ""{
		appG.ResponseError(http.StatusOK, e.SUCCESS,err)
		return
	}

	if models.ExistBookByID(req.ID).ID > 0 {
		data := make(map[string]interface{})
		data["title"] = req.Title
		data["author"] = req.Author
		data["summary"] = req.Summary
		data["image"] = req.Image
		models.EditBook(req.ID, data)
	} else {
		errCode = e.ERROR_NOT_EXIST_BOOK
	}

	if errCode != 200 {
		appG.ResponseError(http.StatusBadRequest,errCode,nil)
	}else {
		appG.Response(http.StatusOK, nil)
	}
}

// @Summary 删除图书
// @Tags 图书
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /v1/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	var appG = lin.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistBookByID(id).ID > 0 {
			models.DeleteBook(id)
		} else {
			code = e.ERROR_NOT_EXIST_BOOK
		}
	}

	if code != 200 {
		appG.ResponseError(http.StatusBadRequest,code,nil)
	}else {
		appG.ResponseSuccess(http.StatusOK,code, nil)
	}
}