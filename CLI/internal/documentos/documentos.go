package documentos

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"baliance.com/gooxml/document"
)

// GerarProcuracao gera um documento de procuração baseado em um modelo e dados fornecidos
func GerarProcuracao(modeloTipo string, dados DadosUsuario, outputPath string) error {
	// Define o caminho do modelo baseado no tipo
	modeloPath := fmt.Sprintf("internal/documentos/modelo-%s.docx", modeloTipo)

	// Abre o modelo DOCX
	doc, err := document.Open(modeloPath)
	if err != nil {
		return fmt.Errorf("erro ao abrir o modelo: %w", err)
	}

	// Substitui os placeholders pelos dados do usuário
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			text := r.Text()
			if strings.TrimSpace(text) == "{{CIDADE}}" {
				r.Clear()
				r.AddText(" " + dados.Cidade)
			}
			switch text {
			case "{{NOME}}":
				r.Clear()
				r.AddText(strings.ToUpper(dados.Nome))
			case "{{STATUS_CIVIL}}":
				r.Clear()
				r.AddText(dados.StatusCivil)
			case "{{NACIONALIDADE}}":
				r.Clear()
				r.AddText(dados.Nacionalidade)
			case "{{CPF}}":
				r.Clear()
				r.AddText(dados.CPF)
			case "{{ENDERECO}}":
				r.Clear()
				r.AddText(dados.Endereco)
			case "{{DATA}}":
				r.Clear()
				r.AddText(dados.Data)
			}
		}
	}

	// Salva o documento gerado no caminho especificado
	outputDir := filepath.Dir(outputPath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.MkdirAll(outputDir, os.ModePerm)
	}

	if err := doc.SaveToFile(outputPath); err != nil {
		return fmt.Errorf("erro ao salvar o documento: %w", err)
	}

	return nil
}
