package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func InitLogger() {
	logFolder := os.Getenv("LOG_FOLDER")
	appName := os.Getenv("APP_NAME")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error loading .env file")
		return
	}

	logFile := fmt.Sprintf("%s/%s/%s-%s.log", pwd, logFolder, appName, "%Y-%m-%d")
	logFileLink := fmt.Sprintf("%s/%s/%s.log", pwd, logFolder, appName)
	writer, err := rotatelogs.New(
		logFile,
		rotatelogs.WithLinkName(logFileLink),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationCount(10000),
	)
	if err != nil {
		fmt.Println("Failed to initialize log file ", err.Error())
	}
	log.SetOutput(writer)
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Log(msg ...interface{}) {
	msg = append(msg, "\n-----------------------------------------------------------------------------")
	log.Println(msg...)
}
func LogWarning(msg ...interface{}) {
	msgStr := removeBraces(msg)
	fmt.Printf(fmt.Sprintf("%s: WARNING: %v\n", CurrentTime(), msgStr))
	log.Println("WARNING: ", msgStr)
}

func LogError(msg ...interface{}) {
	msgStr := removeBraces(msg)
	fmt.Printf(fmt.Sprintf("%s: ERROR: %v\n", CurrentTime(), msgStr))
	log.Printf(fmt.Sprintf("ERROR: %v\n", msgStr))
}

func removeBraces(msg []interface{}) string {
	msgStr := fmt.Sprintf("%v", msg)
	msgStr = strings.Replace(msgStr, "[", "", 2)
	msgStr = strings.Replace(msgStr, "]", "", 2)
	return msgStr
}
