/*
1. What is a Go Module?
	A Go Module is like a container for Go packages.
	Just like npm for nodejs packages, Go modules manage Go packages.
	They ensure you have the right packages and versions needed for your Go project.

2. Purpose of creating modules
	You cannot build your application if you don't make your go application as a module
	You need modules for dependency mangment, Version Control, locking and Updating packages

3. Module Creation:
	In Go, a module is created using the command:
		go mod init <module-name>
	This is similar to initializing a Node.js project with npm init, which creates a package.json file.
	In Go, this command creates a go.mod file, which serves a similar purpose like package.json file for node.

4. Dependency Management:
    In JavaScript, you use npm (or yarn) to manage your packages.
	In Go, the Go toolchain itself handles this.
	There's no separate package manager like npm.
	The go get command is similar to npm install.

5. Lock File:
    npm uses package-lock.json to lock the versions of all packages and their dependencies.
	Similarly, Go has the go.sum file, which ensures the integrity of the modules used.

6. Modules vs Packages:
    In JavaScript, each npm package can be a module.
	In Go, a module is a collection of packages (Go files), and it's bigger than a single package.
	Think of it like a project in JavaScript that contains multiple npm packages.

7. Create and build a Go app
	go mod init <nameOfTheModule>
	go build
	run in the terminal => ./nameOfTheModule
	in windows the user will get .exe file he can click on to run it
	now anyone can run our go program without needing go installed in their system
*/

package main

import "fmt"

// The 'main' function is the entry point of any executable Go program.
// When the program is run, the code inside the 'main' function is executed first.
// This function does not take any arguments and does not return any value.
func main() {
	fmt.Print("Hello world!")
	fmt.Print(`Hello world!`)
}