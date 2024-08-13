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
			generateHTTPServerCode()
			return
		case 2:
			generateGRPCServerCode()
			return
		default:
			fmt.Println("invalid option")
			time.Sleep(1 * time.Second)
		}
	}
}

func generateHTTPServerCode() {
	fmt.Println("generating http server code")
	cmd := exec.Command("go", "run", "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen", "-config=./codegen-config.yaml", "./openapi.json")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("http code generation failed")
		return
	}
	fmt.Println("code generation done")
}

func generateGRPCServerCode() {
	fmt.Println("generating grpc server code")

	cmd := exec.Command("protoc", "--go_out=./protobuf/generated", "--go-grpc_out=./protobuf/generated", "./protobuf/abac.proto")
	cmd.Dir = "../"
	x, err := cmd.Output()
	fmt.Println(string(x))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("grpc code generation failed")
		return
	}
	fmt.Println("grpc code generation done")
}
