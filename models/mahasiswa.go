package models

import "time"

type Mahasiswa struct {
	Nama              string `json:"nama"`
	Usia              int    `json:"usia"`
	Gender            string `json:"gender"`
	TanggalRegistrasi time.Time
}
