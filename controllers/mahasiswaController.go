package controllers

import (
	"api-mahasiswa/models"
	"database/sql"
	"fmt"
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

	_, err = m.DB.Exec("INSERT INTO mahasiswa (nama, usia, gender, tanggal_registrasi) VALUES (?, ?, ?, ?)", mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, tanggalRegistrasi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
	})
}

func (m *MahasiswaController) Read(c *gin.Context) {
	rows, err := m.DB.Query("SELECT * FROM mahasiswa")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	daftarMahasiswa := []models.Mahasiswa{}
	for rows.Next() {
		var mahasiswa models.Mahasiswa
		fmt.Printf("rows: %v\n", rows)
		err := rows.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TanggalRegistrasi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		daftarMahasiswa = append(daftarMahasiswa, mahasiswa)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
		"data":    daftarMahasiswa[0].Nama,
		"rows":    rows,
	})
}

func (m *MahasiswaController) ReadById(c *gin.Context) {
	id := c.Param("id")

	var mahasiswa models.Mahasiswa

	query := fmt.Sprintf("SELECT * FROM mahasiswa WHERE id = %s", id)
	err := m.DB.QueryRow(query).Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TanggalRegistrasi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
		"data":    mahasiswa,
	})
}
