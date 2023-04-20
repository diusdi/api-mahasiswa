package models

import "time"

type Mahasiswa struct {
	Id                int    `json:"id"`
	Nama              string `json:"nama"`
	Usia              int    `json:"usia"`
	Gender            string `json:"gender"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi"`
}
