// Package vcs contains the interface to connect to version control systems
package vcs

// System is the interface containing methods required for all version control systems gh will connect to
type System interface {
	CreateRepository(name string, organization string, private bool) error
}
