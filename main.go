package main

import (
	"os"
	"text/template"
)

type Color struct {
	Color_50  string
	Color_100 string
	Color_200 string
	Color_300 string
	Color_400 string
	Color_500 string
	Color_600 string
	Color_700 string
	Color_800 string
	Color_900 string
	Color_950 string
}

type SemanticColors struct {
	Transparent string

	BackgroundDefault string
	BackgroundMuted   string
	BackgroundSurface string
	BackgroundHover   string

	BorderDefault string
	BorderMuted   string

	ForegroundDefault  string
	ForegroundMuted    string
	ForegroundDisabled string

	Yellow            string
	YellowLight       string
	YellowBackground  string
	Orange            string
	OrangeLight       string
	OrangeBackground  string
	Red               string
	RedLight          string
	RedBackground     string
	Magenta           string
	MagentaLight      string
	MagentaBackground string
	Violet            string
	VioletLight       string
	VioletBackground  string
	Blue              string
	BlueLight         string
	BlueBackground    string
	Cyan              string
	CyanLight         string
	CyanBackground    string
	Green             string
	GreenLight        string
	GreenBackground   string
	Gray              string
	GrayLight         string
	GrayBackground    string
}

func main() {
	var yellow = Color{
		Color_50:  "#FDF2E5",
		Color_100: "#FCE5C8",
		Color_200: "#FACA7D",
		Color_300: "#E8B246",
		Color_400: "#C5963A",
		Color_500: "#A17B2E",
		Color_600: "#806122",
		Color_700: "#5F4717",
		Color_800: "#44320E",
		Color_900: "#271C05",
		Color_950: "#181002",
	}

	var orange = Color{
		Color_50:  "#FCEEE7",
		Color_100: "#F9DDCF",
		Color_200: "#F4B68F",
		Color_300: "#E8954E",
		Color_400: "#C57E41",
		Color_500: "#A26734",
		Color_600: "#815127",
		Color_700: "#643E1D",
		Color_800: "#462A11",
		Color_900: "#2A1707",
		Color_950: "#1C0E04",
	}

	var red = Color{
		Color_50:  "#F7E8E8",
		Color_100: "#EFD0D0",
		Color_200: "#E2A0A0",
		Color_300: "#D96B6B",
		Color_400: "#CA4E4E",
		Color_500: "#A94040",
		Color_600: "#873131",
		Color_700: "#692424",
		Color_800: "#4D1818",
		Color_900: "#2F0C0C",
		Color_950: "#200606",
	}

	var magenta = Color{
		Color_50:  "#F5EBEF",
		Color_100: "#EBD8E0",
		Color_200: "#D8B1C2",
		Color_300: "#C788A5",
		Color_400: "#B96B91",
		Color_500: "#9A5878",
		Color_600: "#79445E",
		Color_700: "#5C3347",
		Color_800: "#412231",
		Color_900: "#28131D",
		Color_950: "#1C0C14",
	}

	var violet = Color{
		Color_50:  "#EFEDF3",
		Color_100: "#DCD8E6",
		Color_200: "#BBB1CE",
		Color_300: "#9D8EB8",
		Color_400: "#8876A8",
		Color_500: "#735E96",
		Color_600: "#5C4A77",
		Color_700: "#45375B",
		Color_800: "#322742",
		Color_900: "#1D1629",
		Color_950: "#110C19",
	}

	var blue = Color{
		Color_50:  "#E3ECF6",
		Color_100: "#CADCEE",
		Color_200: "#8CB8DD",
		Color_300: "#6B96B8",
		Color_400: "#5B809E",
		Color_500: "#4A6982",
		Color_600: "#395367",
		Color_700: "#2B4050",
		Color_800: "#1C2C37",
		Color_900: "#101B23",
		Color_950: "#080F15",
	}

	var cyan = Color{
		Color_50:  "#C8F7F8",
		Color_100: "#95EEF1",
		Color_200: "#80CED0",
		Color_300: "#6BADAF",
		Color_400: "#5A9294",
		Color_500: "#4A7A7B",
		Color_600: "#396061",
		Color_700: "#2B4A4A",
		Color_800: "#1B3233",
		Color_900: "#0F1E1F",
		Color_950: "#071313",
	}

	var green = Color{
		Color_50:  "#DBF5DD",
		Color_100: "#AFECB4",
		Color_200: "#94C898",
		Color_300: "#7AA67E",
		Color_400: "#678D6A",
		Color_500: "#557558",
		Color_600: "#435D46",
		Color_700: "#304532",
		Color_800: "#213022",
		Color_900: "#121C13",
		Color_950: "#0B130B",
	}

	var neutral = Color{
		Color_50:  "#F0F1F2",
		Color_100: "#E1E2E6",
		Color_200: "#C4C6CE",
		Color_300: "#A8ABB6",
		Color_400: "#8C90A0",
		Color_500: "#5B5E6D",
		Color_600: "#32343D",
		Color_700: "#282A31",
		Color_800: "#1E1F25",
		Color_900: "#14151A",
		Color_950: "#101115",
	}

	var colors = SemanticColors{
		Transparent:       "#0000",
		BackgroundDefault: neutral.Color_950,
		BackgroundMuted:   neutral.Color_800,
		BackgroundSurface: neutral.Color_700,
		BackgroundHover:   neutral.Color_600,

		BorderDefault: neutral.Color_600,
		BorderMuted:   neutral.Color_700,

		ForegroundDefault:  neutral.Color_200,
		ForegroundMuted:    neutral.Color_400,
		ForegroundDisabled: neutral.Color_500,

		Yellow:            yellow.Color_300,
		YellowLight:       yellow.Color_200,
		YellowBackground:  yellow.Color_800,
		Orange:            orange.Color_300,
		OrangeLight:       orange.Color_200,
		OrangeBackground:  orange.Color_800,
		Red:               red.Color_300,
		RedLight:          red.Color_200,
		RedBackground:     red.Color_800,
		Magenta:           magenta.Color_300,
		MagentaLight:      magenta.Color_200,
		MagentaBackground: magenta.Color_800,
		Violet:            violet.Color_300,
		VioletLight:       violet.Color_200,
		VioletBackground:  violet.Color_800,
		Blue:              blue.Color_300,
		BlueLight:         blue.Color_200,
		BlueBackground:    blue.Color_800,
		Cyan:              cyan.Color_300,
		CyanLight:         cyan.Color_200,
		CyanBackground:    cyan.Color_800,
		Green:             green.Color_300,
		GreenLight:        green.Color_200,
		GreenBackground:   green.Color_800,
	}

	var templateFile = "template.tmpl"
	tmpl, err := template.New(templateFile).ParseFiles(templateFile)

	if err != nil {
		panic(err)
	}

	file, _ := os.Create("themes/retrofit.json")

	defer file.Close()

	err = tmpl.Execute(file, colors)

	if err != nil {
		panic(err)
	}
}
