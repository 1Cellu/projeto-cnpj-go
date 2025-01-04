package modules

type CompanyInfo struct {
	// ID                      primitive.ObjectID `bson:"_id"`
	Cnpj                    string `bson:"cnpj"`
	TipoEmpresa             string `bson:"tipoEmpresa"`
	RazaoSocial             string `bson:"razaoSocial"`
	NomeFantasia            string `bson:"nomeFantasia"`
	SituacaoCadastral       string `bson:"situacaoCadastral"`
	DataSituacaoCadastral   string `bson:"dataSituacaoCadastral"`
	MotivoSituacaoCadastral string `bson:"motivoSituacaoCadastral"`
	NaturezaJuridica        string `bson:"naturezaJuridica"`
	DataAbertura            string `bson:"dataAbertura"`
	CnaePrincipal           string `bson:"cnaePrincipal"`
	TipoLogradouro          string `bson:"tipoLogradouro"`
	Logradouro              string `bson:"logradouro"`
	Numero                  string `bson:"numero"`
	Complemento             string `bson:"complemento"`
	Bairro                  string `bson:"bairro"`
	Cep                     string `bson:"cep"`
	Uf                      string `bson:"uf"`
	Municipio               string `bson:"municipio"`
	Telefone                string `bson:"telefone"`
	Telefone2               string `bson:"telefone2"`
	Email                   string `bson:"email"`
	Porte                   string `bson:"porte"`
	SituacaoEspecial        string `bson:"situacaoespecial"`
	DataSituacaoEspecial    string `bson:"datasituacaoespecial"`
}
