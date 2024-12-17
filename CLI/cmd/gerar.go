package cmd

import (
	"github.com/spf13/cobra"
)

// gerarCmd representa o comando "gerar"
var gerarCmd = &cobra.Command{
	Use:   "gerar",
	Short: "Gera documentos baseados em modelos",
	Long:  "O comando gerar suporta subcomandos para criação de diferentes documentos.",
}

func init() {
	rootCmd.AddCommand(gerarCmd)
}
