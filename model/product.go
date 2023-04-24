package model

import (
	"gorm.io/gorm"
) 


type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Album struct {
	gorm.Model
	Name   string `json:"name,omitempty" form:"name" gorm:"not null;size:255"`
	Detail string `json:"detail,omitempty" form:"detail" gorm:"type:text"`
	ImageCover  string `json:"imageCover,omitempty" form:"imageCover" gorm:"type:longtext"`
	Images []AlbumExtraImage
}

type AlbumExtraImage struct {
	gorm.Model
	AlbumID uint `json:"albumId,omitempty" form:"albumId" gorm:"type:int(10) unsigned;not null"`
	Image   string `json:"image,omitempty" form:"image" gorm:"type:longtext"`
}
