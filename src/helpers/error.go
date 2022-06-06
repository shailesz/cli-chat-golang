package helpers

import "log"

// handleError is helper function to handle errors.
func handleError(err error) {
	log.Panicln(err)
}
