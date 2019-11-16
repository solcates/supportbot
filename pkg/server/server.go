package server

type Server interface {
	Serve() (err error)
}
type SupportBot struct {
	host string
}

func NewSupportBot(host string) *SupportBot {
	return &SupportBot{host: host}
}

func (s *SupportBot) Serve() (err error) {
	panic("implement me")
}
