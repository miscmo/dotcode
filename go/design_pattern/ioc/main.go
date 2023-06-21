package main

import "fmt"

type Widget struct {
	X, Y int
}

type Label struct {
	Widget
	Text string
}

type Button struct {
	Label
}

func NewButton(x, y int, text string) *Button {
	return &Button{Label{Widget{x, y}, text}}
}

type ListBox struct {
	Widget
	Texts []string
	Index int
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

func (label Label) Paint() {
	fmt.Printf("%p: Label.Paint(%q)\n", &label, label.Text)
}

func (button Button) Paint() {
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}
func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}

func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

func main()  {
	//label := Label{Widget{10, 10}, "State:"}
	//
	//label.X = 11
	//label.Y = 12
	//
	//label.Widget.Y = 11
	//
	//fmt.Println(label)
	//
	//button1 := Button{Label{Widget{10, 70}, "OK"}}
	//button2 := NewButton(50, 70, "Cancel")
	//listBox := ListBox{}
	//
	//for _, painter := range []Painter{label, ListBox{}} {
	//
	//}
}
