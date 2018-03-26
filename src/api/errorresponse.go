package api

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"

	"hubimage/src/pkg/common"
)

func handleInternalError(response *restful.Response, err error) {
	log.Print(err)

	statusCode := http.StatusInternalServerError
	if err == common.ErrConflict {
		statusCode = http.StatusConflict
	}
	/*statusError, ok := err.(*errorsK8s.StatusError)
	if ok && statusError.Status().Code > 0 {
		statusCode = int(statusError.Status().Code)
	}*/

	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(statusCode, err.Error()+"\n")
}
