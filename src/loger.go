package main

import (
    "os"
    "fmt"
    "time"
    "strconv"
)

const logDirPath = "../data/logs/"

func logMessage(text string, logType string) {
    logFilePath    := getLogFilePath(logType)
    oldLogFilePath := getOldLogFilePath(logType)

    currTime := time.Now().Format("[2006-01-02 15:04:05] ")
    text      = currTime + text + "\n"

    fmt.Print(text)

    if _, err := os.Stat(logDirPath + logType); os.IsNotExist(err) {
        os.Mkdir(logDirPath + logType, os.ModePerm)
    }

    if _, err := os.Stat(oldLogFilePath); ! os.IsNotExist(err) {
        os.Remove(oldLogFilePath)
    }

    if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
        os.Create(logFilePath)
    }

    logFile, err := os.OpenFile(logFilePath, os.O_WRONLY | os.O_APPEND, os.ModePerm)

    if err != nil {
        panic(err)
    }

    defer func() {
        if err := logFile.Close(); err != nil {
            panic(err)
        }
    }()

    if _, err := logFile.WriteString(text); err != nil {
        panic(err)
    }
}

func logError(text string) {
    logMessage(text, "error")

    err := sendToTelegramAdmin(text)

    if (err != nil) {
        logMessage(err.Error(), "error")
        panic(err.Error())
    }
}

func getLogFilePath(logType string) string {
    currDate := time.Now().Format("2006-01-02")

    return logDirPath + logType + "/" + currDate + ".log"
}

func getOldLogFilePath(logType string) string {
    currYear := time.Now().Format("2006")

    prevYear, _ := strconv.Atoi(currYear)
    prevYear--

    prevYearStr := strconv.Itoa(prevYear)

    oldDate := prevYearStr + time.Now().Format("-01-02")

    return logDirPath + logType + "/" + oldDate + ".log"
}
