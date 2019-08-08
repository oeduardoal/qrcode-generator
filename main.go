package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gorilla/mux"
	"image/png"
	"net/http"
	"os"
)

func GenerateQrCode(w http.ResponseWriter, r *http.Request){
	dataString := r.FormValue("q")
	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	png.Encode(w, qrCode)


}

func getPort() string {
	p := os.Getenv("PORT")
	if p != ""{
		return ":" + p
	}else {
		return ":8080"
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/generate", GenerateQrCode).Methods("GET")
	p := getPort()
	http.ListenAndServe(p, r)
	
}
