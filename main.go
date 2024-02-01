package main

import (http "github.com/Mossaka/hello-wasi-http-go/target_world")

func init() {
	server := HttpServer{}
	http.SetExportsWasiHttp0_2_0_IncomingHandler(server)
}

type HttpServer struct {}

// Helper type aliases to make generated code more readable
type HttpRequest = http.ExportsWasiHttp0_2_0_IncomingHandlerIncomingRequest
type HttpResponseWriter = http.ExportsWasiHttp0_2_0_IncomingHandlerResponseOutparam
type HttpOutgoingResponse = http.WasiHttp0_2_0_TypesOutgoingResponse
type HttpErrorCode = http.WasiHttp0_2_0_TypesErrorCode
type HttpTrailers = http.WasiHttp0_2_0_TypesTrailers
type HttpError = http.WasiHttp0_2_0_TypesErrorCode


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

//go:generate wit-bindgen tiny-go wit --out-dir=target_world --gofmt
func main() {}