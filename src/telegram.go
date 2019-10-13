package main

import (
    "strconv"
    "errors"
    "net/url"
)

type TelegramResponse struct {
    Status      bool   `json:"ok"`
    Description string `json:"description"`
}

const TelegramAPIBaseURL = "https://api.telegram.org/bot"

func sendToTelegram(text string, chatID int) error {
    var telegramResponse TelegramResponse

    text = url.QueryEscape(text)

    url := getSendToTelegramURL(text, chatID)
    err := getRemoteTelegram(url, &telegramResponse)

    if (err != nil) {
        return err
    }

    if (len(telegramResponse.Description) < 1) {
        telegramResponse.Description = "Unknown Telegram API Error"
    }

    if (!telegramResponse.Status) {
        return errors.New(telegramResponse.Description)
    }

    return nil
}

func sendToTelegramChannel(text string) error {
    credentials := getCredentials()

    return sendToTelegram(text, credentials.ChannelID)
}

func sendToTelegramAdmin(text string) error {
    credentials := getCredentials()

    return sendToTelegram(text, credentials.AdminID)
}

func getSendToTelegramURL(text string, chatID int) string {
    credentials = getCredentials()

    url := TelegramAPIBaseURL + credentials.Token + "/sendMessage?chat_id="
    url	= url + strconv.Itoa(chatID) + "&text=" + text

    if (len(url) > 2048) {
        url = url[:2048]
    }

    return url
}
