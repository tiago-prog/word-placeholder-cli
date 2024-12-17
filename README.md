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
