/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"example/cmd"
	_ "example/cmd/bar"
	_ "example/cmd/foo"
)

func main() {
	cmd.Execute()
}
