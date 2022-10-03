package repo

import (
	"eta/model"
	"eta/utils"

	"github.com/jinzhu/gorm"
)

type CompanyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) CompanyRepo {
	return CompanyRepo{
		db: db,
	}
}

func (ur *CompanyRepo) Find() (*model.CompanyInfo, error) {
	var resp model.CompanyInfo
	err := ur.db.Raw("EXEC CompanyInfoFind").Row().Scan(
		&resp.EtaRegistrationId,
		&resp.EtaActivityCode,
		&resp.EtaType,
		&resp.ComName,
	)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return &resp, nil
}
