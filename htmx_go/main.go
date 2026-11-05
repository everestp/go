package main

import (
	"context"
	"fmt"
	"time"
)

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation canceled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func main() {
	ctx := context.TODO()
	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with contxt", result)
	ctx = context.Background()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
}

//+===== DIFFERENCE BETWEEN toDO AND background()
// func main(){
// 	 http.HandleFunc("/",  func( w http.ResponseWriter, r *http.Request){
//           templ ,err := template.ParseFiles("index.html")
// 		  if err != nil{
// 			http.Error(w, err.Error() ,http.StatusInternalServerError )
// 		  }
// 		  data :=  struct {
// 			Message string
// 		  }{Message : "Hellor"}
// 		  templ.Execute(w, data)
// 	 })

// 	 http.HandleFunc("/time",  func( w http.ResponseWriter,r *http.Request){
// 		w.Write([]byte("Current Time "+ time.Now().Format(time.RFC1123)))
// 	 })
// 	 fmt.Println("Started http server at port :8080")
// 	 http.ListenAndServe(":8080", nil)
// }
