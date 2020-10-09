/** Created By wene<354007048@qq.com> . Date at 2020/7/3 */
package lin

import (
	"fmt"
	"lin-cms-gin/internal/models"
	"lin-cms-gin/pkg/logging"
)

var PermissionMapping = make(map[string]permissionMapping)

type Permission struct {
}

type PermissionMate interface {
	Mapping(action string, method string, name string, module string, mount int)
	AuthMapping()
}


type permissionMapping struct {
	ID int // 权限ID
	Action string // 方法名称
	Method string // 请求方法
	Name string // 权限名称，如：分类列表
	Module string // 权限分组，如：分类
	Mount int  // 是否挂载权限，0 不挂载，1挂载
}


// lin cms 输出权限模型
type LinPermission struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Module string `json:"module"`
}

func (p Permission) Mapping(action string, method string, name string, module string, mount int) {
	logging.Info("*Lin Permission*:",fmt.Sprintf("[%s]模块：%s | 权限：%s | 是否挂载：%v",action, name, module, mount))

	var pMate = permissionMapping{
		ID: 0,
		Action: action,
		Method: method,
		Name: name,
		Module: module,
		Mount: mount,
	}

	linPermission := models.GetLinPermissionByName(pMate.Name)
	if linPermission.ID > 0 {
		if linPermission.Name != linPermission.Name || linPermission.Module != linPermission.Module || linPermission.Mount != linPermission.Mount {
			linPermission.UpdatePermission(pMate.Name,pMate.Module,pMate.Mount)
		}
	} else {
		linPermission = models.AddLinPermission(pMate.Name, pMate.Module, pMate.Mount)
	}
	pMate.ID = linPermission.ID
	PermissionMapping[GetPermissionEndpoint(action,method)] = pMate
}

func (p Permission) AuthMapping() {

}

func Include(v interface{})  {
	if per,ok := v.(PermissionMate);ok {
		per.AuthMapping()
	}
}


func GetPermissionMappingByName(method string,action string) (per permissionMapping) {
	fmt.Printf("method - action: %v \n", GetPermissionEndpoint(action,method))
	fmt.Printf("count: %v \n",len(PermissionMapping))
	for item,_ := range PermissionMapping {
		fmt.Printf("PermissionMapping item k: %v \n",item)
	}
	if v,ok := PermissionMapping[GetPermissionEndpoint(action,method)]; ok{
		return v
	}
	return
}

func GetPermissionEndpoint(action string, method string) (endpoint string) {
	return method+"-"+action
}

func FormatLinPermission(userPermissions []models.LinPermission) (linPermission map[string][]LinPermission) {
	linPermission = make(map[string][]LinPermission)
	for _,v := range userPermissions {
		permission := LinPermission{
			Id: v.ID,
			Name: v.Name,
			Module: v.Module,
		}

		linPermission[v.Module] = append(linPermission[v.Module],permission)
	}
	return linPermission
}

func FormatLinGroupPermission(userPermissions []models.JsonLinGroupPermission) (linPermission map[string][]LinPermission) {
	linPermission = make(map[string][]LinPermission)
	for _,v := range userPermissions {
		permission := LinPermission{
			Id: v.ID,
			Name: v.Name,
			Module: v.Module,
		}

		linPermission[v.Module] = append(linPermission[v.Module],permission)
	}
	return linPermission
}