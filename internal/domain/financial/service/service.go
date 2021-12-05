package service

type Service struct {
	Message     string
	Destination string
}

func (s *Service) Dispatch() string {

	return s.Message + s.Destination
}
