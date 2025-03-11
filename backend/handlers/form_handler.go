package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

// Submit Form API
func SubmitForm(c *gin.Context) {
	var form models.FormData

	// Debugging - Print JSON data received
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println("❌ Invalid JSON input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	fmt.Println("✅ Received Data:", form.Name, form.Email)

	// Check if Name & Email are empty
	if form.Name == "" || form.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required"})
		return
	}

	// Database Query
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.DB.Exec(query, form.Name, form.Email)
	if err != nil {
		fmt.Println("❌ Database Insert Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	// Fetch Last Inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("❌ Error Fetching Inserted ID:", err)
	} else {
		fmt.Println("✅ Data Inserted with ID:", id)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data submitted successfully",
		"id":      id,
	})
}

// ------------------------------------------------------ to see only last update row
func GetData(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, name, email FROM users ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Println("Database Query Error:", err) // ✅ Debugging ke liye error print karo
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	defer rows.Close()

	var form models.FormData
	if rows.Next() {
		err := rows.Scan(&form.ID, &form.Name, &form.Email)
		if err != nil {
			fmt.Println("Row Scan Error:", err) // ✅ Row scan error print karo
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan data"})
			return
		}
	} else {
		fmt.Println("No data found") // ✅ Debugging ke liye message print karo
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found"})
		return
	}

	fmt.Println("Fetched Data:", form) // ✅ Data successfully fetch hua ya nahi, ye check karo
	c.JSON(http.StatusOK, form)
}

/*

//-------------------   mysql ki table ka saara data dekhena ho

func GetData(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, name, email FROM users") // Data DESC order me fetch kar raha hu
	if err != nil {
		fmt.Println("Database Query Error:", err) // ✅ Debugging ke liye error print karo
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	defer rows.Close()

	var formDataList []models.FormData

	for rows.Next() {
		var form models.FormData
		err := rows.Scan(&form.ID, &form.Name, &form.Email)
		if err != nil {
			fmt.Println("Row Scan Error:", err) // ✅ Row scan error print karo
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan data"})
			return
		}
		formDataList = append(formDataList, form)
	}

	// Agar koi data nahi mila
	if len(formDataList) == 0 {
		fmt.Println("No data found") // ✅ Debugging ke liye message print karo
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found"})
		return
	}

	fmt.Println("Fetched Data:", formDataList) // ✅ Saara data fetch hone ka confirmation
	c.JSON(http.StatusOK, formDataList)        // ✅ Saara data JSON response me bhej raha hu
}



*/

func DownloadCSV(c *gin.Context) {
	rows, _ := db.DB.Query("SELECT id, name, email FROM users")
	defer rows.Close()

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=data.csv")

	writer := csv.NewWriter(c.Writer)
	writer.Write([]string{"ID", "Name", "Email"})

	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		writer.Write([]string{fmt.Sprint(id), name, email})
	}

	writer.Flush()
}

func DownloadPDF(c *gin.Context) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Submitted Data")
	pdf.Ln(10)

	rows, _ := db.DB.Query("SELECT id, name, email FROM users")
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		pdf.Cell(40, 10, fmt.Sprintf("%d - %s - %s", id, name, email))
		pdf.Ln(10)
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment;filename=data.pdf")
	_ = pdf.Output(c.Writer)
}
