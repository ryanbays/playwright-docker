package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const apiLogFile = "/logs/api.log"
const workerLogFile = "/logs/worker.log"

var logChan = make(chan string, 1000)

func startLogger() {
	// ensure log directory exists
	os.MkdirAll("/logs", os.ModePerm)

	go func() {
		for msg := range logChan {
			f, err := os.OpenFile(apiLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				continue
			}

			f.WriteString(msg + "\n")
			f.Close()
		}
	}()
}

// async logger helper
func logAPI(message string) {
	select {
	case logChan <- message:
	default:
		// drop log if channel is full (prevents blocking API)
	}
}

// Gin middleware
func apiLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		entry := fmt.Sprintf(
			"%s | %s | %s | %d | %s",
			start.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency.String(),
		)

		logAPI(entry)
	}
}

// DELETE /remove/:id
func removeState(c *gin.Context) {
	id := c.Param("id")
	filePath := fmt.Sprintf("/auth/%s.json", id)

	if err := os.Remove(filePath); err != nil {
		logAPI(fmt.Sprintf("Failed to remove file %s: %v", filePath, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove the file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File removed successfully", "id": id})
}

// POST /upload
func uploadState(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logAPI(fmt.Sprintf("Failed to get uploaded file: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	id := uuid.New().String()
	dst := fmt.Sprintf("/auth/%s.json", id)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		logAPI(fmt.Sprintf("Failed to save uploaded file %s: %v", dst, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    file.Filename,
		"id":      id,
	})
}

// GET /status
func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// GET /logs/api
func apiLogs(c *gin.Context) {
	logs, err := os.ReadFile(apiLogFile)
	if err != nil {
		logAPI(fmt.Sprintf("Failed to read log file %s: %v", apiLogFile, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read the log file"})
		return
	}

	c.Data(http.StatusOK, "text/plain; charset=utf-8", logs)
}

// DELETE /logs/api
func deleteApiLogs(c *gin.Context) {
	if err := os.Remove(apiLogFile); err != nil {
		logAPI(fmt.Sprintf("Failed to delete log file %s: %v", apiLogFile, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the log file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Log file deleted successfully"})
}

type LogRequest struct {
	Logs string `json:"logs"`
}

// GET /logs/worker
func workerLogs(c *gin.Context) {
	logs, err := os.ReadFile(workerLogFile)
	if err != nil {
		logAPI(fmt.Sprintf("Failed to read log file %s: %v", workerLogFile, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read the log file"})
		return
	}

	c.Data(http.StatusOK, "text/plain; charset=utf-8", logs)
}

// DELETE /logs/worker
func deleteWorkerLogs(c *gin.Context) {
	if err := os.Remove(workerLogFile); err != nil {
		logAPI(fmt.Sprintf("Failed to delete log file %s: %v", workerLogFile, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the log file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Log file deleted successfully"})
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	startLogger()
	logAPI("API server started")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(apiLogger())

	// core endpoints
	r.GET("/status", status)

	// auth state endpoints
	r.POST("/upload", uploadState)
	r.POST("/remove/:id", removeState)

	// api logs
	r.GET("/logs/api", apiLogs)
	r.DELETE("/logs/api", deleteApiLogs)

	// worker logs
	r.GET("/logs/worker", workerLogs)
	r.DELETE("/logs/worker", deleteWorkerLogs)

	r.Run(":8080")
}

