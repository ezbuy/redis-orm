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
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "generate sql script from yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		GenerateSQL()
	},
}

func init() {
	RootCmd.AddCommand(sqlCmd)

	sqlCmd.PersistentFlags().StringP("input", "i", ".", "directory of yaml files")
	sqlCmd.PersistentFlags().StringP("driver", "d", "mysql", "database type")
	sqlCmd.PersistentFlags().StringP("model", "m", "", "model need to generate sql script")
	viper.BindPFlag("sql_input", sqlCmd.PersistentFlags().Lookup("input"))
	viper.BindPFlag("sql_model", sqlCmd.PersistentFlags().Lookup("model"))
	viper.BindPFlag("sql_driver", sqlCmd.PersistentFlags().Lookup("driver"))
}
