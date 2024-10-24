package initialize

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/devpenguin/go-ecommerce/global"
	"gopkg.in/yaml.v3"
)

func LoadConfig(env string) {
	yamlFile, err := os.ReadFile(fmt.Sprintf("./configs/%s.yaml", env))
	if err != nil {
        log.Fatalf("Error reading YAML file:: %v", err)
    }

	// Parse the YAML file into the Config structure
	err = yaml.Unmarshal(yamlFile, &global.Config)
	if err != nil {
		log.Fatalf("Error parsing YAML file:: %v", err)
	}
	// read server configuration
	fmt.Println("Server port::", strconv.Itoa(global.Config.Server.Port))
	fmt.Println("MODE::", global.Config.Server.Mode)
}