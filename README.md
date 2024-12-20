# 📝 **CLI Gerador de Documentos**

Este projeto é uma CLI (Interface de Linha de Comando) feita em Go, que permite gerar documentos personalizados a partir de modelos de procuração. É a solução perfeita para automatizar a criação de documentos legais e outros tipos de arquivos com placeholders! ⚙️

## 💡 **O que é?**

Este projeto permite que você crie documentos docx substituindo **placeholders** com os dados fornecidos pelo usuário, como nome, CPF e endereço ou qualquer outra informação. Essa cli foi pensanda para uso exclusivo e é mais engessado no quesito do placeholder, ele ta programado para procurar por dados como nome, cpf, endereco, cidade e data. Mas não impende de você reescrever o código a seu belprazer.

## 🛠️ **Frameworks/bibliotecas**

-  **Golang cobra:** feito para criação de CLI [https://cobra.dev/]🖋️
-  

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
 - gerar procuracao 
    Gera um documento de procuração preenchendo os placeholders com as informações fornecidas pelo usuário.
   
```bash
go run main.go gerar procuracao --modelo [masculino|feminino] --nome [nome] --status [status_civil] --nacionalidade [nacionalidade] --cpf [cpf] --endereco [endereco] --cidade [cidade] --data [data]
```

## 4️⃣ Exemplo de Execução
Imagine que você quer gerar uma procuração para João da Silva. O comando seria:

```bash
go run main.go gerar procuracao --modelo masculino --nome "João da Silva" --status solteiro --nacionalidade brasileiro --cpf "123.456.789-10" --endereco "Rua A, 123" --cidade "São Paulo" --data "17/12/2024"
```

Isso geraria o arquivo de procuração na pasta output.

## 📝 Importante
Os arquivos de modelo não estão incluídos neste repositório devido a questões de privacidade/licenciamento.
Você precisará fornecer seus próprios arquivos de modelo para que o gerador funcione corretamente.🖼️



