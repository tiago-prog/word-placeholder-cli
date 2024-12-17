package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"go.mod/CLI/internal/documentos"
)

var (
	modeloTipo    string
	nome          string
	statusCivil   string
	nacionalidade string
	cpf           string
	endereco      string
	cidade        string
	data          string
)

// procuracaoCmd representa o subcomando "gerar procuracao"
var procuracaoCmd = &cobra.Command{
	Use:   "procuracao",
	Short: "Gera um documento de procuração",
	Long:  "Gera uma procuração preenchendo os placeholders do modelo com base nos dados fornecidos.",
	Run: func(cmd *cobra.Command, args []string) {
		// Dados do usuário
		dados := documentos.DadosUsuario{
			Nome:          nome,
			StatusCivil:   statusCivil,
			Nacionalidade: nacionalidade,
			CPF:           cpf,
			Endereco:      endereco,
			Cidade:        cidade,
			Data:          data,
		}

		// Define o diretório de saída
		outputDir := "output"
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			os.Mkdir(outputDir, os.ModePerm)
		}

		// Define o caminho do arquivo de saída
		outputPath := filepath.Join(outputDir, fmt.Sprintf("procuracao-%s.docx", dados.Nome))

		// Verifica se o arquivo já existe e cria um nome único
		counter := 1
		for fileExists(outputPath) {
			// Se o arquivo existir, adiciona um sufixo "(n)"
			outputPath = fmt.Sprintf("%s (%d).docx", strings.TrimSuffix(outputPath, ".docx"), counter)
			counter++
		}

		// Gera a procuração usando os dados fornecidos
		if err := documentos.GerarProcuracao(modeloTipo, dados, outputPath); err != nil {
			fmt.Println("Erro ao gerar procuração:", err)
			os.Exit(1)
		}

		// Confirma a geração do documento
		fmt.Println("Procuração gerada com sucesso:", outputPath)
	},
}

// Função auxiliar para verificar se o arquivo já existe
func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func init() {
	// Adiciona o comando "procuracao" ao comando principal "gerar"
	gerarCmd.AddCommand(procuracaoCmd)

	// Define as flags do comando "procuracao"
	procuracaoCmd.Flags().StringVarP(&modeloTipo, "modelo", "m", "", "Tipo de modelo (masculino/feminino) [obrigatório]")
	procuracaoCmd.Flags().StringVarP(&nome, "nome", "n", "", "Nome do usuário [obrigatório]")
	procuracaoCmd.Flags().StringVarP(&statusCivil, "status", "s", "", "Status civil [obrigatório]")
	procuracaoCmd.Flags().StringVarP(&nacionalidade, "nacionalidade", "a", "", "Nacionalidade [obrigatório]")
	procuracaoCmd.Flags().StringVarP(&cpf, "cpf", "c", "", "CPF [obrigatório]")
	procuracaoCmd.Flags().StringVarP(&endereco, "endereco", "e", "", "Endereço [obrigatório]")
	procuracaoCmd.Flags().StringVarP(&cidade, "cidade", "i", "", "Cidade [obrigatório]")
	procuracaoCmd.Flags().StringVarP(&data, "data", "d", "", "Data [obrigatório]")

	// Marca as flags como obrigatórias
	procuracaoCmd.MarkFlagRequired("modelo")
	procuracaoCmd.MarkFlagRequired("nome")
	procuracaoCmd.MarkFlagRequired("status")
	procuracaoCmd.MarkFlagRequired("nacionalidade")
	procuracaoCmd.MarkFlagRequired("cpf")
	procuracaoCmd.MarkFlagRequired("endereco")
	procuracaoCmd.MarkFlagRequired("cidade")
	procuracaoCmd.MarkFlagRequired("data")
}
