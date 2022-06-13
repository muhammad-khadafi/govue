package util

import (
	"backend/entity"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

/**
 * Created by muhammad.khadafi on 27/05/2022
 */

func APIResponse(c *gin.Context, message string, code int, status string, data interface{}) {
	newToken, _ := c.Get("newToken")
	if newToken == nil {
		c.JSON(code,
			gin.H{
				"message": message,
				"code":    code,
				"status":  status,
				"data":    data,
			})
	} else {
		c.JSON(code,
			gin.H{
				"message":      message,
				"code":         code,
				"status":       status,
				"data":         data,
				"token_extend": newToken,
			})
	}

}

func APIResponseLogout(c *gin.Context, message string, code int, status string, data interface{}) {
	c.JSON(code,
		gin.H{
			"message": message,
			"code":    code,
			"status":  status,
			"data":    data,
		})
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func GetErrors(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}

//source https://blog.depa.do/post/gin-validation-errors-handling
func ParseData(c *gin.Context, e entity.EntityInterface) error {

	if err := c.ShouldBind(&e); err != nil {
		//// ValidationErrors is a slice of FieldError's
		//type ValidationErrors []FieldError
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			log.Info().Err(verr).Msg("Terdapat kesalahan di input JSON: " + err.Error())
			//APIResponse(c, err.Error(), http.StatusBadRequest, "error", gin.H{"Terdapat kesalahan di input JSON: ": GetErrors(verr)})
			APIResponse(c, "Terdapat kesalahan di input JSON", http.StatusBadRequest, "error", GetErrors(verr))
		} else {
			// We now know that this error is not a validation error
			// probably a malformed JSON
			log.Info().Err(err).Msg("NOT a validation error. Unable to bind")
			APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		}
		return errors.New("JSON Parse Error")
	}
	// Data is OK
	return nil
}
