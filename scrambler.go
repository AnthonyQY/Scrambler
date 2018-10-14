package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {

	// Create a ticker object
	instructionTicker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-instructionTicker.C:
			executeInstruction(getInstruction())
		}
	}
}

func executeInstruction(instruction string) {
	switch instruction {

	case "mouse":
		// Move, click, or scroll the mouse
		mouseInstruction()

	case "keyboard":
		// Press a random series of keys
		keyboardInstruction()

	case "sleep":
		// Go to sleep for n seconds
		sleepInstruction()

	case "error":
		// Error; Do nothing

	default:
		// Error; Do nothing

	}
}

func getInstruction() string {
    // Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
	randInstruction := rand.Intn(3)

	switch randInstruction {

	case 0:
		// Mouse event
		return "mouse"
	case 1:
		// Keyboard event
		return "keyboard"
	case 2:
		// Sleep event
		return "sleep"
	default:
		// Error; Do nothing
		return "error"
	}
}

func mouseInstruction() {
    // Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
    
	eventType := rand.Intn(2)

	switch eventType {

	case 0:
		// Move mouse
		mouseMoveEvent()
	case 1:
		// Click mouse
		mouseClickEvent()
	case 2:
		// Scroll mouse
		mouseScrollEvent()
	default:
		// Error; Do nothing
	}
}

func mouseMoveEvent() {
	// Get screen size of monitor
	xScreenSize, yScreenSize := robotgo.GetScreenSize()

    // Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
    
	// Generate random mouse coordinates
	xMousePosition := rand.Intn(xScreenSize)
	yMousePosition := rand.Intn(yScreenSize)

    // Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
    
	moveType := rand.Intn(2)

	switch moveType {

	case 0:
		// Smooth movement
		robotgo.MoveMouseSmooth(xMousePosition, yMousePosition)

	case 1:
		// Instant movement
		robotgo.MoveMouse(xMousePosition, yMousePosition)

	default:
		// Error; Do nothing
	}

}

func mouseClickEvent() {
    // Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
    
	clickType := rand.Intn(2)
	
	// Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
	doubleClick := rand.Intn(2)

	switch clickType {
	case 0:
		// Left click

		// Double click
		if doubleClick == 1 {
			robotgo.MouseClick("left", true)
		} else { // Single Click
			robotgo.MouseClick("left")
		}

	case 1:
		// Right click

		// Double click
		if doubleClick == 1 {
			robotgo.MouseClick("right", true)
		} else { // Single click
			robotgo.MouseClick("right")
		}

	default:
		// Error; Do nothing
	}
}

func mouseScrollEvent() {
    // Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
	scrollDir := rand.Intn(2)
	
    // Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
	scrollMagnitude := rand.Intn(100)

	switch scrollDir {

	case 0:
		// Scroll up
		robotgo.ScrollMouse(scrollMagnitude, "up")

	case 1:
		// Scroll down
		robotgo.ScrollMouse(scrollMagnitude, "down")

	default:
		// Error; Do nothing
	}
}

func keyboardInstruction() {
    // Random seeding using time
    rand.Seed(time.Now().UTC().UnixNano())
	numOfKeys := rand.Intn(100)
	for i := 1; i < numOfKeys; i++ {
		pressRandomKey()
	}
}

func pressRandomKey() {
	characters := [36]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	keyIndex := rand.Intn(36)
	robotgo.KeyTap(characters[keyIndex])
}

func sleepInstruction() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}
