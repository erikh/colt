package main

import "github.com/fatih/color"

var colorMap = map[string]color.Attribute{
	"black":     color.FgBlack,
	"red":       color.FgRed,
	"green":     color.FgGreen,
	"yellow":    color.FgYellow,
	"blue":      color.FgBlue,
	"magenta":   color.FgMagenta,
	"cyan":      color.FgCyan,
	"white":     color.FgWhite,
	"hiblack":   color.FgHiBlack,
	"hired":     color.FgHiRed,
	"higreen":   color.FgHiGreen,
	"hiyellow":  color.FgHiYellow,
	"hiblue":    color.FgHiBlue,
	"himagenta": color.FgHiMagenta,
	"hicyan":    color.FgHiCyan,
	"hiwhite":   color.FgHiWhite,
}

var bgColorMap = map[string]color.Attribute{
	"black":     color.BgBlack,
	"red":       color.BgRed,
	"green":     color.BgGreen,
	"yellow":    color.BgYellow,
	"blue":      color.BgBlue,
	"magenta":   color.BgMagenta,
	"cyan":      color.BgCyan,
	"white":     color.BgWhite,
	"hiblack":   color.BgHiBlack,
	"hired":     color.BgHiRed,
	"higreen":   color.BgHiGreen,
	"hiyellow":  color.BgHiYellow,
	"hiblue":    color.BgHiBlue,
	"himagenta": color.BgHiMagenta,
	"hicyan":    color.BgHiCyan,
	"hiwhite":   color.BgHiWhite,
}

var attrMap = map[string]color.Attribute{
	"reset":        color.Reset,
	"bold":         color.Bold,
	"faint":        color.Faint,
	"italic":       color.Italic,
	"underline":    color.Underline,
	"blinkslow":    color.BlinkSlow,
	"blinkrapid":   color.BlinkRapid,
	"reversevideo": color.ReverseVideo,
	"concealed":    color.Concealed,
	"crossedout":   color.CrossedOut,
}
