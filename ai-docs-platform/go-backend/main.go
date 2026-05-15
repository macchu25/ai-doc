package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Project struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Status     string   `json:"status"`
	Frameworks []string `json:"frameworks"`
	Path       string   `json:"path"`
}

var projects = make(map[string]Project)
var uploadDir = "./uploads"

func main() {
	// Create upload directory
	os.MkdirAll(uploadDir, os.ModePerm)

	// Initialize Qdrant
	InitQdrant()

	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-API-Key"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "AI Software Documentation Platform Go API is running",
		})
	})

	r.POST("/upload", handleUpload)
	r.GET("/projects", listProjects)
	r.GET("/analyze/:id", analyzeProject)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}

func handleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	if filepath.Ext(file.Filename) != ".zip" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only ZIP files are supported"})
		return
	}

	projectID := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
	zipPath := filepath.Join(uploadDir, file.Filename)
	extractPath := filepath.Join(uploadDir, projectID)
	
	if err := c.SaveUploadedFile(file, zipPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Unzip the project
	_, err = Unzip(zipPath, extractPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract ZIP: " + err.Error()})
		return
	}

	// Scan the project
	_, frameworks := ScanCodebase(extractPath)
	
	p := Project{
		ID:         projectID,
		Name:       projectID,
		Status:     "Analyzing",
		Frameworks: frameworks,
		Path:       extractPath,
	}
	projects[projectID] = p

	// Trigger Indexing in background
	go func() {
		err := IndexProject(p.ID, p.Path)
		if err != nil {
			fmt.Printf("Indexing error for %s: %v\n", p.ID, err)
			return
		}
		// Update status when done
		p.Status = "Ready"
		projects[p.ID] = p
		fmt.Printf("Successfully indexed project: %s\n", p.ID)
	}()

	c.JSON(http.StatusOK, p)
}

func listProjects(c *gin.Context) {
	var list []Project
	for _, p := range projects {
		list = append(list, p)
	}
	c.JSON(http.StatusOK, list)
}

func analyzeProject(c *gin.Context) {
	id := c.Param("id")
	p, ok := projects[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project_id": p.ID,
		"status":     "Analysis Complete",
		"findings":   []string{},
	})
}
