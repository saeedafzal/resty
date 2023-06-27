package util

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// HexToColour is a helper function to convert hex colour strings to 32-bit integers recognised by [tcell].
func HexToColour(hex string) tcell.Color {
	if strings.HasPrefix(hex, "#") {
		hex = hex[1:]
	}

	c, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return tcell.ColorDefault
	}

	return tcell.NewHexColor(int32(c))
}
