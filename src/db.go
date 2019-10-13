package main

import (
    "fmt"
    "os"

    "github.com/boltdb/bolt"
)

const dbFilePath   = "../data/links.db"
const dbBucketName = "foobar"

func isLinkSended(link string) (bool, error) {
    err, linkValue := getLinkFromDB(link)

    if (err != nil) {
        return true, err
    }

    if (linkValue != "true" && linkValue != "false") {
        return false, nil
    }

    return true, nil
}

func getLinkFromDB(link string) (error, string) {
    var linkValue []byte

    db, err := bolt.Open(dbFilePath, 0600, &bolt.Options{ReadOnly: true})

    if (err != nil) {
        return err, string(linkValue)
    }

    err = db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(dbBucketName))

        if bucket == nil {
            return fmt.Errorf("Failed To Get DB Bucket")
        }

        linkValue = bucket.Get([]byte(link))

        return nil
    })

    defer db.Close()

    return err, string(linkValue)
}

func saveLinkToDB(link string, status bool) error {
    statusString := "false"
    if (status) {
        statusString = "true"
    }

    db, err := bolt.Open(dbFilePath, 0600, nil)

    if (err != nil) {
        return err
    }

    err = db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists([]byte(dbBucketName))

        if err != nil {
            return fmt.Errorf("Failed To Create DB Bucket: %v", err)
        }

        err = bucket.Put([]byte(link), []byte(statusString))

        if (err != nil) {
            return fmt.Errorf("Failed Insert/Update To Bucket: %v", err)
        }

        return nil
    })

    defer db.Close()

    return err
}

func updLinkInDB(link string, status bool) error {
    return saveLinkToDB(link, status)
}

func initDB() error {
    if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
        return saveLinkToDB("#", true)
    }

    return nil
}
