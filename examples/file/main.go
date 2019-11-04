
package main

import (
    "os"

    logger "github.com/bogoevskig/golog"
)

func main() {
    f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        panic("error opening log file")
    }
    defer f.Close()

    // configure logger level
    logger.SetLevel("INFO")
    logger.SetOutput(f)

    logger.Debug("debug log")
    logger.Info("info log")
}