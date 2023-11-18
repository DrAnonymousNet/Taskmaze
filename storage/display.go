package storage

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/DrAnonymousNet/taskmaze/utils"
)



// DisplayManyTasks displays details of multiple tasks in a formatted table
func DisplayManyTasks(tasks []*Task) {
	screenWidth, _, err := utils.GetTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	// Adjusted column widths based on screen size
	titleWidth := utils.CalculateTitleWidth(screenWidth)

	// Create a new tabwriter with padding and a minimum width of 8 characters
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)

	// Generate the top border with adjusted title column width
	utils.PrintBorder(w, "╔", "═", "╗", titleWidth+12) //12 picked from trial and error

	// Generate and print header with adjusted title column width
	header := fmt.Sprintf("║  ID ║ Title%-"+fmt.Sprintf("%d", titleWidth-3)+"s ║ Deadline            ║ Done  ║", "")
	fmt.Fprintln(w, header)

	// Generate the header-bottom border with adjusted title column width
	utils.PrintBorder(w, "╠", "═", "╣", titleWidth+12) //12 picked from trial and error

	// Generate and print tasks with adjusted title column width and title wrapping
	for _, task := range tasks {
		wrappedTitle := utils.TruncateString(task.Title, titleWidth)
		fmt.Fprintf(w, "║ %2d ║ %-"+fmt.Sprintf("%d", titleWidth+3)+"s ║ %s ║ %t ║\n", task.ID, wrappedTitle, task.Deadline.Format("2006-01-02 by 15:04"), task.Done)
	}

	// Generate the bottom border with adjusted title column width
	utils.PrintBorder(w, "╚", "═", "╝", titleWidth+12) //12 picked from trial and error

	// Flush the buffer
	w.Flush()
}

// Display displays the details of a task in a formatted table
func (t *Task) Display() {
	screenWidth, _, err := utils.GetTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	// Adjusted column widths based on screen size
	titleWidth := utils.CalculateTitleWidth(screenWidth)

	// Display borders and headers
	utils.PrintBorder(os.Stdout, "╔", "═", "╗", titleWidth)
	fmt.Printf("║  Task Details: %-"+fmt.Sprintf("%d", titleWidth+11)+"s ║\n", " ")
	utils.PrintBorder(os.Stdout, "╠", "═", "╣", titleWidth)

	// Display task details
	fmt.Printf("║  ID:        %d\n", t.ID)
	wrappedTitle := strings.Split(utils.WrapString(t.Title, titleWidth), "\n")
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
	utils.PrintBorder(os.Stdout, "╚", "═", "╝", titleWidth)
}
