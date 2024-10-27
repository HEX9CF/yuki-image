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
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Insert Failure", Data: gin.H{"error": err}})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "插入成功", Data: gin.H{"album": album}})
}

func UpdateAlbum(ctx *gin.Context) {
	var album model.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	err = ialbum.Update(album, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "更新失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "更新成功", Data: nil})
}

func SelectAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	album, err := ialbum.Select(uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"album": album}})
}

func SelectAllAlbum(ctx *gin.Context) {
	albums, err := ialbum.SelectAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"album": albums}})
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
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Insert Failure", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "插入成功", Data: nil})
}

func SelectFormatSupport(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	format_support, err := ialbum.SelectFormatSupport(uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"format_support": format_support}})
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
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "删除失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "删除成功", Data: nil})
}

func SelectImageFromAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	upage := uint64(page)
	size, err := strconv.Atoi(ctx.Query("size"))
	if err != nil {
		size = conf.Conf.Server.ImageListDefalutSize
	}
	usize := uint64(size)

	imageList, err := ialbum.SelectImage(uint64(id), upage, usize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"image_list": imageList}})
}