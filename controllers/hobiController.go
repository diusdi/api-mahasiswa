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

// @Tags Mengelola data hobi
// @Summary update data hobi
// @ID update-hobi
// @Produce json
// @Param id path int true "Id hobi"
// @Param mahasiswa body models.Hobi true "Data yang bisa diupdate : nama hobi"
// @Success 200 {string} message
// @Failure 400 {object} error
// @Router /mhs/hobi/{id} [put]
func (m *HobiController) Update(c *gin.Context) {
	id := c.Param("id")

	var hobi models.Hobi

	query := fmt.Sprintf("SELECT id FROM hobi WHERE id = %s", id)
	err := m.DB.QueryRow(query).Scan(&hobi.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data hobi tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&hobi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	queryUpdate := "UPDATE hobi SET nama_hobi=? WHERE id=?"
	_, errUpdate := m.DB.Exec(queryUpdate, hobi.NamaHobi, id)
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data hobi berhasil diupdate",
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
