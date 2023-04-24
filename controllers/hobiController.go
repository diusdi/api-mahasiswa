package controllers

import (
	"api-mahasiswa/models"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	_ "api-mahasiswa/docs"

	"github.com/gin-gonic/gin"
)

type HobiController struct {
	DB *sql.DB
}

// @Tags Mengelola data hobi
// @Summary menambahkan data hobi baru
// @ID create-hobi
// @Param hobi body models.Hobi true "Data yang bisa ditambahkan : nama hobi"
// @Produce json
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs/hobi [post]
func (m *HobiController) Create(c *gin.Context) {
	var hobi models.Hobi

	err := c.BindJSON(&hobi)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	namaHobi := strings.ToLower(hobi.NamaHobi)
	query := fmt.Sprintf("SELECT nama_hobi FROM hobi WHERE nama_hobi='%s'", namaHobi)
	errFind := m.DB.QueryRow(query).Scan(&hobi.NamaHobi)
	if errFind == nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hobi sudah ada"})
		return
	}

	_, err = m.DB.Exec("INSERT INTO hobi (nama_hobi) VALUES (?)", hobi.NamaHobi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menambahkan data hobi",
	})
}

// @Tags Mengelola data hobi
// @Summary menampilkan semua data hobi
// @ID read-hobi
// @Produce json
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs/hobi [get]
func (m *HobiController) Read(c *gin.Context) {
	rows, err := m.DB.Query("SELECT * FROM hobi")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	daftarHobi := []models.Hobi{}
	for rows.Next() {
		var hobi models.Hobi
		fmt.Printf("rows: %v\n", rows)
		err := rows.Scan(&hobi.Id, &hobi.NamaHobi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		daftarHobi = append(daftarHobi, hobi)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menampilkan data mahasiswa",
		"data":    daftarHobi,
	})
}

// @Summary menampilkan data mahasiswa berdasarkan id
// @ID read-mahasiswa-by-id
// @Produce json
// @Param id path int true "Id mahasiswa"
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs/{id} [get]
func (m *HobiController) ReadById(c *gin.Context) {
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
func (m *HobiController) Update(c *gin.Context) {
	id := c.Param("id")

	var mahasiswa models.Mahasiswa

	query := fmt.Sprintf("SELECT id FROM mahasiswa WHERE id = %s AND is_active = '1'", id)
	err := m.DB.QueryRow(query).Scan(&mahasiswa.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data mahasiswa tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	queryUpdate := "UPDATE mahasiswa SET nama=?, usia=?, gender=? WHERE id=?"
	_, errUpdate := m.DB.Exec(queryUpdate, mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, id)
	if errUpdate != nil {
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
func (m *HobiController) Delete(c *gin.Context) {
	id := c.Param("id")

	var mahasiswa models.Mahasiswa

	query := fmt.Sprintf("SELECT id FROM mahasiswa WHERE id = %s AND is_active = '1'", id)
	err := m.DB.QueryRow(query).Scan(&mahasiswa.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data mahasiswa tidak ditemukan"})
		return
	}

	// fitur soft delete
	queryDelete := "UPDATE mahasiswa SET is_active=? WHERE id=?"
	_, errDelete := m.DB.Exec(queryDelete, "0", id)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data mahasiswa berhasil dihapus",
	})
}
