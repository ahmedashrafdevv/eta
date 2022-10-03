package repo

import (
	"database/sql"
	"eta/model"
	"eta/utils"

	"github.com/jinzhu/gorm"
)

type StoreRepo struct {
	db *gorm.DB
}

func NewStoreRepo(db *gorm.DB) StoreRepo {
	return StoreRepo{
		db: db,
	}
}

func (ur *StoreRepo) ListAll() (*[]model.Store, error) {
	rows, err := ur.db.Raw("EXEC StoreCodeList").Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	if utils.CheckErr(&err) {
		return nil, err
	}
	result, err := scanStoreResult(rows)
	return result, nil
}

func scanStoreResult(rows *sql.Rows) (*[]model.Store, error) {
	var resp []model.Store
	for rows.Next() {
		var rec model.Store
		err := rows.Scan(&rec.StoreCode, &rec.StoreName)
		if utils.CheckErr(&err) {
			return nil, err
		}
		resp = append(resp, rec)

	}

	return &resp, nil
}
