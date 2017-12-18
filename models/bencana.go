package models

import ()

type Bencana struct {
	ID        int    `json:"id"`
	Kota      string `json:"kota"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Tewas     string `json:"tewas"`
	Luka      string `json:"luka"`
	Kerusakan string `json:"kerusakan"`
}

func GetSeveralBencana() ([]*Bencana, error) {
	bencana := []*Bencana{}
	err = db.Find(&bencana).Error
	return bencana, err
}

func GetBencana(id int) (*Bencana, error) {
	bencana := &Bencana{}
	err = db.Where("id = ?", id).First(&bencana).Error
	return bencana, err
}

func CreateBencana(bencana *Bencana) (*Bencana, error) {
	err = db.Create(bencana).Error
	return bencana, err
}

func UpdateBencana(bencana *Bencana) (*Bencana, error) {
	err = db.Save(bencana).Error
	return bencana, err
}

func DeleteBencana(email string) error {
	err = db.Where("email = ?", email).Delete(&Bencana{}).Error
	return err
}
