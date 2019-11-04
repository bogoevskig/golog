package main

import (
    logger "github.com/bogoevskig/golog"
)

func main() {
    // configure logger level
    logger.SetLevel("INFO")

    logger.Debug("debug log")
    logger.Info("info log")
}
