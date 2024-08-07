package server

func NewServer(s Server, urlListener string) Server {
	s.SetURLListener(urlListener)
	return s
}
