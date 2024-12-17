# 📝 **CLI Gerador de Documentos**

Bem-vindo ao **Gerador de Documentos**! 🚀  
Este projeto é uma CLI (Interface de Linha de Comando) feita em Go, que permite gerar documentos personalizados a partir de modelos de procuração. É a solução perfeita para automatizar a criação de documentos legais e outros tipos de arquivos com placeholders! ⚙️

## 💡 **O que é?**

Este projeto permite que você crie documentos como **Procurações** usando um modelo e substituindo os campos de placeholder com os dados fornecidos pelo usuário, como nome, CPF, endereço e muito mais! 🙌

## 🛠️ **Funcionalidades**

- Geração de **Procuração** personalizada 🖋️
- Suporte para modelos **masculino** e **feminino** 📄
- Armazenamento do documento gerado em uma pasta `output` 📂
- Detecção automática de nomes duplicados (ex: `tiago`, `tiago (1)`) 📑

## ⚡ **Como Usar?**

### 1️⃣ **Instale as dependências**

Certifique-se de que você tem o Go instalado em sua máquina. Se ainda não, você pode instalá-lo a partir de [aqui](https://go.dev/dl/). 🌍

Em seguida, instale as dependências do projeto com o comando:

```bash
go mod tidy 

```

## 2️⃣ Executar a CLI
Após configurar tudo, você pode rodar o comando da CLI com o seguinte formato:

```bash
go run main.go [comando] [flags]
```


## 3️⃣ Comandos Disponíveis:
 - gerar procuracao 🎯
    Gera um documento de procuração preenchendo os placeholders com as informações fornecidas pelo usuário.

go run main.go gerar procuracao --modelo [masculino|feminino] --nome [nome] --status [status_civil] --nacionalidade [nacionalidade] --cpf [cpf] --endereco [endereco] --cidade [cidade] --data [data]

## 4️⃣ Exemplo de Execução
Imagine que você quer gerar uma procuração para João da Silva. O comando seria:

```bash
go run main.go gerar procuracao --modelo masculino --nome "João da Silva" --status solteiro --nacionalidade brasileiro --cpf "123.456.789-10" --endereco "Rua A, 123" --cidade "São Paulo" --data "17/12/2024"
```

Isso geraria o arquivo de procuração na pasta output.

## 🔄 Verificando a pasta de saída:
O documento gerado será salvo na pasta output com o nome do usuário. Caso já exista um documento com o mesmo nome, ele será renomeado automaticamente como nome (1) para evitar sobrescrever. 🚀

## 🛠️ Comandos Avançados:
procuracao 📜: Gera o documento de procuração baseado no modelo e dados fornecidos.
help ❓: Exibe ajuda sobre como usar a CLI.
🔧 Como Funciona?
O processo funciona da seguinte forma:

Escolha do Modelo: O usuário escolhe entre modelo masculino ou feminino.
Substituição dos Placeholders: O programa substitui os placeholders no modelo de documento com os dados fornecidos (nome, CPF, endereço, etc.).
Geração do Documento: O documento é salvo na pasta output com um nome único, evitando sobreposições.

## 🚀 Exemplo de saída do Documento

Procuração:
```vbnet
OUTORGANTE: João da Silva, brasileiro, solteiro, inscrito no CPF nº 123.456.789-10, residente e domiciliado na Rua A, 123, cidade de São Paulo. 

OUTORGADO (s): [Nome do Advogado] com sede na Avenida Jardim Osorio, nº 231-Bairro Centro-São Paulo-SP, inscrita na OAB/RS sob nº 12.345...

PODERES: Por este instrumento particular de mandato, para fim específico de ajuizar AÇÃO JUNTO AO INSS...

{{DATA}}.
```

A partir desse modelo, os placeholders são preenchidos com os dados fornecidos e um novo arquivo DOCX é gerado com as informações.

## 📂 Estrutura do Projeto

```bash
/cmd
  /gerar-procuracao.go    # Comando para gerar procuração
/internal
  /documentos
    gerar.go              # Função que lida com a geração do documento
    dados.go              # Função que lida com os dados do usuário
  /models
    modelo-masculino.docx # Modelo para masculino (não incluído no repositório)
    modelo-feminino.docx  # Modelo para feminino (não incluído no repositório)
/output
  # Arquivos gerados aqui
```


## 📝 Importante
Os arquivos de modelo modelo-masculino.docx e modelo-feminino.docx não estão incluídos neste repositório devido a questões de privacidade/licenciamento.
Você precisará fornecer seus próprios arquivos de modelo para que o gerador funcione corretamente.🖼️



