package main

import (
	"fmt"
	"maker-checker/rest"
	"maker-checker/service"
)

func main() {
	serviceContainer := service.NewServiceContainer()
	rest.StartServer(serviceContainer)
	fmt.Println("========== Rest Server Started ============")
	select {}
}
