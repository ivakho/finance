package state

type UserState struct {
	Step     string
	TempData map[string]string
}

func New() *UserState {
	return &UserState{
		Step:     "main_menu",
		TempData: make(map[string]string),
	}
}
