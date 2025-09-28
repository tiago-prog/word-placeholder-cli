---

# 📝 CLI Gerador de Documentos

Uma **interface de linha de comando (CLI)** desenvolvida em **Go** com o framework **Cobra**, projetada para **automatizar a geração de documentos personalizados** (como procurações) a partir de modelos `.docx`.

## 💡 Funcionalidade
- Substitui **placeholders fixos** no modelo (como `nome`, `cpf`, `endereco`, `cidade`, `data`, etc.) pelos dados fornecidos via linha de comando.
- Atualmente otimizada para documentos legais, mas o código pode ser adaptado para outros usos.

## 🛠️ Tecnologias
- **Go** (linguagem)
- **Cobra** (para criação da CLI)

## ⚡ Como Usar
1. Instale as dependências:
   ```bash
   go mod tidy
   ```
2. Execute o comando:
   ```bash
   go run main.go gerar procuracao --modelo [masculino|feminino] --nome "..." --status ... --nacionalidade ... --cpf ... --endereco ... --cidade ... --data ...
   ```

### 📌 Exemplo:
```bash
go run main.go gerar procuracao --modelo masculino --nome "João da Silva" --status solteiro --nacionalidade brasileiro --cpf "123.456.789-10" --endereco "Rua A, 123" --cidade "São Paulo" --data "17/12/2024"
```
→ Gera o documento preenchido na pasta `output/`.

## ⚠️ Observação
- **Modelos `.docx` não estão incluídos** no repositório por questões de privacidade/licenciamento.  
- Você deve fornecer seus próprios arquivos de modelo com os placeholders esperados.

--- 

Pronto para uso ou personalização! 🚀
