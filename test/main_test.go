package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kristiyanyanchev/Gophercise3/api"
)

// func TestNoEndpoint(t *testing.T) {
// 	go api.StartServer()
// 	internalNoEndpoint(t)
// }
// func internalNoEndpoint(t *testing.T) {
// 	resp, err := http.Get("http://localhost:9021/new-york")
// 	if err != nil {
// 		t.Fatalf(err.Error())
// 	}
// 	if resp.StatusCode != 200 {
// 		t.Fatalf("Error status code is " + resp.Status)
// 	}
// }

func TestJsonStoryMap_(t *testing.T) {
	str := api.JsonToStoryMap("../api/stories.json")
	if str == nil {
		t.Fatal("cound open file")
	}

}
func Test(t *testing.T) {
	go func() {
		fmt.Println("rts")
		os.WriteFile("sss.txt", []byte("sea"), 777)
		t.Errorf("str")
	}()
	fmt.Println("fta")
}
