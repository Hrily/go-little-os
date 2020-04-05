package keyboard

type Key struct {
	Code       int32
	Characters [2]byte
	OnPress    func()
	OnRelease  func()
}

type State struct {
	ModifierLevel int32
}

func (s *State) ToggleModifier() {
	s.ModifierLevel = 1 - s.ModifierLevel
}
