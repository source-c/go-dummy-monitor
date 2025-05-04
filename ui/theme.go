package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type customTheme struct {
	isDark bool
	base   fyne.Theme
}

func NewCustomTheme(isDark bool) fyne.Theme {
	var base fyne.Theme
	if isDark {
		base = theme.DefaultTheme()
	} else {
		base = theme.DefaultTheme()
	}
	return &customTheme{isDark: isDark, base: base}
}

func (t *customTheme) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	variant := theme.VariantLight
	if t.isDark {
		variant = theme.VariantDark
	}
	return t.base.Color(n, variant)
}

func (t *customTheme) Font(s fyne.TextStyle) fyne.Resource {
	return t.base.Font(s)
}

func (t *customTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return t.base.Icon(n)
}

func (t *customTheme) Size(n fyne.ThemeSizeName) float32 {
	return t.base.Size(n)
}
