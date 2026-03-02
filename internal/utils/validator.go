package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type Utils struct {
	Validate *validator.Validate
}

func NewUtils(v *validator.Validate) *Utils {
	return &Utils{
		Validate: v,
	}
}
func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func (u *Utils) BindAndValidate(r *http.Request, target any) error {
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(target); err != nil {
			return err
		}
	} else {
		if err := r.ParseForm(); err != nil {
			return err
		}
		if err := schema.NewDecoder().Decode(target, r.Form); err != nil {
			return err
		}
	}
	return u.Validate.Struct(target)
}
