package repo

import "time"

type File struct {
	ID      int64     `gorm:"column:id;primarykey"`
	MD5     string    `gorm:"column:md5"`
	Name    string    `gorm:"column:name"`
	Path    string    `gorm:"column:path"`
	Key     string    `gorm:"column:key;index"`
	Updated time.Time `gorm:"column:updated"`
}
