package Utils

import (
	"fmt"
	"log/slog"
)

func HandleErr(err error) {
	slog.Error("error reading incoming message", err)
	fmt.Println("ERROR READING INCOMING TRANSMISSION %w", err)
}
