package tools

type GinDemoError struct {
	Msg string
}

func (m *GinDemoError) Error() string {
	return m.Msg
}
