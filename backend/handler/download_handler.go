package handler

import (
	"net/http"
	"backend/entity"
	"github.com/gin-gonic/gin"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
	"time"
	"strings"
	"strconv"
)

type downloadHandler struct {

}

func NewDownloadHandler() *downloadHandler {
	return &downloadHandler{}
}

func (downloadHandler *downloadHandler) DynamicHtmlTest(c *gin.Context) {
	userData := []entity.User{
        {Username: "adam", Password: "099911111", Email: "tes@tes.com"},
        {Username: "brown", Password: "0933222222", Email: "tes@tes.com"},
        {Username: "cayla", Password: "0933333333", Email: "tes@tes.com"},
    }
	c.HTML(http.StatusOK, "table.tmpl", gin.H{
		"userdata": userData,
	})
}

func (downloadHandler *downloadHandler) GeneratePdf(c *gin.Context) {
	
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println("Error generate pdf", err)
	}

	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	var url strings.Builder
	url.WriteString(os.Getenv("URL_APLIKASI"))
	url.WriteString("/api/html-test")
	page1 := wkhtmltopdf.NewPage(url.String())
	pdfg.AddPage(page1)

	pdfg.Dpi.Set(300)
	err = pdfg.Create()
	if err != nil {
		log.Println("Error generate pdf", err)
	}
	var filename strings.Builder
	t := time.Now()
	filename.WriteString("./pdf/output-")
	s := []string{strconv.Itoa(t.Year()), t.Month().String(), strconv.Itoa(t.Day()),
		strconv.Itoa(t.Hour()),strconv.Itoa(t.Minute()), strconv.Itoa(t.Second())}
	filename.WriteString(strings.Join(s, "-"))
	filename.WriteString(".pdf")
	err = pdfg.WriteFile(filename.String())
	if err != nil {
		log.Println("Error generate pdf", err, filename.String())
	}

	c.FileAttachment(filename.String(), "output-"+strings.Join(s, "-")+".pdf")

	log.Println("Done")
}