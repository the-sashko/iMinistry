package main

import (
    "net/http"
    "encoding/xml"
    "encoding/json"
)

func getFromRemote(url string) (*http.Response, error) {
    HTTPResponse, err := http.Get(url)

    if (err != nil) {
        return nil, err
    }

    return HTTPResponse, nil
}

func getRemoteRSS(url string, rss *RSS) error {
    HTTPResponse, err := getFromRemote(url)

    if (err != nil) {
        return err
    }

    xmlDecoder := xml.NewDecoder(HTTPResponse.Body)
    err        = xmlDecoder.Decode(&rss)

    defer HTTPResponse.Body.Close()

    if (err != nil) {
        return err
    }

    return nil
}

func getRemoteTelegram(url string, telegramResponse *TelegramResponse) error {
    HTTPResponse, err := getFromRemote(url)

    if (err != nil) {
        return err
    }

    jsonDecoder := json.NewDecoder(HTTPResponse.Body)
    err          = jsonDecoder.Decode(&telegramResponse)

    defer HTTPResponse.Body.Close()

    return err
}
