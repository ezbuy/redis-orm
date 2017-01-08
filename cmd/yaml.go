// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// yamlCmd represents the yaml command
var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "generate yaml files from database or just sample yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		GenerateYaml()
	},
}

func init() {
	RootCmd.AddCommand(yamlCmd)

	yamlCmd.PersistentFlags().StringP("model", "m", "", "sample yaml file's model name")
	yamlCmd.PersistentFlags().StringP("driver", "D", "mysql", "database driver name, like: mysql, mssql etc")
	yamlCmd.PersistentFlags().StringP("database", "d", "", "database name")
	yamlCmd.PersistentFlags().StringP("host", "H", "localhost", "database host")
	yamlCmd.PersistentFlags().IntP("port", "P", 3306, "database port")
	yamlCmd.PersistentFlags().StringP("username", "u", "root", "database username")
	yamlCmd.PersistentFlags().StringP("password", "p", "", "database password")

	viper.BindPFlag("yaml_model", yamlCmd.PersistentFlags().Lookup("model"))
	viper.BindPFlag("yaml_driver", yamlCmd.PersistentFlags().Lookup("driver"))
	viper.BindPFlag("yaml_host", yamlCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("yaml_port", yamlCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("yaml_database", yamlCmd.PersistentFlags().Lookup("database"))
	viper.BindPFlag("yaml_username", yamlCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("yaml_password", yamlCmd.PersistentFlags().Lookup("password"))
}
