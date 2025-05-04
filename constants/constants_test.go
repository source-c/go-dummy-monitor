package constants

import (
	"image/color"
	"testing"
)

func TestColorConstants(t *testing.T) {
	// Test alpha values
	if FULL_ALPHA != 255 {
		t.Errorf("Expected FULL_ALPHA to be 255, got %d", FULL_ALPHA)
	}
	if TRANSLUCENT_ALPHA != 180 {
		t.Errorf("Expected TRANSLUCENT_ALPHA to be 180, got %d", TRANSLUCENT_ALPHA)
	}
	if SEMI_TRANSLUCENT_ALPHA != 100 {
		t.Errorf("Expected SEMI_TRANSLUCENT_ALPHA to be 100, got %d", SEMI_TRANSLUCENT_ALPHA)
	}
	if TRANSPARENT_ALPHA != 0 {
		t.Errorf("Expected TRANSPARENT_ALPHA to be 0, got %d", TRANSPARENT_ALPHA)
	}
}

func TestLightColorScheme(t *testing.T) {
	// Test some key colors in the light scheme
	if LightColors.CPU.(color.RGBA).R != LIGHT_GREEN_R {
		t.Errorf("Expected CPU color R to be %d, got %d", LIGHT_GREEN_R, LightColors.CPU.(color.RGBA).R)
	}
	if LightColors.RAM.(color.RGBA).G != LIGHT_PURPLE_G {
		t.Errorf("Expected RAM color G to be %d, got %d", LIGHT_PURPLE_G, LightColors.RAM.(color.RGBA).G)
	}
	if LightColors.DISK.(color.RGBA).B != LIGHT_CYAN_B {
		t.Errorf("Expected DISK color B to be %d, got %d", LIGHT_CYAN_B, LightColors.DISK.(color.RGBA).B)
	}
}

func TestDarkColorScheme(t *testing.T) {
	// Test some key colors in the dark scheme
	if DarkColors.CPU.(color.RGBA).R != DARK_GREEN_R {
		t.Errorf("Expected CPU color R to be %d, got %d", DARK_GREEN_R, DarkColors.CPU.(color.RGBA).R)
	}
	if DarkColors.RAM.(color.RGBA).G != DARK_PURPLE_G {
		t.Errorf("Expected RAM color G to be %d, got %d", DARK_PURPLE_G, DarkColors.RAM.(color.RGBA).G)
	}
	if DarkColors.DISK.(color.RGBA).B != DARK_CYAN_B {
		t.Errorf("Expected DISK color B to be %d, got %d", DARK_CYAN_B, DarkColors.DISK.(color.RGBA).B)
	}
}

func TestEmptyRectangle(t *testing.T) {
	// Test the empty rectangle color
	if EmptyRectangle.A != TRANSPARENT_ALPHA {
		t.Errorf("Expected EmptyRectangle alpha to be %d, got %d", TRANSPARENT_ALPHA, EmptyRectangle.A)
	}
}
