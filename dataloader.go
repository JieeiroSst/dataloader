package dataloader

import (
	"gorm.io/gorm"
)

type DataLoader struct {
	db *gorm.DB
}

type Ids struct {
	id int
}

type Data struct {
	node interface{}
}

func (d *DataLoader) GetById(table interface{})([]Ids,error){
	var ids []Ids
	if err:=d.db.Model(table).Select("id").Scan(&ids).Error ; err!=nil{
		return nil, err
	}
	return ids,nil
}

func (d *DataLoader) DataLoader(table interface{},ids []int) ([]Data,error) {
	var (
		data Data
		datasets []Data
	)
	for _,id:=range ids {
		if err:=d.db.Model(table).Where("id = ?",id).Find(&data).Error;err!=nil{
			return nil,err
		}
		datasets = append(datasets,data)
	}
	return datasets,nil
}

func (d *DataLoader) CreateDataLoader(table interface{},ids []int,key string,data map[string]interface{}) error {
	for _,id:=range ids {
		data[key] = id
		if err:=d.db.Model(table).Create(data).Error;err!=nil{
			return err
		}
	}
	return nil
}