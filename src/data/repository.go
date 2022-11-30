package data

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

type Repository interface {
	BatchInsert([]*Data) (int, error)
	FindByYearMonthType(year, month, typ string) (*Data, error)
}

type repositoryImpl struct {
	DB *gorm.DB
}

func (d *repositoryImpl) BatchInsert(items []*Data) (int, error) {
	tx := d.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&items, len(items))
	return int(tx.RowsAffected), tx.Error
}

func (d *repositoryImpl) FindByYearMonthType(y, m, t string) (*Data, error) {
	var (
		result *Data
		begin  time.Time
		end    time.Time
		err    error
	)

	if begin, err = time.Parse("20060102", y+m+"01"); err != nil {
		return result, errors.New("failed to parse y and m: " + y + "." + m)
	}
	end = begin.AddDate(0, 1, 0)
	tx := d.DB.Where("target_type=? AND generated_at >= ? AND generated_at < ?", t, begin, end).First(&result)
	err = tx.Error
	if err != nil && strings.Contains(err.Error(), "record not found") {
		err = fmt.Errorf(fmt.Sprintf("%+v data not found from y:%+v m:%+v", t, y, m))
	}
	return result, err
}

func NewDataRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db}
}
