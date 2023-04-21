package controllers

import (
	"api-mahasiswa/models"
	"database/sql"
	"fmt"
	"net/http"

	_ "api-mahasiswa/docs"

	"github.com/gin-gonic/gin"
)

type MahasiswaController struct {
	DB *sql.DB
}

// @Summary menambahkan data mahasiswa baru
// @ID create-mahasiswa
// @Param mahasiswa body models.Mahasiswa true "Menambahkan data mahasiswa"
// @Produce json
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs [post]
func (m *MahasiswaController) Create(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	err := c.BindJSON(&mahasiswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = m.DB.Exec("INSERT INTO mahasiswa (nama, usia, gender, tanggal_registrasi) VALUES (?, ?, ?, ?)", mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, mahasiswa.TanggalRegistrasi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menambahkan data mahasiswa",
	})
}

// @Summary menampilkan semua data mahasiswa
// @ID read-mahasiswa
// @Produce json
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs [get]
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
		"message": "Berhasil menampilkan data mahasiswa",
		"data":    daftarMahasiswa,
	})
}

// @Summary menampilkan data mahasiswa berdasarkan id
// @ID read-mahasiswa-by-id
// @Produce json
// @Param id path int true "Id mahasiswa"
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs/{id} [get]
func (m *MahasiswaController) ReadById(c *gin.Context) {
	id := c.Param("id")

	var mahasiswa models.Mahasiswa

	query := fmt.Sprintf("SELECT * FROM mahasiswa WHERE id = %s", id)
	err := m.DB.QueryRow(query).Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TanggalRegistrasi)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data mahasiswa tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
		"data":    mahasiswa,
	})
}

func (m *MahasiswaController) Update(c *gin.Context) {
	id := c.Param("id")

	var mahasiswa models.Mahasiswa

	querySearch := fmt.Sprintf("SELECT * FROM mahasiswa WHERE id = %s", id)
	row := m.DB.QueryRow(querySearch)
	if row == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Mahasiswa tidak ditemukan",
		})
		return
	}

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	queryUpdate := "UPDATE mahasiswa SET nama=?, usia=?, gender=? WHERE id=?"
	_, err := m.DB.Exec(queryUpdate, mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil diupdate",
	})
}

func (m *MahasiswaController) Delete(c *gin.Context) {
	id := c.Param("id")
	querySearch := fmt.Sprintf("SELECT * FROM mahasiswa WHERE id = %s", id)
	err := m.DB.QueryRow(querySearch)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data mahasiswa tidak ditemukan",
		})
		return
	}

	query, errors := m.DB.Prepare("DELETE FROM mahasiswa WHERE id=?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errors.Error(),
		})
		return
	}

	_, errExt := query.Exec(id)
	if errExt != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errExt.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil dihapus",
	})
}
