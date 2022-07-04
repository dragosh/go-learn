/*
* -----------------------------------------------------------
* Errors and Logging
* -----------------------------------------------------------
 */
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

/*
* -----------------------------------------------------------
* Errors
* The error type is a built-in interface similar to fmt.Stringer:
* More https://go.dev/blog/error-handling-and-go
* Logger https://github.com/sirupsen/logrus
* -----------------------------------------------------------
 */
// Throwing
func retError(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg, nil
}

// Custom Erors
type MyError struct {
	arg     int
	message string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("I really %s %d", err.message, err.arg)
}
func retCustomError(arg int) (int, error) {
	if arg == 42 {
		return -1, &MyError{arg, "can't work with"}
	}
	return arg, nil
}

// It formats a string according to Printf’s rules and returns it as an error created by errors.New.
func retErrorf(arg int) (int, error) {
	if arg == 42 {
		return -1, fmt.Errorf("can't work with %d formated", arg)
	}
	return arg, nil
}

/*
* -----------------------------------------------------------
* Logging
* -----------------------------------------------------------
 */

func logging() {
	// The code above prints the text "Hello world!" to the standard error `> /dev/null`,
	// but it also includes the date and time, which is handy for filtering log messages by date.
	log.Println("Hello world!")
}

// Custom loggers

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// log to file init-ialized
func init() {
	file, err := os.OpenFile("logger.out", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	for _, val := range []int{7, 42} {
		// inline error check on the if line is a common idiom in Go code
		if ret, err := retError(val); err != nil {
			fmt.Println("Woops:", err)
		} else {
			fmt.Println("OK:", ret)
		}
	}
	for _, val := range []int{7, 42} {
		if ret, err := retCustomError(val); err != nil {
			fmt.Println("Woops:", err)
			// fmt.Errorf("Woops %g", err)
		} else {
			fmt.Println("OK:", ret)
		}
	}

	arg, err := retErrorf(42)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arg)
	}

	// If you want to programmatically use the data in a custom error,
	// you’ll need to get the error as an instance of the custom error type
	// via type assertion
	_, e := retCustomError(42)
	if ae, ok := e.(*MyError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	}

	/*
	* -----------------------------------------------------------
	* Logging
	* -----------------------------------------------------------
	 */

	logging()

	// Log to file
	InfoLogger.Println("Starting the application...")
	WarningLogger.Println("There is something you should know about")
	ErrorLogger.Println("Something went wrong")
}
