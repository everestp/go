package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)


func main(){
	http.HandleFunc("/",  func(w http.ResponseWriter, r *http.Request){
		logRequestDetails(r)
		fmt.Fprintln(w,"Hello server")
	})
		http.HandleFunc("/orders",  func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w,"Handling orders")
	})
	const port string = "3000"
	//Load the tls certificate
	cert :="cert.pem"
	key :="key.pem"

	//Config the tls
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//Create a customer server
	server := &http.Server{
		Addr: fmt.Sprintf(":%v",port),
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	//Enable http2
	http2.ConfigureServer(server,&http2.Server{})
	fmt.Println("Server is listening om port :",port)
    err := server.ListenAndServeTLS(cert, key)
   if err != nil{
	 	log.Fatalln("error starting the server",err)
	 }
// This is the HTTP 1.1 Server with TLS

	// err :=http.ListenAndServe(port, nil)
	// if err != nil{
	// 	log.Fatalln("error starting the server",err)
	// }
	
}

func logRequestDetails(r *http.Request){
	httpVersion :=r.Proto
	fmt.Println("Recieved Request with HTTP Version",httpVersion)
	if r.TLS != nil{
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received Request With TLS version",tlsVersion)
	} else{
		fmt.Println("Received Request with out TLS")
	}
}
func getTLSVersionName(version uint16) string{
	 switch version{
	 case tls.VersionTLS10:
		return "TLS 1.0"
		case tls.VersionTLS11:
		return "TLS 1.1"
		case tls.VersionTLS12:
		return "TLS 1.2"
		case tls.VersionTLS13:
		return "TLS 1.3"
		default:
			return "Unknown TLS Version"
	 }
}
