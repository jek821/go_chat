package Utils

import (
	"fmt"
	"log/slog"
)

func HandleErr(err error) {
	slog.Error("error occurred", "error", err)
	fmt.Printf("ERROR: %v\n", err)
}
