package modified

type State bool

func NewState() *State {
	return new(State)
}

func (s *State) Set(modified bool) {
	*s = State(modified)
}

func (s *State) Get() bool {
	return bool(*s)
}
