package main

import (
	"bytes"
	"fmt"
	"nft_service/app/global/variable"
	_ "nft_service/bootstrap"
	"nft_service/routers"
	"os"
	"os/exec"
)

// @title Revolution service API
// @version 1.0
// @description  Revolution service API
// @termsOfService http://github.com
// @contact.name API Support
// @contact.url Fuck!
// @contact.email U!
func main() {
	router := routers.InitRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Web.Port"))
}

func test() {
	cmd := exec.Command("swag", "init")
	fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}
