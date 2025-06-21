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
		YellowBackground:  "#4A3A1E",
		Orange:            "#E8954E",
		OrangeBackground:  "#473024",
		Red:               "#D96B6B",
		RedBackground:     "#432626",
		Magenta:           "#C788A5",
		MagentaBackground: "#3F2D34",
		Violet:            "#9D8EB8",
		VioletBackground:  "#322E3C",
		Blue:              "#6B96B8",
		BlueBackground:    "#243039",
		Cyan:              "#6BADAF",
		CyanBackground:    "#243837",
		Green:             "#7AA67E",
		GreenBackground:   "#263328",
		Gray:              "#8D909F",
		GrayBackground:    "#282A31",
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
