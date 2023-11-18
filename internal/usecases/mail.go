package usecases

import "fmt"

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

func (c *Core) SendMail(m Mail) error {
	fmt.Println(m)
	return nil
}
