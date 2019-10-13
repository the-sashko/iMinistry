package main

import (
    "os"
    "strconv"
)

const dataDirPath = "../data"

func main() {
    var RSSPosts []RSSItem

    SentLinksCount := 0

    prepareData()

    logMessage("Start Parsing", "default")
    siteList := getSiteList()

    for _, RSSLink := range siteList.RSSLinks {
        rss, err := getRSS(RSSLink)

        if (err != nil) {
            logError(err.Error())
        }

        RSSPosts = append(RSSPosts, rss.Channel.Items...)
    }

    for _, RSSPost := range RSSPosts {
        isLinkSended, err := isLinkSended(RSSPost.Link)

        if (err != nil) {
            logError(err.Error())
            continue
        }

        if (isLinkSended) {
            continue
        }

        err = saveLinkToDB(RSSPost.Link, false)

        if (err != nil) {
            logError(err.Error())
            continue
        }

        err = sendToTelegramChannel(RSSPost.Link)
        if (err != nil) {
            logError(err.Error())
            continue
        }

        err = updLinkInDB(RSSPost.Link, true)
        if (err != nil) {
            logError(err.Error())
        }

        SentLinksCount++
    }

    message := strconv.Itoa(SentLinksCount) + " Links Sent To Telegram"
    logMessage(message, "default")
    logMessage("End Parsing", "default")
}

func prepareData() {
    if _, err := os.Stat(dataDirPath); os.IsNotExist(err) {
        os.Mkdir(dataDirPath, os.ModePerm)
    }

    if _, err := os.Stat(dataDirPath + "/logs"); os.IsNotExist(err) {
        os.Mkdir(dataDirPath + "/logs", os.ModePerm)
    }

    setCredentials()

    err := initDB()

    if (err != nil) {
        logError(err.Error())
        panic(err.Error())
    }
}
