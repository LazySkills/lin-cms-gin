/** Created By 嗝嗝<354007048@qq.com>. Date 2020/10/13 */
package cms

import "lin-cms-gin/pkg/time"

// 日志验证器
type LogFindValidator struct {
	Name int `form:"name" json:"v" binding:"alphanum"`
	Start time.JSONTime `form:"start" json:"start" binding:""`
	End time.JSONTime `form:"end" json:"end" binding:""`
}