package main

import (
    "os"
    "io/ioutil"
    "encoding/json"
)

type SiteList struct {
    RSSLinks []string `json:"rss"`
}

const siteListFilePath = "../data/config/sites.json"

var siteList SiteList

func setSiteList() bool {
    if (isValidSiteList()) {
        return true
    }

    if _, err := os.Stat(siteListFilePath); os.IsNotExist(err) {
        errorMessage := "Site List File Is Not Exist"
        handleSiteListError(errorMessage)
    }

    siteListFile, err := os.Open(siteListFilePath)

    if err != nil{
        handleSiteListError(err.Error())
    }

    defer func() {
        if err := siteListFile.Close(); err != nil {
            handleSiteListError(err.Error())
        }
    }()
     
    jsonData, err := ioutil.ReadAll(siteListFile)

    if err != nil{
        handleSiteListError(err.Error())
    }

    json.Unmarshal(jsonData, &siteList)

    if (!isValidSiteList()) {
        errorMessage := "Sites List Has Bad Format"
        handleSiteListError(errorMessage)
    }

    return true
}

func getSiteList() SiteList {
    if (!isValidSiteList()) {
        setSiteList()
    }

    return siteList
}

func isValidSiteList() bool {
    if (len(siteList.RSSLinks) < 1) {
        return false
    }

    return true
}

func handleSiteListError(errorMessage string) {
    logError(errorMessage)
    panic(errorMessage)
}
