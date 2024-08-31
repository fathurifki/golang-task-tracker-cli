package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Activity struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

func NewActivity(id int, description string) Activity {
	now := time.Now()
	return Activity{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func ListActivities(activities map[int]Activity, listType string) {
	list := make(map[string][]Activity)
	fmt.Println("List Activities:")

	for _, activity := range activities {
		if activity.Status == listType || listType == "all" {
			list[activity.Status] = append(list[activity.Status], activity)
			fmt.Printf("ID: %d, Description: '%s', Status: %s, Created: %s\n",
				activity.ID, activity.Description, activity.Status, activity.CreatedAt.Format(time.RFC822))
		}
	}

	if len(list) == 0 {
		fmt.Printf("There are no activities with status '%s'\n", listType)
	}
}

func UpdateActivity(id int, activity map[int]Activity, field string, value string) {
	act, err := activity[id]

	if !err {
		fmt.Printf("Activity with ID %d not found.\n", id)
		return
	}

	switch field {
	case "desc":
		act.Description = value
		fmt.Printf("Updated activity %d description to %s\n", id, value)
	case "status":
		act.Status = value
		fmt.Printf("Updated activity %d status to %s\n", id, value)
	}

	act.UpdatedAt = time.Now()
	activity[id] = act
}

func main() {
	activities := make(map[int]Activity)
	inputID := 1

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Activity Tracker")
	fmt.Println("----------------")
	fmt.Println("Commands: add, list, update, delete, quit")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "quit", "q":
			return
		case "add":
			description := parts[1]

			activity := NewActivity(inputID, description)
			activities[activity.ID] = activity
			inputID++

			fmt.Printf("Added activity: ID=%d, Description='%s', Status=%s\n",
				activity.ID, activity.Description, activity.Status)

		case "update":
			if len(parts) < 2 {
				fmt.Println("Invalid update command. Usage: update <id> status <new_status>")
				fmt.Println(parts, len(parts))
				continue
			}

			updatedParts := strings.Fields(parts[1])
			update_id, err := strconv.Atoi(updatedParts[0])
			if err != nil {
				fmt.Println("Invalid ID. Please provide a valid integer ID.")
				continue
			}

			update_command := updatedParts[1]
			value := strings.Trim(updatedParts[2], "'\"")

			_, found := activities[update_id]
			if !found {
				fmt.Printf("Activity with ID %d not found.\n", update_id)
				continue
			}

			UpdateActivity(update_id, activities, update_command, value)

		case "list":
			if len(parts) < 2 {
				fmt.Println("Usage: list <all|todo|in-progress|done>")
				continue
			}

			listType := parts[1]
			ListActivities(activities, listType)

		case "delete":
			updatedParts := strings.Fields(parts[1])
			deleteId, err := strconv.Atoi(updatedParts[0])
			if err != nil {
				fmt.Printf("Activity with ID %d not found.\n", deleteId)
				continue
			}

			delete(activities, deleteId)
			fmt.Printf("Activity Id %d has been delete \n", deleteId)

		default:
			fmt.Println("Unknown command. Available commands: add, quit")
		}
	}

}
