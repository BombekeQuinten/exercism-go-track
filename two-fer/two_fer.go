// Package twofer provides function that returns slogan.
package twofer

import "fmt"

// ShareWith returns slogan with a name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
