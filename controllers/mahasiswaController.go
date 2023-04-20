package controllers

import (
	"api-mahasiswa/models"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type MahasiswaController struct {
	DB *sql.DB
}

func (m *MahasiswaController) Create(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	err := c.BindJSON(&mahasiswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tanggalRegistrasi := time.DateTime

	res, err := m.DB.Exec("INSERT INTO mahasiswa (nama, usia, gender, tanggal_registrasi) VALUES (?, ?, ?, ?)", mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, tanggalRegistrasi)
	_ = res
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
	})
}
