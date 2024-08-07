package server

func NewServer(s Server, urlListener string) Server {
	s.SetUrlListener(urlListener)
	return s
}
