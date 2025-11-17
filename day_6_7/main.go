package day_6_7

import (
	"fmt"
)

type chatApp struct {
	users     *userHashMap
	status    *statusArray
	history   *messageHistory
	queue     *messageQueue
	undo      *undoStack
	userCount int
}

func newChatApp() *chatApp {
	return &chatApp{
		users:     newUserHashMap(),
		status:    newStatusArray(2),
		history:   &messageHistory{},
		queue:     &messageQueue{},
		undo:      &undoStack{},
		userCount: 0,
	}
}

func (c *chatApp) register(username, password string) int {
	_, exists := c.users.Get(username)
	if exists {
		fmt.Println("Username already exists!")
		return -1
	}

	id := c.userCount
	c.userCount++

	newUser := user{
		username: username,
		password: password,
	}
	c.users.Put(username, newUser)
	c.status.setStatus(id, "online")

	fmt.Println("User registered:", username, "ID:", id)
	return id
}

func (c *chatApp) login(username, password string) bool {
	u, ok := c.users.Get(username)
	if !ok || u.password != password {
		fmt.Println("Invalid login")
		return false
	}
	fmt.Println(username, "logged in successfully")
	return true
}

func (c *chatApp) sendMessage(from, to, msg string) {
	fullMsg := fmt.Sprintf("%s → %s: %s", from, to, msg)

	c.queue.enqueue(fullMsg)

	c.history.addMessage(fullMsg)

	c.undo.push(fullMsg)

	fmt.Println("Message queued:", fullMsg)
}

func (c *chatApp) receiveMessage() {
	msg, err := c.queue.dequeue()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delivered:", msg)
}

func (c *chatApp) undoLastMessage() {
	msg, err := c.undo.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Undo successful. Removed:", msg)
}

func ChatApp() {
	chat := newChatApp()

	chat.register("alice", "123")
	chat.register("bob", "456")
	chat.login("alice", "123")

	chat.sendMessage("alice", "bob", "Hey Bob!")
	chat.sendMessage("bob", "alice", "Hello Alice!")

	chat.receiveMessage()
	chat.receiveMessage()

	fmt.Println("\nMessage History:")
	chat.history.print()
	fmt.Println("")

	chat.undoLastMessage()
}
