package deploy

import (
	_ "embed"
	"os"
)

//go:embed deploy.sh
var data []byte

func init() {
	err := os.Mkdir("./packagedeploy", 0744)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./packagedeploy/deploy.sh", data, 0544)
	if err != nil {
		panic(err)
	}
}
