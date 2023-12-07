package usecases

// Mail is a struct that represents a mail information.
type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
	Type    string
}

// SendMail is a use case that will send a mail.
func (c *Core) SendMail(m Mail) error {
	return nil
}
