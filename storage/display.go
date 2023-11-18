package storage

import (
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"text/tabwriter"
	"unsafe"

	"github.com/mattn/go-runewidth"
)

// Constants for column widths
const (
	idWidth       = 4
	dateWidth     = 21
	priorityWidth = 8
)

// DisplayManyTasks displays details of multiple tasks in a formatted table
func DisplayManyTasks(tasks []*Task) {
	screenWidth, _, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	// Adjusted column widths based on screen size
	titleWidth := calculateTitleWidth(screenWidth)

	// Create a new tabwriter with padding and a minimum width of 8 characters
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)

	// Generate the top border with adjusted title column width
	printBorder(w, "╔", "═", "╗", titleWidth+12) //12 picked from trial and error

	// Generate and print header with adjusted title column width
	header := fmt.Sprintf("║  ID ║ Title%-"+fmt.Sprintf("%d", titleWidth-3)+"s ║ Deadline            ║ Done  ║", "")
	fmt.Fprintln(w, header)

	// Generate the header-bottom border with adjusted title column width
	printBorder(w, "╠", "═", "╣", titleWidth+12) //12 picked from trial and error

	// Generate and print tasks with adjusted title column width and title wrapping
	for _, task := range tasks {
		wrappedTitle := truncateString(task.Title, titleWidth)
		fmt.Fprintf(w, "║ %2d ║ %-"+fmt.Sprintf("%d", titleWidth+3)+"s ║ %s ║ %t ║\n", task.ID, wrappedTitle, task.Deadline.Format("2006-01-02 by 15:04"), task.Done)
	}

	// Generate the bottom border with adjusted title column width
	printBorder(w, "╚", "═", "╝", titleWidth+12) //12 picked from trial and error

	// Flush the buffer
	w.Flush()
}

// Display displays the details of a task in a formatted table
func (t *Task) Display() {
	screenWidth, _, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	// Adjusted column widths based on screen size
	titleWidth := calculateTitleWidth(screenWidth)

	// Display borders and headers
	printBorder(os.Stdout, "╔", "═", "╗", titleWidth)
	fmt.Printf("║  Task Details: %-"+fmt.Sprintf("%d", titleWidth+11)+"s ║\n", " ")
	printBorder(os.Stdout, "╠", "═", "╣", titleWidth)

	// Display task details
	fmt.Printf("║  ID:        %d\n", t.ID)
	wrappedTitle := strings.Split(wrapString(t.Title, titleWidth), "\n")
	fmt.Printf("║  Title:     %s\n", wrappedTitle[0])
	if len(wrappedTitle) > 1 {
		for _, title := range wrappedTitle[1:] {
			fmt.Printf("║             %s\n", title)
		}
	}
	fmt.Printf("║  Due Date:  %s\n", t.Deadline.Format("2006-01-02 by 15:04"))
	fmt.Printf("║  Priority:  %s\n", t.Priority)
	fmt.Printf("║  Remind Me: %s\n", t.RemindMe.Format("2006-01-02 by 15:04"))
	fmt.Printf("║  Completed: %t\n", t.Done)

	// Display bottom border
	printBorder(os.Stdout, "╚", "═", "╝", titleWidth)
}

// calculateTitleWidth calculates the adjusted title column width based on screen size
func calculateTitleWidth(screenWidth int) int {
	if screenWidth-32 < 35 {
		return 4 // Just an arbitrary small number to accommodate
	}
	return screenWidth - 45
}

// printBorder prints a border string with the specified corner, horizontal, and end characters
func printBorder(output io.Writer, corner, horizontal, end string, titleWidth int) {
	border := fmt.Sprintf("%s%s%s%s%s",
		corner,
		strings.Repeat(horizontal, 4),            // Width of " ID "
		strings.Repeat(horizontal, titleWidth+3), // Width of " Title "
		strings.Repeat(horizontal, dateWidth),    // Width of " Deadline "
		end,
	)
	fmt.Fprintln(output, border)
}

// truncateString wraps a string to the specified width
func truncateString(input string, maxWidth int) string {
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

func wrapString(input string, width int) string {
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

// getTerminalSize returns the width and height of the terminal
func getTerminalSize() (int, int, error) {
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
