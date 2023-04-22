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
		err := rows.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.IsActive, &mahasiswa.TanggalRegistrasi)
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

	query := fmt.Sprintf("SELECT nama, usia, gender, tanggal_registrasi FROM mahasiswa WHERE id = %s AND is_active = '1'", id)
	err := m.DB.QueryRow(query).Scan(&mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TanggalRegistrasi)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data mahasiswa tidak ditemukan"})
		return
	}

	var gender map[string]string
	if mahasiswa.Gender == "1" {
		gender = map[string]string{
			"1": "laki-laki",
		}
	} else {
		gender = map[string]string{
			"0": "perempuan",
		}
	}

	tanggalRegistrasi := mahasiswa.TanggalRegistrasi.Format("02-01-2006")
	data := map[string]any{
		"nama":               mahasiswa.Nama,
		"usia":               mahasiswa.Usia,
		"gender":             gender,
		"tanggal_registrasi": tanggalRegistrasi,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil menampilkan data",
		"data":    data,
	})
}

// @Summary update data mahasiswa
// @ID update-mahasiswa
// @Produce json
// @Param id path int true "Id mahasiswa"
// @Param mahasiswa body models.Mahasiswa true "Data yang bisa diupdate : nama, usia, gender ('0' untuk perempuan dan '1' untuk laki-laki)"
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs/{id} [put]
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
		"message": "Data mahasiswa berhasil diupdate",
	})
}

// @Summary menghapus data mahasiswa
// @ID delete-mahasiswa
// @Produce json
// @Param id path int true "Id mahasiswa"
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs/{id} [delete]
func (m *MahasiswaController) Delete(c *gin.Context) {
	id := c.Param("id")

	query, err := m.DB.Prepare("DELETE FROM mahasiswa WHERE id=?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	query.Exec(id)
	// if errExt != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil dihapus",
	})
}
