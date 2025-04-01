package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)


var logFile *os.File

func init() {
	moduleRoot, err := findModuleRoot()
	if err != nil {
		fmt.Printf("Error finding module root: %v\n", err)
		os.Exit(1)
	}

	logFilePath := filepath.Join(moduleRoot, "pkg", "log", "app.log")

	fmt.Printf("Log file path: %s\n", logFilePath) 

	logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		os.Exit(1)
	}
}

func findModuleRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); !os.IsNotExist(err) {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}
	return "", fmt.Errorf("go.mod not found")
}

func logTime() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
}

func Info(message string) {
	logMessage("INFO", message, nil, nil)
}

func Infof(format string, args ...interface{}) {
    logMessage("INFO", fmt.Sprintf(format, args...), nil, nil)
}

func Error(message string, err error) {
	logMessage("ERROR", message, err, nil)
}

func JWTInfo(message string, c *fiber.Ctx) {
	logMessage("JWT_INFO", message, nil, c)
}

func JWTError(message string, err error, c *fiber.Ctx) {
	logMessage("JWT_ERROR", message, err, c)
}

func logMessage(level, message string, err error, c *fiber.Ctx) {
	timestamp := logTime()
	errorMsg := ""
	if err != nil {
		errorMsg = " - Error: " + err.Error()
	}

	ip := ""
	userAgent := ""
	if c != nil {
		ip = getIP(c)
		userAgent = c.Get("User-Agent")
	}

	logEntry := fmt.Sprintf(
		"%s [%s] - IP: %s - User-Agent: %s - Message: %s%s\n",
		timestamp,
		level,
		ip,
		userAgent,
		message,
		errorMsg,
	)

	if _, writeErr := logFile.WriteString(logEntry); writeErr != nil {
		fmt.Printf("Error writing to log file: %v\n", writeErr)
	}
}

func getIP(c *fiber.Ctx) string {
	if ip := c.Get("X-Forwarded-For"); ip != "" {
		return ip
	}

	return c.IP()
}

func Close() {
	if err := logFile.Close(); err != nil {
		fmt.Printf("Error closing log file: %v\n", err)
	}
}
