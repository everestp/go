package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled", ctx.Err())
			return
		default:
			fmt.Println("We are working")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	rootCtx := context.Background()
	// ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second) // time of the context start here  After this time duration the context will send a cancelation signals

	ctx, cancel := context.WithCancel(rootCtx) // time of the context start here  After this time duration the context will send a cancelation signals

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	ctx = context.WithValue(ctx, "requestID", "dfdfd")
	go doWork(ctx)
	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID", requestID)
	} else {
		fmt.Println("No Request id found")

	}
	logWithContext(ctx, "this is  test log mesage")

}

func logWithContext(ctx context.Context, message string) {
	requestIdval := ctx.Value("requestID")
	log.Printf("Request id ; %v - %v", requestIdval, message)
}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()
// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result with contxt", result)
// 	ctx = context.Background()

// 	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
// 	defer cancel()
// 	result = checkEvenOdd((ctx), 10)
// 	fmt.Println("Result with timeput context", result)
// 	time.Sleep(3 * time.Second)
// 	result = checkEvenOdd((ctx), 10)
// 	fmt.Println("Result with after time  ", result)

// }

//+===== DIFFERENCE BETWEEN toDO AND background()
// func main() {
// 	todoContext := context.TODO()
// 	contextBg := context.Background()
// 	ctx := context.WithValue(todoContext, "name", "everest")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx1 := context.WithValue(contextBg, "city", "Kathmadnu")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))

// }
