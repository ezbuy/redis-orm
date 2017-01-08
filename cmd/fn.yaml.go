package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func GenerateYaml() {
	model := viper.GetString("yaml_model")

	fmt.Println("generate yaml => ", model)

	s := make([]string, 6)
	fmt.Println("s =>", strings.Join(s, ","))

}
