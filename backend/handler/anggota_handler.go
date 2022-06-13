package handler

import (
	"backend/request"
	"backend/service"
	"backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type anggotaHandler struct {
	anggotaService service.AnggotaService
}

func NewAnggotaHandler(anggotaService service.AnggotaService) *anggotaHandler {
	return &anggotaHandler{anggotaService}
}

func (anggotaHandler *anggotaHandler) FindAnggotaById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	anggota, err := anggotaHandler.anggotaService.FindAnggotaByID(uint(id))
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", anggota)
}

func (anggotaHandler *anggotaHandler) ListAnggota(c *gin.Context) {
	anggotas, err := anggotaHandler.anggotaService.ListAnggota()
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", anggotas)
}

func (anggotaHandler *anggotaHandler) InsertAnggota(c *gin.Context) {
	var anggota request.AnggotaRequest
	if err := c.ShouldBindJSON(&anggota); err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	err := anggotaHandler.anggotaService.InsertAnggota(anggota)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Insert data success!", http.StatusOK, "ok", nil)
}

//nyoba generic json parser
func (anggotaHandler *anggotaHandler) AddAnggota(c *gin.Context) {
	//create object of AnggotaRequest
	var anggota request.AnggotaRequest

	//lempar request ke parser untuk mengubah json jadi struct
	err := util.ParseData(c, &anggota)
	if err == nil {
		//jika tidak ada masalah, lanjut insert data
		err := anggotaHandler.anggotaService.InsertAnggota(anggota)
		if err != nil {
			util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
			return
		}
		util.APIResponse(c, "Berhasil menambah Anggota", http.StatusOK, "ok", &anggota)
	}
}

func (anggotaHandler *anggotaHandler) UpdateAnggota(c *gin.Context) {
	var anggota request.AnggotaRequest
	if err := c.ShouldBindJSON(&anggota); err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	err := anggotaHandler.anggotaService.UpdateAnggota(anggota)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Update data success!", http.StatusOK, "ok", nil)
}

func (anggotaHandler *anggotaHandler) DeleteAnggotaById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	err := anggotaHandler.anggotaService.DeleteAnggotaById(uint(id))
	if err == 0 {
		util.APIResponse(c, "Failed to delete data", http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Delete data success!", http.StatusOK, "ok", nil)
}
