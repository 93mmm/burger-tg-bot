package definitions

type Button struct {
	Text string
	Data string
	URL  string
}

type Message struct {
	GifFileID      *int
	ReplyMessageID int
	ChatID         any
	Text           string

	Buttons [][]Button
}
