package main

type commandID int

const (
	cmdNick commandID = iota
	cmdJoin
	cmdRooms
	cmdMSG
	cmdQuit
)

type command struct {
	id     commandID
	client *client
	args   []string
}
