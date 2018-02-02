package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ezbuy/redis-orm/fs"
	"github.com/ezbuy/redis-orm/parser"
	"github.com/spf13/viper"
)

func GenerateCode() {
	packageName := viper.GetString("package")

	inputDir, err := filepath.Abs(viper.GetString("code_input"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputDir, err := filepath.Abs(viper.GetString("output"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if packageName == "" {
		_, packageName = path.Split(outputDir)
	}

	yamls, err := fs.GetDirectoryFilesBySuffix(inputDir, ".yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	model := viper.GetString("code_model")
	metaObjs := map[string]*parser.MetaObject{}
	confTpls := map[string]bool{
		"orm": true,
	}
	utilTpls := make(map[string]bool)
	i := 0
	for _, yaml := range yamls {
		objs, err := parser.ReadYaml(packageName, yaml)
		if err != nil {
			fmt.Println("failed: ", err)
			os.Exit(1)
		}

		i = i + 1
		if model != "" {
			for _, obj := range objs {
				if obj.Tag == "" {
					obj.Tag = fmt.Sprint(i)
				}
				if strings.ToLower(obj.Name) == strings.ToLower(model) {
					metaObjs[obj.Name] = obj
					for _, db := range obj.Dbs {
						confTpls[db] = true
						utilTpls[db] = true
					}
					goto GeneratePoint
				}
			}
		} else {
			for _, obj := range objs {
				obj.Tag = fmt.Sprint(i)
				metaObjs[obj.Name] = obj
				for _, db := range obj.Dbs {
					confTpls[db] = true
					utilTpls[db] = true
				}
			}
		}
	}

GeneratePoint:
	for _, metaObj := range metaObjs {
		err = fs.ExecuteMetaObjectCodeTemplate(outputDir, metaObj)
		if err != nil {
			panic(err.Error())
		}
	}

	for conf := range confTpls {
		err = fs.ExecuteConfigTemplate(outputDir, conf, packageName)
		if err != nil {
			panic(err.Error())
		}
	}
	for db := range utilTpls {
		err = fs.ExecuteUtilTemplate(outputDir, db, packageName)
		if err != nil {
			panic(err.Error())
		}
	}

}
