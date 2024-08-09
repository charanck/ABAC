package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	for {

		fmt.Println("Select below option:")
		fmt.Println("1) generate http server code")
		fmt.Println("2) generate grpc server code")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		switch input {
		case 1:
			fmt.Println("generating http server code")
			cmd := exec.Command("go", "run", "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen", "-config=./codegen-config.yaml", "./openapi.json")
			_, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("code generation done")
			return
		case 2:
			fmt.Println("NOT IMPLEMENTED YET")
			return
		default:
			fmt.Println("invalid option")
			time.Sleep(1 * time.Second)
		}
	}
}
