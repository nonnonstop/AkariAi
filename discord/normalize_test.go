package discord_test

import (
	"nonnonstop/akariai/discord"
	"testing"
)

func TestNormalizeStringRemoveChar1(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABC あいう 123")
	expected := "ABCあいう123"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringRemoveChar1: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringRemoveChar2(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABC	あいう	123")
	expected := "ABCあいう123"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringRemoveChar2: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringQuestionMark(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABC？あいう？123")
	expected := "ABC?あいう?123"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringQuestionMark: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringFullUpperCase1(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123ＡＢＣ")
	expected := "ABCあいう123ABC"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringFullUpperCase1: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringFullUpperCase2(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123ＸＹＺ")
	expected := "ABCあいう123XYZ"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringFullUpperCase2: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringFullLowerCase1(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123ａｂｃ")
	expected := "ABCあいう123ABC"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringFullLowerCase1: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringFullLowerCase2(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123ｘｙｚ")
	expected := "ABCあいう123XYZ"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringFullLowerCase2: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringHalfLowerCase1(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123abc")
	expected := "ABCあいう123ABC"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringHalfLowerCase1: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringHalfLowerCase2(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123xyz")
	expected := "ABCあいう123XYZ"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringHalfLowerCase2: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringFullKatakana1(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123ア")
	expected := "ABCあいう123あ"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringFullKatakana1: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringFullKatakana2(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123ヶ")
	expected := "ABCあいう123ゖ"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringFullKatakana2: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringHalfKatakana1(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123ｦ")
	expected := "ABCあいう123を"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringHalfKatakana1: %s %s",
			actual, expected)
	}
}

func TestNormalizeStringHalfKatakana2(t *testing.T) {
	t.Parallel()
	actual := discord.ExportNormalizeString("ABCあいう123ﾝ")
	expected := "ABCあいう123ん"
	if actual != expected {
		t.Errorf(
			"Failed TestNormalizeStringHalfKatakana2: %s %s",
			actual, expected)
	}
}
