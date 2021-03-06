package result

// import (
// 	"github.com/wunderkraut/radi-api/property"
// )

/**
 * This file holds the definition of Result, which is what an operation
 * returns, and a few usefull base structs that implement Result, which
 * can be used directly or for inheritance
 */

// Result is an what an operation returns
type Result interface {
	// Give a bool channel indicating that the operation is finished
	Finished() chan bool

	// Did the operation execute successfully? Return any error that occured
	Success() bool

	// return any errors that occured.
	Errors() []error

	// Return the Properties for the result
	// Properties() property.Properties
}
