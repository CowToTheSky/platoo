package cmd

import "github.com/spf13/cobra"

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use:   "plato",
	Short: "优版权对接项目",
	Run:   Plato,
}

func Plato(cmd *cobra.Command, args []string) {

}

func initConfig() {

}
