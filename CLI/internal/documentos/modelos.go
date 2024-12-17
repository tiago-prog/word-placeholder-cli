package documentos

// DadosUsuario representa os dados necessários para preenchimento dos documentos
type DadosUsuario struct {
	Nome          string
	StatusCivil   string
	Nacionalidade string
	CPF           string
	Endereco      string
	Cidade        string
	Data          string
}
