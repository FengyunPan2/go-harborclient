package filter

import (
	"fmt"
	"time"

	restful "github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

const (
	// RequestLogString is a template for request log message.
	RequestLogString = "[%s] Incoming %s %s %s request from %s: %s"

	// ResponseLogString is a template for response log message.
	ResponseLogString = "[%s] Outcoming response to %s with %d status code"
)

// LogRequestAndResponse Log Request And Response
func LogRequestAndResponse(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	log.Info(formatRequestLog(request))
	chain.ProcessFilter(request, response)
	log.Info(formatResponseLog(response, request))
}

// formatRequestLog formats request log string.
func formatRequestLog(request *restful.Request) string {
	uri := ""
	if request.Request.URL != nil {
		uri = request.Request.URL.RequestURI()
	}

	return fmt.Sprintf(RequestLogString, time.Now().Format(time.RFC3339), request.Request.Proto,
		request.Request.Method, uri, request.Request.RemoteAddr)
}

// formatResponseLog formats response log string.
func formatResponseLog(response *restful.Response, request *restful.Request) string {
	return fmt.Sprintf(ResponseLogString, time.Now().Format(time.RFC3339),
		request.Request.RemoteAddr, response.StatusCode())
}
