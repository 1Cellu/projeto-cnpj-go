package services

import (
	"go.mongodb.org/mongo-driver/mongo"

	"projeto-cnpj-go/internal/modules"
	"projeto-cnpj-go/internal/repository"
)

type Service struct{}

func (s *Service) List() ([]modules.CompanyInfo, error) {
	items, err := repository.ListRecords()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(cnpj string) (*modules.CompanyInfo, error) {
	item, err := repository.GetRecord(cnpj)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) AddRecord(cnpj string, companyInfo modules.CompanyInfo) error {
	_, err := repository.GetRecord(cnpj)
	if err == mongo.ErrNoDocuments {
		err := repository.AddRecord(cnpj, companyInfo)
		if err != nil {
			return err
		}
		return nil
	}

	err = repository.UpdateRecord(cnpj, companyInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(cnpj string) error {
	err := repository.DeleteRecord(cnpj)
	if err != nil {
		return err
	}
	return nil
}
