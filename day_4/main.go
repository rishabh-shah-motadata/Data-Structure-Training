package day4

import "fmt"

type node struct {
	tabName string
	next    *node
}

type browserTabs struct {
	head *node
}

func (b *browserTabs) addTabAtEnd(name string) {
	newNode := &node{tabName: name}
	if b.head == nil {
		b.head = newNode
		return
	}

	curr := b.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (b *browserTabs) addTabAtBeginning(name string) {
	newNode := &node{tabName: name, next: b.head}
	b.head = newNode
}

func (b *browserTabs) addTabAfter(afterTab, newTab string) {
	if b.head == nil {
		fmt.Println("No open tabs!")
		return
	}

	curr := b.head
	for curr != nil && curr.tabName != afterTab {
		curr = curr.next
	}

	if curr == nil {
		fmt.Printf("Tab '%s' not found so cannot add '%s'\n", afterTab, newTab)
		return
	}

	newNode := &node{tabName: newTab, next: curr.next}
	curr.next = newNode
	fmt.Printf("Added '%s' after '%s'\n", newTab, afterTab)
}

func (b *browserTabs) closeTab(name string) {
	if b.head == nil {
		fmt.Println("No open tabs!")
		return
	}

	if b.head.tabName == name {
		temp := b.head
		b.head = b.head.next
		temp.next = nil
		fmt.Printf("Closed tab: %s\n", name)
		return
	}

	curr := b.head
	for curr.next != nil && curr.next.tabName != name {
		curr = curr.next
	}

	if curr.next == nil {
		fmt.Printf("Tab '%s' not found\n", name)
		return
	}

	temp := curr.next
	curr.next = curr.next.next
	temp.next = nil
	fmt.Printf("Closed tab: %s\n", name)
}

func (b *browserTabs) closeTabsToRight(tabName string) {
	if b.head == nil {
		fmt.Println("No open tabs!")
		return
	}

	curr := b.head
	for curr != nil && curr.tabName != tabName {
		curr = curr.next
	}

	if curr == nil {
		fmt.Printf("Tab '%s' not found.\n", tabName)
		return
	}

	right := curr.next
	for right != nil {
		next := right.next
		right.next = nil
		right = next
	}

	curr.next = nil
	fmt.Printf("Closed all tabs to the right of '%s'\n", tabName)
}

func (b *browserTabs) closeTabsToLeft(tabName string) {
	if b.head == nil {
		fmt.Println("No open tabs!")
		return
	}

	if b.head.tabName == tabName {
		fmt.Println("No tabs to the left")
		return
	}

	curr := b.head
	for curr != nil && curr.tabName != tabName {
		curr = curr.next
	}

	if curr == nil {
		fmt.Printf("Tab '%s' not found\n", tabName)
		return
	}

	temp := b.head
	for temp != curr {
		next := temp.next
		temp.next = nil
		temp = next
	}
	b.head = curr
	fmt.Printf("Closed all tabs to the left of '%s'\n", tabName)
}

func (b *browserTabs) closeAllTabs() {
	curr := b.head
	for curr != nil {
		next := curr.next
		curr.next = nil
		curr = next
	}
	b.head = nil
	fmt.Println("All tabs closed")
}

func (b *browserTabs) displayTabs() {
	if b.head == nil {
		fmt.Println("No open tabs!")
		return
	}

	fmt.Println("Open Tabs:")
	curr := b.head
	for curr != nil {
		fmt.Println("-", curr.tabName)
		curr = curr.next
	}
	fmt.Println()
}

func BrowserTabsDemo() {
	tabs := &browserTabs{}

	tabs.addTabAtEnd("YouTube")
	tabs.addTabAtEnd("GitHub")
	tabs.addTabAtBeginning("Gmail")
	tabs.addTabAtEnd("ChatGPT")
	tabs.displayTabs()

	tabs.addTabAfter("GitHub", "Reddit")
	tabs.displayTabs()

	tabs.closeTabsToRight("Reddit")
	tabs.displayTabs()

	tabs.closeTab("GitHub")
	tabs.displayTabs()

	tabs.addTabAtEnd("Twitter")
	tabs.addTabAtEnd("StackOverflow")
	tabs.addTabAtBeginning("LinkedIn")
	tabs.displayTabs()

	tabs.closeTabsToLeft("Reddit")
	tabs.displayTabs()

	tabs.closeAllTabs()
	tabs.displayTabs()
}
