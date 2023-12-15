package request

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type ResponseStatus struct {
	IsError          bool
	Body             string
	StatusCode       int
	Status           string
	RedirectLocation string
	ErrorMessage     string
}

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func (r *RequestError) StatusMessage() string {
	switch statusCode := r.StatusCode; statusCode {
	case http.StatusContinue: //100
		return fmt.Sprint("Continue")
	case http.StatusSwitchingProtocols: //101
		return fmt.Sprint("Switching Protocols")
	case http.StatusProcessing: //102
		return fmt.Sprint("Processing")
	case http.StatusEarlyHints: //103
		return fmt.Sprint("Early Hints")

	case http.StatusOK: //200
		return fmt.Sprint("OK")
	case http.StatusCreated: //201
		return fmt.Sprint("Created")
	case http.StatusAccepted: //202
		return fmt.Sprint("Accepted")
	case http.StatusNonAuthoritativeInfo: //203
		return fmt.Sprint("Non-Authoritative Information")
	case http.StatusNoContent: //204
		return fmt.Sprint("No Content")
	case http.StatusResetContent: //205
		return fmt.Sprint("Reset Content")
	case http.StatusPartialContent: //206
		return fmt.Sprint("Partial Content")
	case http.StatusMultiStatus: //207
		return fmt.Sprint("Multi-Status")
	case http.StatusAlreadyReported: //208
		return fmt.Sprint("Already Reported")
	case http.StatusIMUsed: //226
		return fmt.Sprint("IM Used")

	case http.StatusMultipleChoices: //300
		return fmt.Sprint("Multiple Choices")
	case http.StatusMovedPermanently: //301
		return fmt.Sprint("Moved Permanently")
	case http.StatusFound: //302 - Found (Previously "Moved Temporarily")
		return fmt.Sprint("Found")
	case http.StatusSeeOther: //303
		return fmt.Sprint("See Other")
	case http.StatusNotModified: //304
		return fmt.Sprint("Not Modified")
	case http.StatusUseProxy: //305
		return fmt.Sprint("Use Proxy")
	//306 - Switch Proxy - Undefined
	case http.StatusTemporaryRedirect: //307
		return fmt.Sprint("Temporary Redirect")
	case http.StatusPermanentRedirect: //308
		return fmt.Sprint("Permanent Redirect")

	case http.StatusBadRequest: //400
		return fmt.Sprint("Bad Request")
	case http.StatusUnauthorized: //401
		return fmt.Sprint("Unauthorized")
	case http.StatusPaymentRequired: //402
		return fmt.Sprint("Payment Required")
	case http.StatusForbidden: //403
		return fmt.Sprint("Forbidden")
	case http.StatusNotFound: //404
		return fmt.Sprint("Not Found")
	case http.StatusMethodNotAllowed: //405
		return fmt.Sprint("Method Not Allowed")
	case http.StatusNotAcceptable: //406
		return fmt.Sprint("Not Acceptable")
	case http.StatusProxyAuthRequired: //407
		return fmt.Sprint("Proxy Authentication Required")
	case http.StatusRequestTimeout: //408
		return fmt.Sprint("Request Timeout")
	case http.StatusConflict: //409
		return fmt.Sprint("Conflict")
	case http.StatusGone: //410
		return fmt.Sprint("Forbidden")
	case http.StatusLengthRequired: //411
		return fmt.Sprint("Length Required")
	case http.StatusPreconditionFailed: //412
		return fmt.Sprint("Precondition Failed")
	case http.StatusRequestEntityTooLarge: //413
		return fmt.Sprint("Entity Too Large")
	case http.StatusRequestURITooLong: //414
		return fmt.Sprint("URI Too Long")
	case http.StatusUnsupportedMediaType: //415
		return fmt.Sprint("Unsupported Media Type")
	case http.StatusRequestedRangeNotSatisfiable: //416
		return fmt.Sprint("Range Not Satisfiable")
	case http.StatusExpectationFailed: //417
		return fmt.Sprint("Expectation Failed")
	case http.StatusTeapot: //418
		return fmt.Sprint("I'm a Teapot")
	case http.StatusMisdirectedRequest: //421
		return fmt.Sprint("Misdirected Request")
	case http.StatusUnprocessableEntity: //422
		return fmt.Sprint("Unprocessable Entity")
	case http.StatusLocked: //423
		return fmt.Sprint("Locked")
	case http.StatusFailedDependency: //424
		return fmt.Sprint("Failed Dependency")
	case http.StatusTooEarly: //425
		return fmt.Sprint("Too Early")
	case http.StatusUpgradeRequired: //426
		return fmt.Sprint("Upgrade Required")
	case http.StatusPreconditionRequired: //428
		return fmt.Sprint("Precondition Required")
	case http.StatusTooManyRequests: //429
		return fmt.Sprint("Too Many Requests")
	case http.StatusRequestHeaderFieldsTooLarge: //431
		return fmt.Sprint("Request Header Fields Too Large")
	case http.StatusUnavailableForLegalReasons: //451
		return fmt.Sprint("Unavailable For Legal Reasons")

	case http.StatusInternalServerError: //500
		return fmt.Sprint("Internal Server Error")
	case http.StatusNotImplemented: //501
		return fmt.Sprint("Not Implemented")
	case http.StatusBadGateway: //502
		return fmt.Sprint("Bad Gateway")
	case http.StatusServiceUnavailable: //503
		return fmt.Sprint("Service Unavailable")
	case http.StatusGatewayTimeout: //504
		return fmt.Sprint("Gateway Timeout")
	case http.StatusHTTPVersionNotSupported: //505
		return fmt.Sprint("HTTP Version Not Supported")
	case http.StatusVariantAlsoNegotiates: //506
		return fmt.Sprint("Variant Also Negotiates")
	case http.StatusInsufficientStorage: //507
		return fmt.Sprint("Insufficient Storage")
	case http.StatusLoopDetected: //508
		return fmt.Sprint("Loop Detected")
	case http.StatusNotExtended: //510
		return fmt.Sprint("Not Extended")
	case http.StatusNetworkAuthenticationRequired: //511
		return fmt.Sprint("Network Authentication Required")
	default:
		return fmt.Sprint("Network Error!")
	}
}

