package main

import (
    "os"
    "io/ioutil"
    "encoding/json"
)

type Credentials struct {
    Token     string `json:"token"`
    ChannelID int    `json:"channel_id"`
    AdminID   int    `json:"admin_id"`
}

const credentialsFilePath = "../data/config/credentials.json"

var credentials Credentials

func setCredentials() bool {
    if (isValidCredentials()) {
        return true
    }

    if _, err := os.Stat(credentialsFilePath); os.IsNotExist(err) {
        errorMessage := "Credentials File Is Not Exist"
        handleCredentialsError(errorMessage)
    }

    credentialsFile, err := os.Open(credentialsFilePath)

    if err != nil{
        handleCredentialsError(err.Error())
    }

    defer func() {
        if err := credentialsFile.Close(); err != nil {
            handleCredentialsError(err.Error())
        }
    }()
     
    jsonData, err := ioutil.ReadAll(credentialsFile)

    if err != nil{
        handleCredentialsError(err.Error())
    }

    json.Unmarshal(jsonData, &credentials)

    if (!isValidCredentials()) {
        errorMessage := "Credentials Has Bad Format"
        handleCredentialsError(errorMessage)
    }

    return true
}

func getCredentials() Credentials {
    if (!isValidCredentials()) {
        setCredentials()
    }

    return credentials
}

func isValidCredentials() bool {
    if (len(credentials.Token) < 1) {
        return false
    }

    if (credentials.AdminID == 0) {
        return false
    }

    if (credentials.ChannelID == 0) {
        return false
    }

    return true
}

func handleCredentialsError(errorMessage string) {
    logMessage(errorMessage, "error")
    panic(errorMessage)
}
