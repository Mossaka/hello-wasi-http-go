package main

import (http "github.com/Mossaka/hello-wasi-http-go/target_world/2023_11_10")

func init() {
	server := HttpServer{}
	http.SetExportsWasiHttp0_2_0_rc_2023_11_10_IncomingHandler(server)
}

type HttpServer struct {}

// Helper type aliases to make generated code more readable
type HttpRequest = http.ExportsWasiHttp0_2_0_rc_2023_11_10_IncomingHandlerIncomingRequest
type HttpResponseWriter = http.ExportsWasiHttp0_2_0_rc_2023_11_10_IncomingHandlerResponseOutparam
type HttpOutgoingResponse = http.WasiHttp0_2_0_rc_2023_11_10_TypesOutgoingResponse
type HttpErrorCode = http.WasiHttp0_2_0_rc_2023_11_10_TypesErrorCode
type HttpTrailers = http.WasiHttp0_2_0_rc_2023_11_10_TypesTrailers


func (h HttpServer) Handle(request HttpRequest, response_out HttpResponseWriter) {
	// Create HTTP response
	headers := http.NewFields()
	response := http.NewOutgoingResponse(headers)
	response.SetStatusCode(200)
	body := response.Body().Unwrap()
	res_result := http.Ok[HttpOutgoingResponse, HttpErrorCode](response)
	http.StaticResponseOutparamSet(response_out, res_result)

	out := body.Write().Unwrap()
	out.BlockingWriteAndFlush([]uint8("Hello world from Go!!!\n")).Unwrap()
	out.Drop()

	http.StaticOutgoingBodyFinish(body, http.None[HttpTrailers]())
	
}

//go:generate wit-bindgen tiny-go wit/2023_11_10 --out-dir=target_world/2023_11_10 --gofmt
func main() {}