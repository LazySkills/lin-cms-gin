/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

import jtime "lin-cms-gin/internal/pkg/time"

// LinFile [...]
type LinFile struct {
	Model
	Path       string    `gorm:"column:path;type:varchar(500);not null;comment:'路径'"`
	Type       string    `gorm:"column:type;type:varchar(10);not null;default:'LOCAL';comment:'LOCAL 本地，REMOTE 远程'"`
	Name       string    `gorm:"column:name;type:varchar(100);not null;comment:'名称'"`
	Extension  string    `gorm:"column:extension;type:varchar(50);not null;comment:'后缀'"`
	Size       int       `gorm:"column:size;type:int(11);not null;comment:'大小'"`
	Md5        string    `gorm:"unique_index:md5_del;column:md5;type:varchar(40);not null;comment:'图片md5值，防止上传重复图片'"`
	DeleteTime jtime.JSONTime `gorm:"unique_index:md5_del;column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
}

func AddLinFile(name string, path string, fileType string, ext string, size int, md5 string) int {
	var linFile = LinFile{
		Name: name,
		Path: path,
		Type: fileType,
		Extension: ext,
		Size: size,
		Md5: md5,
	}
	db.Create(&linFile)
	return linFile.ID
}

func GetLinFileByFileMd5(md5 string) (file LinFile) {
	db.Where("md5 = ?", md5).Find(&file)
	return file
}

