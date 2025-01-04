package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"projeto-cnpj-go/internal/modules"
)

var db *mongo.Database

const (
	collectionName = "record"
	limite         = 10
)

func New(cli *mongo.Client) {
	db = cli.Database("testdb")
}

func ListRecords() ([]modules.CompanyInfo, error) {
	coll := getCollection(collectionName)

	limit := int64(limite)
	var results []modules.CompanyInfo
	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find().SetLimit(limit))
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem modules.CompanyInfo
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())
	return results, nil
}

func GetRecord(cnpj string) (*modules.CompanyInfo, error) {
	coll := getCollection(collectionName)
	filter := bson.D{{Key: "cnpj", Value: cnpj}}

	var companyInfo modules.CompanyInfo

	err := coll.FindOne(context.TODO(), filter).Decode(&companyInfo)
	if err != nil {
		return nil, err
	}

	return &companyInfo, nil
}

func AddRecord(key string, companyInfo modules.CompanyInfo) error {
	coll := getCollection(collectionName)

	_, err := coll.InsertOne(context.TODO(), bson.M{
		"cnpj":                    companyInfo.Cnpj,
		"tipoEmpresa":             companyInfo.TipoEmpresa,
		"razaoSocial":             companyInfo.RazaoSocial,
		"nomeFantasia":            companyInfo.NomeFantasia,
		"situacaoCadastral":       companyInfo.SituacaoCadastral,
		"dataSituacaoCadastral":   companyInfo.DataSituacaoCadastral,
		"motivoSituacaoCadastral": companyInfo.MotivoSituacaoCadastral,
		"naturezaJuridica":        companyInfo.NaturezaJuridica,
		"dataAbertura":            companyInfo.DataAbertura,
		"cnaePrincipal":           companyInfo.CnaePrincipal,
		"tipoLogradouro":          companyInfo.TipoLogradouro,
		"logradouro":              companyInfo.Logradouro,
		"numero":                  companyInfo.Numero,
		"complemento":             companyInfo.Complemento,
		"bairro":                  companyInfo.Bairro,
		"cep":                     companyInfo.Cep,
		"uf":                      companyInfo.Uf,
		"municipio":               companyInfo.Municipio,
		"telefone":                companyInfo.Telefone,
		"telefone2":               companyInfo.Telefone2,
		"email":                   companyInfo.Email,
		"porte":                   companyInfo.Porte,
		"situacaoespecial":        companyInfo.SituacaoEspecial,
		"datasituacaoespecial":    companyInfo.DataSituacaoEspecial,
	})
	if err != nil {
		return err
	}

	return nil
}

func UpdateRecord(cnpj string, companyInfo modules.CompanyInfo) error {
	coll := getCollection(collectionName)

	filter := bson.M{"cnpj": cnpj}

	update := bson.M{"$set": bson.M{
		"cnpj":                    companyInfo.Cnpj,
		"tipoEmpresa":             companyInfo.TipoEmpresa,
		"razaoSocial":             companyInfo.RazaoSocial,
		"nomeFantasia":            companyInfo.NomeFantasia,
		"situacaoCadastral":       companyInfo.SituacaoCadastral,
		"dataSituacaoCadastral":   companyInfo.DataSituacaoCadastral,
		"motivoSituacaoCadastral": companyInfo.MotivoSituacaoCadastral,
		"naturezaJuridica":        companyInfo.NaturezaJuridica,
		"dataAbertura":            companyInfo.DataAbertura,
		"cnaePrincipal":           companyInfo.CnaePrincipal,
		"tipoLogradouro":          companyInfo.TipoLogradouro,
		"logradouro":              companyInfo.Logradouro,
		"numero":                  companyInfo.Numero,
		"complemento":             companyInfo.Complemento,
		"bairro":                  companyInfo.Bairro,
		"cep":                     companyInfo.Cep,
		"uf":                      companyInfo.Uf,
		"municipio":               companyInfo.Municipio,
		"telefone":                companyInfo.Telefone,
		"telefone2":               companyInfo.Telefone2,
		"email":                   companyInfo.Email,
		"porte":                   companyInfo.Porte,
		"situacaoespecial":        companyInfo.SituacaoEspecial,
		"datasituacaoespecial":    companyInfo.DataSituacaoEspecial,
	}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRecord(cnpj string) error {
	coll := getCollection(collectionName)

	result, err := coll.DeleteOne(context.TODO(), bson.M{"cnpj": cnpj})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no item was found to be deleted")
	}

	return nil
}

func getCollection(name string) *mongo.Collection {
	return db.Collection(name)
}
