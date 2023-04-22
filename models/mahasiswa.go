package models

import "time"

type Mahasiswa struct {
	Id                int       `json:"id" example:"1"`
	Nama              string    `json:"nama" example:"Dion"`
	Usia              int       `json:"usia" example:"21"`
	Gender            string    `json:"gender" example:"1"`
	IsActive          string    `json:"is_active" example:"1"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi" example:"2020-01-02T15:04:05Z"`
}
