package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"log"

	tcell "github.com/gdamore/tcell/v2"
	tview "github.com/gord-project/gview"

	"github.com/gord-project/gord/config"
)

type WalColorsStruct struct {
	Wallpaper string `json:"wallpaper"`
	Alpha     string `json:"alpha"`
	Special   struct {
		Background string `json:"background"`
		Foreground string `json:"foreground"`
		Cursor     string `json:"cursor"`
	} `json:"special"`
	Colors struct {
		Color0  string `json:"color0"`
		Color1  string `json:"color1"`
		Color2  string `json:"color2"`
		Color3  string `json:"color3"`
		Color4  string `json:"color4"`
		Color5  string `json:"color5"`
		Color6  string `json:"color6"`
		Color7  string `json:"color7"`
		Color8  string `json:"color8"`
		Color9  string `json:"color9"`
		Color10 string `json:"color10"`
		Color11 string `json:"color11"`
		Color12 string `json:"color12"`
		Color13 string `json:"color13"`
		Color14 string `json:"color14"`
		Color15 string `json:"color15"`
	} `json:"colors"`
}


var walFile string = fmt.Sprintf("%s/.cache/wal/colors.json", os.Getenv("HOME"))

func main() {
	jsonColors, err := ioutil.ReadFile(walFile)
	if err != nil {
		log.Fatal("Cannot open wal config [{}]", walFile)
	}
	
	var wcstruct WalColorsStruct  
	json.Unmarshal([]byte(jsonColors), &wcstruct)

	theme := &config.Theme{
		Theme: &tview.Theme{
			PrimitiveBackgroundColor:    fromHex(wcstruct.Special.Background),
			ContrastBackgroundColor:     fromHex(wcstruct.Special.Foreground),
			MoreContrastBackgroundColor: fromHex(wcstruct.Special.Foreground),
			BorderColor:                 fromHex(wcstruct.Colors.Color0),
			BorderFocusColor:            fromHex(wcstruct.Colors.Color1),
			TitleColor:                  fromHex(wcstruct.Colors.Color13),
			GraphicsColor:               fromHex(wcstruct.Colors.Color2),
			PrimaryTextColor:           fromHex(wcstruct.Special.Foreground),
			SecondaryTextColor:         fromHex(wcstruct.Special.Foreground),
			TertiaryTextColor:          fromHex(wcstruct.Special.Foreground),
			InverseTextColor:            fromHex(wcstruct.Colors.Color0),
			ContrastSecondaryTextColor:  fromHex(wcstruct.Colors.Color0),
		},
		BlockedUserColor: tcell.ColorGray,
		InfoMessageColor: tcell.ColorGray,
		BotColor:         tcell.NewRGBColor(0x94, 0x96, 0xfc),
		MessageTimeColor: fromHex(wcstruct.Colors.Color3),
		LinkColor:        fromHex(wcstruct.Colors.Color4),
		DefaultUserColor: fromHex(wcstruct.Colors.Color5),
		AttentionColor:   fromHex(wcstruct.Colors.Color3),
		ErrorColor:       tcell.ColorRed,
		RandomUserColors: []tcell.Color{
			fromHex(wcstruct.Colors.Color6),
			fromHex(wcstruct.Colors.Color7),
			fromHex(wcstruct.Colors.Color8),
			fromHex(wcstruct.Colors.Color9),
			fromHex(wcstruct.Colors.Color10),
			fromHex(wcstruct.Colors.Color11),
			fromHex(wcstruct.Colors.Color12),
			fromHex(wcstruct.Colors.Color13),
			fromHex(wcstruct.Colors.Color14),
			fromHex(wcstruct.Colors.Color15),
		},
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(theme)
}

// Usage: fromHex("#FF0000")
func fromHex(hexString string) tcell.Color {
	trimmed := strings.TrimPrefix(strings.TrimSpace(hexString), "#")
	var r, g, b int32
	fmt.Sscanf(trimmed, "%02x%02x%02x", &r, &g, &b)
	return tcell.NewRGBColor(r, g, b)
}
