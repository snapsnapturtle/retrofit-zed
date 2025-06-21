package main

import (
	"os"
	"text/template"
)

type Colors struct {
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
	var colors = Colors{
		Transparent:       "#0000",
		BackgroundDefault: "#101115",
		BackgroundMuted:   "#14151a",
		BackgroundSurface: "#1E1F25",
		BackgroundHover:   "#282A31",

		BorderDefault: "#32343D",
		BorderMuted:   "#282A31",

		ForegroundDefault:  "#A8ABB6",
		ForegroundMuted:    "#5B5E6D",
		ForegroundDisabled: "#32343D",

		Yellow:            "#E8B246",
		YellowLight:       "#FACA7D",
		YellowBackground:  "#44320E",
		Orange:            "#E8954E",
		OrangeLight:       "#F4B68F",
		OrangeBackground:  "#462A11",
		Red:               "#D96B6B",
		RedLight:          "#E2A0A0",
		RedBackground:     "#4D1818",
		Magenta:           "#C788A5",
		MagentaLight:      "#D8B1C2",
		MagentaBackground: "#412231",
		Violet:            "#9D8EB8",
		VioletLight:       "#BBB1CE",
		VioletBackground:  "#322742",
		Blue:              "#6B96B8",
		BlueLight:         "#8CB8DD",
		BlueBackground:    "#1C2C37",
		Cyan:              "#6BADAF",
		CyanLight:         "#80CED0",
		CyanBackground:    "#1B3233",
		Green:             "#7AA67E",
		GreenLight:        "#94C898",
		GreenBackground:   "#213022",
		Gray:              "#A8ABB6",
		GrayLight:         "#C4C6CE",
		GrayBackground:    "#2E3037",
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
