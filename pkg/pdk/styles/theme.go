package styles

type _style Style

// Theme define a composition of style.
type Theme interface {
	Style

	// AddStyle adds the given Style to the StyleList.
	AddStyle(Style)
	// DelStyle deletes the given Style from the StyleList.
	DelStyle(Style)

	// Styles returns all the Style present in this Theme.
	Styles() []Style
}

// theme is a composition of Style object.
type theme struct {
	_style
	shared []Style
}

// NewTheme return a new Theme object with the given internal Style.
func NewTheme(defaultStyle Style) Theme {
	shared := make([]Style, 0, 8)
	if defaultStyle != nil {
		shared[0] = defaultStyle
	}

	return &theme{
		_style: MakeStyle(),
		shared: shared,
	}
}

func (t *theme) Styles() []Style {
	return t.shared
}

func (t *theme) AddStyle(s Style) {
	panic("implement me")
}

func (t *theme) DelStyle(s Style) {
	panic("implement me")
}
