// Licence is the license of the package

// Package first sentance is important in the commnet . This is the description of the package.

/*
My Package documentation
*/
package mypkg

import (
	"fmt"

	"github.com/dragosh/go-learn/_pack-mod_/internal/ports"
)

// Start server documentation
func StartServer() {
	fmt.Println("Call mypkg StartServer() method", ports.ServicePort)
}

func init() {
	fmt.Println("Call mypkg.init() method")
}
