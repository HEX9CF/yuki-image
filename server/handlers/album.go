package handlers

import (
	"net/http"
	"strconv"
	ialbum "yuki-image/internal/album"
	"yuki-image/internal/conf"
	"yuki-image/internal/model"

	"github.com/gin-gonic/gin"
)

func InsertAlbum(ctx *gin.Context) {
	var album model.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	id, err := ialbum.Insert(album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Insert Failure"})
		return
	}
	album, err = ialbum.Select(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Insert Failure", Data: err})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "插入成功", Data: album})
}

func UpdateAlbum(ctx *gin.Context) {
	var album model.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}

	err = ialbum.Update(album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "更新失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "更新成功", Data: nil})
}

func SelectAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	idHasValue := err == nil && id > 0
	name := ctx.Query("name")
	nameHasValue := name != ""
	if !idHasValue && !nameHasValue {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Query error"})
		return
	}
	var album model.Album
	if !idHasValue {
		album, err = ialbum.SelectFromName(name)
	} else {
		album, err = ialbum.Select(uint64(id))
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: album})
}

func SelectAllAlbum(ctx *gin.Context) {
	albums, err := ialbum.SelectAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: albums})
}

func InsertFormatSupport(ctx *gin.Context) {
	var formatSupport model.FormatSupport
	err := ctx.ShouldBindJSON(&formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	err = ialbum.InsertFormatSupport(formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Insert Failure", Data: err})
		return
	}
	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "插入成功", Data: nil})
}

func SelectFormatSupport(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	idHasValue := err == nil && id > 0
	name := ctx.Query("name")
	nameHasValue := name != ""
	if !idHasValue && !nameHasValue {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Query error"})
		return
	}
	var format_support []model.Format
	if !idHasValue {
		format_support, err = ialbum.SelectFormatSupportFromName(name)
	} else {
		format_support, err = ialbum.SelectFormatSupport(uint64(id))
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: format_support})
}

func SelectImageFromAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	idHasValue := err == nil && id > 0
	name := ctx.Query("name")
	nameHasValue := name != ""
	if !idHasValue && !nameHasValue {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Query error"})
		return
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Query error"})
		return
	}
	upage := uint64(page)
	size, err := strconv.Atoi(ctx.Query("size"))
	if err != nil {
		size = conf.Conf.Image.ImageListDefalutSize
	}
	usize := uint64(size)
	var imageList model.ImageList
	if !idHasValue {
		imageList, err = ialbum.SelectImageFromName(name, upage, usize)
	} else {
		imageList, err = ialbum.SelectImage(uint64(id), upage, usize)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: imageList})
}

func DeleteFormatSupport(ctx *gin.Context) {
	var formatSupport model.FormatSupport
	err := ctx.ShouldBindJSON(&formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	err = ialbum.DeleteFormatSupport(formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "删除失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "删除成功", Data: nil})
}
