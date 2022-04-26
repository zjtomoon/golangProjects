package dao

import "time"

type FileTable struct {
	UserId         int    `gorm:"primaryKey column:UserId"`
	FilePath       string `gorm:"column:FilePath"`
	FileSM3        string `gorm:"column:FileSM3"`
	LastUpdateTime string `gorm:"column:LastUpdateTime"`
	isUpdated      bool   `gorm:"column:isUpdated"`
}

func NewFile(userid int, filepath string, filesm3 string, lastupdateTime time.Time, isupdated bool) *FileTable {
	return &FileTable{
		UserId:         userid,
		FilePath:       filepath,
		FileSM3:        filesm3,
		LastUpdateTime: lastupdateTime.String(),
		isUpdated:      isupdated,
	}
}
