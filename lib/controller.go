package lib

import (
	"encoding/json"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing"
)

type Controller struct {
	Response http.ResponseWriter
}

func (ctrl *Controller) Respond(context *routing.Context) *Controller {
	ctrl.Response = context.Response

	return ctrl
}

func (ctrl *Controller) StatusCode(statusCode int) *Controller {
	ctrl.Response.WriteHeader(statusCode)

	return ctrl
}

func (ctrl *Controller) Json(data interface{}) error {
	content, _ := json.Marshal(data)
	_, err := ctrl.Response.Write(content)

	return err
}
