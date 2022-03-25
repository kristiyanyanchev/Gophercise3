package main

import "github.com/kristiyanyanchev/Gophercise3/api"

func main() {
	go api.StartServer(9019)
}

// func main() {
// 	go func() {
// 		for i := 0; i < 1000000000; i++ {

// 			fmt.Print("ocean")
// 		}
// 	}()
// 	go func() {
// 		for i := 0; i < 1000000000; i++ {

// 			os.WriteFile("sea.txt", []byte("sea"), os.ModeAppend)
// 		}

// 	}()
// }
