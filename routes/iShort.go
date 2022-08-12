package routes

type iShort interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	getSize() int
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) GetLogo() string {
	return s.logo
}

func (s *short) setSize(size int) {
	s.size = size
}

func (s *short) getSize() int {
	return s.size
}
