package main

import (
	"github.com/DrAnonymousNet/taskmaze/storage"
	"github.com/DrAnonymousNet/taskmaze/cmd"
)

func main() {
	storage.InitDB()
	cmd.RootCmd.Execute()

}