func CustomError(resp *resty.Response, err error) *ResponseStatus {

	if resp != nil {

		if err != nil {

			re, ok := err.(*RequestError)

			if ok {

				return &ResponseStatus{
					IsError:      resp.IsError(),
					Status:       resp.Status(),
					StatusCode:   resp.StatusCode(),
					ErrorMessage: re.StatusMessage(),
				}

			}

		}

		return &ResponseStatus{
			IsError:      resp.IsError(),
			Status:       resp.Status(),
			StatusCode:   resp.StatusCode(),
			ErrorMessage: "Network Error!",
		}

	}

	return &ResponseStatus{
		IsError:      true,
		Status:       "",
		StatusCode:   0,
		ErrorMessage: "Network Error!",
	}

}

func Send(url string) *ResponseStatus {

	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))

	resp, err := client.R().EnableTrace().Get(url)

	if err != nil {
		return CustomError(resp, err)
	}

	return &ResponseStatus{
		IsError:      resp.IsError(),
		Status:       resp.Status(),
		StatusCode:   resp.StatusCode(),
		ErrorMessage: "",
	}

	/*client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}*/

	/*redirectUrl := ""
	statusCode := resp.StatusCode
	newUrl, err := resp.Location()
	if err == nil && newUrl.String() != "" && strings.Compare(newUrl.String(), url) > 0 && validation.CompareDomainForUrl(url, newUrl.String()) {
		redirectUrl = newUrl.String()
		statusCode = resp.Request.Response.StatusCode
	}*/

}
