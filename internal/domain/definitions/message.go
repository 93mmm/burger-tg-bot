package definitions

type Button struct {
	Text string
	Data string
	URL  string
}

type Message struct {
	ChatID  any
	Text    string
	Buttons [][]Button
}
