package handler

import (
	"backend/request"
	"backend/service"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type institusiHandler struct {
	institusiService service.InstitusiService
}

func NewInstitusiHandler(institusiService service.InstitusiService) *institusiHandler {
	return &institusiHandler{institusiService}
}

func (institusiHandler *institusiHandler) ListInstitusi(c *gin.Context) {
	institusis, err := institusiHandler.institusiService.ListInstitusi()
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", institusis)
}

//nyoba generic json parser
func (institusiHandler *institusiHandler) AddInstitusi(c *gin.Context) {
	//create object of InstitusiRequest
	var institusi request.InstitusiRequest

	//lempar request ke parser untuk mengubah json jadi struct
	err := util.ParseData(c, &institusi)
	if err == nil {
		//jika tidak ada masalah, lanjut insert data
		err := institusiHandler.institusiService.InsertInstitusi(institusi)
		if err != nil {
			util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
			return
		}
		util.APIResponse(c, "Berhasil menambah Institusi", http.StatusOK, "ok", &institusi)
	}
}
