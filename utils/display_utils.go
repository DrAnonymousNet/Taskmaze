package utils


import (
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/mattn/go-runewidth"
)


// Constants for column widths

const (
	idWidth       = 4
	dateWidth     = 21
	priorityWidth = 8
)

// CalculateTitleWidth calculates the adjusted title column width based on screen size
func CalculateTitleWidth(screenWidth int) int {
	if screenWidth-32 < 35 {
		return 4 // Just an arbitrary small number to accommodate some characters ellipsis when screen is tool small
	}
	return screenWidth - 45
}

// PrintBorder prints a border string with the specified corner, horizontal, and end characters
func PrintBorder(output io.Writer, corner, horizontal, end string, titleWidth int) {
	border := fmt.Sprintf("%s%s%s%s%s",
		corner,
		strings.Repeat(horizontal, 4),            // Width of " ID "
		strings.Repeat(horizontal, titleWidth+3), // Width of " Title "
		strings.Repeat(horizontal, dateWidth),    // Width of " Deadline "
		end,
	)
	fmt.Fprintln(output, border)
}

// TruncateString wraps a string to the specified width
func TruncateString(input string, maxWidth int) string {
	var result strings.Builder
	for _, runeValue := range input {
		widthOfRune := runewidth.RuneWidth(runeValue)
		maxWidth -= widthOfRune
		if maxWidth < 0 {
			break
		}
		result.WriteRune(runeValue)
	}
	if maxWidth < 0 && result.Len() > 3 {
		result.WriteString("...")
	}
	return result.String()
}

func WrapString(input string, width int) string {
	screenWidth := width
	var result strings.Builder
	for _, runeValue := range input {
		width -= runewidth.RuneWidth(runeValue)
		if width < 0 {
			result.WriteString("\n")
			width =  screenWidth
		}
		result.WriteRune(runeValue)
	}
	return result.String()
}

// GetTerminalSize returns the width and height of the terminal
func GetTerminalSize() (int, int, error) {
	file, err := os.Open("/dev/tty")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	var dimensions [4]uint16
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		file.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&dimensions)),
	)
	if errno != 0 {
		return 0, 0, errno
	}

	return int(dimensions[1]), int(dimensions[0]), nil
}
