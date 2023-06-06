package code

import "fmt"

func Main(app, server string) string {
	return fmt.Sprintf(`package main

import "%s/%s/cmd"

func main() {
	cmd.Execute()
}`, app, server)
}
