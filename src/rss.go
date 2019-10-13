package main

import "errors"

type RSSItem struct {
    Link string `xml:"link"`
}

type RSSChannel struct {
    Items []RSSItem `xml:"item"`
}

type RSS struct {
    Channel RSSChannel `xml:"channel"`
}

func getRSS(url string) (RSS, error) {
    var rss RSS

    err := getRemoteRSS(url, &rss)

    if (err != nil) {
        return rss, err
    }

    if (!isValidRSSData(rss)) {
        errorMessage := "RSS (" + url + ") Has Invalid Format"
        err           = errors.New(errorMessage)
    }

    return rss, err
}

func isValidRSSData(rss RSS) bool {
    if (len(rss.Channel.Items) < 1) {
        return false
    }

    return true
}