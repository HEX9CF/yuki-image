package db

import (
	"time"
	dbModel "yuki-image/internal/db/model"
)

func InsertAlbum(album dbModel.Album) (uint64, error) {
	time := time.Now()
	album.CreateTime = time
	album.UpdateTime = time
	tx := db.Create(&album)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return album.Id, nil
}

func UpdateAlbum(album dbModel.Album) error {
	tx := db.Model(&album).Where("id = ?", album.Id).Updates(map[string]interface{}{
		"name":        album.Name,
		"max_height":  album.MaxHeight,
		"max_width":   album.MaxWidth,
		"update_time": time.Now().Format("2006-01-02 15:04:05"),
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SelectAlbum(id uint64) (dbModel.Album, error) {
	var album dbModel.Album
	tx := db.First(&album, "id=?", id)
	if tx.Error != nil {
		return dbModel.Album{}, tx.Error
	}
	return album, nil
}

func SelectAlbumIdFromName(name string) (uint64, error) {
	var albumId uint64
	tx := db.Model(dbModel.Album{}).Where("name = ?", name).Pluck("id", &albumId)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return albumId, nil
}

func SelectAllAlbum() ([]dbModel.Album, error) {
	var albums []dbModel.Album
	tx := db.Find(&albums)
	if tx.Error != nil {
		return []dbModel.Album{}, tx.Error
	}
	return albums, nil
}

func SelectImageFromAlbum(albumId uint64, page uint64, size uint64) ([]dbModel.Image, error) {
	var images []dbModel.Image

	tx := db.Offset(int((page-1)*size)).Limit(int(size)).Where("album_id = ?", albumId).Find(&images)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return images, nil
}

func GetAlbumImageTotal(albumId uint64) (uint64, error) {
	var count int64
	tx := db.Model(&dbModel.Image{}).Where("album_id = ?", albumId).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return uint64(count), nil
}

func DeleteAlbum(albumId uint64) error {
	tx := db.Delete(&dbModel.Album{}, "id=?", albumId)
	if tx.Error != nil {
		return tx.Error
	}
	return tx.Error
}

func ClearAlbum(albumId uint64) error {
	tx := db.Delete(&dbModel.Image{}, "album_id=?", albumId)
	if tx.Error != nil {
		return tx.Error
	}
	return tx.Error
}
