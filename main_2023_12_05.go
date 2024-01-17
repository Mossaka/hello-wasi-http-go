package main

import (
	http "github.com/Mossaka/hello-wasi-http-go/target_world/2023_12_05"
)

// Helper type aliases to make generated code more readable
type HttpRequest = http.ExportsWasiHttp0_2_0_rc_2023_12_05_IncomingHandlerIncomingRequest
type HttpResponseWriter = http.ExportsWasiHttp0_2_0_rc_2023_12_05_IncomingHandlerResponseOutparam
type HttpOutgoingResponse = http.WasiHttp0_2_0_rc_2023_12_05_TypesOutgoingResponse
type HttpError = http.WasiHttp0_2_0_rc_2023_12_05_TypesErrorCode
type HttpTrailers = http.WasiHttp0_2_0_rc_2023_12_05_TypesTrailers

type HttpServer struct{}

func init() {
	httpserver := HttpServer{}
	// Set the incoming handler struct to HttpServer
	http.SetExportsWasiHttp0_2_0_rc_2023_12_05_IncomingHandler(httpserver)
}

func (h HttpServer) Handle(request HttpRequest, responseWriter HttpResponseWriter) {
	// Construct HttpResponse to send back
	headers := http.NewFields()
	
	httpResponse := http.NewOutgoingResponse(headers)
	httpResponse.SetStatusCode(200)

	body := httpResponse.Body().Unwrap()
	okResponse := http.Ok[HttpOutgoingResponse, HttpError](httpResponse)
	http.StaticResponseOutparamSet(responseWriter, okResponse)

	stream := body.Write().Unwrap()
	stream.BlockingWriteAndFlush([]uint8("Hello from Go!\n")).Unwrap()
	stream.Drop()

	http.StaticOutgoingBodyFinish(body, http.None[HttpTrailers]())
	
}

//go:generate wit-bindgen tiny-go wit/2023_12_05 --out-dir=target_world/2023_12_05 --gofmt
func main() {}
