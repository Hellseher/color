package color

import (
	"fmt"
	"strings"
)

/*************************************************************
 * 16 color Style
 *************************************************************/

// Style a 16 color style. can add: fg color, bg color, color options
//
// Example:
// 	color.Style{color.FgGreen}.Print("message")
type Style []Color

// New create a custom style
//
// Usage:
//	color.New(color.FgGreen).Print("message")
//	equals to:
//	color.Style{color.FgGreen}.Print("message")
func New(colors ...Color) Style {
	return colors
}

// Save to styles map
func (s Style) Save(name string) {
	AddStyle(name, s)
}

// Render render text
// Usage:
//  color.New(color.FgGreen).Render("text")
//  color.New(color.FgGreen, color.BgBlack, color.OpBold).Render("text")
func (s Style) Render(a ...interface{}) string {
	return RenderCode(s.String(), a...)
}

// Renderln render text line.
// like Println, will add spaces for each argument
// Usage:
//  color.New(color.FgGreen).Renderln("text", "more")
//  color.New(color.FgGreen, color.BgBlack, color.OpBold).Render("text", "more")
func (s Style) Renderln(a ...interface{}) string {
	return RenderWithSpaces(s.String(), a...)
}

// Sprint is alias of the 'Render'
func (s Style) Sprint(a ...interface{}) string {
	return RenderCode(s.String(), a...)
}

// Sprintf format and render message.
func (s Style) Sprintf(format string, a ...interface{}) string {
	return RenderString(s.String(), fmt.Sprintf(format, a...))
}

// Print render and Print text
func (s Style) Print(a ...interface{}) {
	doPrint(s.String(), s, fmt.Sprint(a...))
}

// Printf render and print text
func (s Style) Printf(format string, a ...interface{}) {
	doPrint(s.Code(), s, fmt.Sprintf(format, a...))
}

// Println render and print text line
func (s Style) Println(a ...interface{}) {
	doPrintln(s.String(), s, a)
}

// Code convert to code string. returns like "32;45;3"
func (s Style) Code() string {
	return s.String()
}

// String convert to code string. returns like "32;45;3"
func (s Style) String() string {
	return colors2code(s...)
}

// IsEmpty style
func (s Style) IsEmpty() bool {
	return len(s) == 0
}

/*************************************************************
 * Theme(extended Style)
 *************************************************************/

// Theme definition. extends from Style
type Theme struct {
	// Name theme name
	Name string
	// Style for the theme
	Style
}

// NewTheme instance
func NewTheme(name string, style Style) *Theme {
	return &Theme{name, style}
}

// Save to themes map
func (t *Theme) Save() {
	AddTheme(t.Name, t.Style)
}

// Tips use name as title, only apply style for name
func (t *Theme) Tips(format string, a ...interface{}) {
	// only apply style for name
	t.Print(strings.ToUpper(t.Name) + ": ")
	Printf(format+"\n", a...)
}

// Prompt use name as title, and apply style for message
func (t *Theme) Prompt(format string, a ...interface{}) {
	title := strings.ToUpper(t.Name) + ":"
	t.Println(title, fmt.Sprintf(format, a...))
}

// Block like Prompt, but will wrap a empty line
func (t *Theme) Block(format string, a ...interface{}) {
	title := strings.ToUpper(t.Name) + ":\n"

	t.Println(title, fmt.Sprintf(format, a...))
}

/*************************************************************
 * Theme: internal themes
 *************************************************************/

// internal themes(like bootstrap style)
// Usage:
// 	color.Info.Print("message")
// 	color.Info.Printf("a %s message", "test")
// 	color.Warn.Println("message")
// 	color.Error.Println("message")
var (
	// Info color style
	Info = &Theme{"info", Style{OpReset, FgGreen}}
	// Note color style
	Note = &Theme{"note", Style{OpBold, FgLightCyan}}
	// Warn color style
	Warn = &Theme{"warning", Style{OpBold, FgYellow}}
	// Light color style
	Light = &Theme{"light", Style{FgLightWhite, BgBlack}}
	// Error color style
	Error = &Theme{"error", Style{FgLightWhite, BgRed}}
	// Danger color style
	Danger = &Theme{"danger", Style{OpBold, FgRed}}
	// Debug color style
	Debug = &Theme{"debug", Style{OpReset, FgCyan}}
	// Notice color style
	Notice = &Theme{"notice", Style{OpBold, FgCyan}}
	// Comment color style
	Comment = &Theme{"comment", Style{OpReset, FgLightYellow}}
	// Success color style
	Success = &Theme{"success", Style{OpBold, FgGreen}}
	// Primary color style
	Primary = &Theme{"primary", Style{OpReset, FgBlue}}
	// Question color style
	Question = &Theme{"question", Style{OpReset, FgMagenta}}
	// Secondary color style
	Secondary = &Theme{"secondary", Style{FgDarkGray}}
)

// Themes internal defined themes.
// Usage:
// 	color.Themes["info"].Println("message")
var Themes = map[string]*Theme{
	"info":  Info,
	"note":  Note,
	"light": Light,
	"error": Error,

	"debug":   Debug,
	"danger":  Danger,
	"notice":  Notice,
	"success": Success,
	"comment": Comment,
	"primary": Primary,
	"warning": Warn,

	"question":  Question,
	"secondary": Secondary,
}

// AddTheme add a theme and style
func AddTheme(name string, style Style) {
	Themes[name] = NewTheme(name, style)
	Styles[name] = style
}

// GetTheme get defined theme by name
func GetTheme(name string) *Theme {
	return Themes[name]
}

/*************************************************************
 * internal styles
 *************************************************************/

// Styles internal defined styles, like bootstrap styles.
// Usage:
// 	color.Styles["info"].Println("message")
var Styles = map[string]Style{
	"info":  {OpReset, FgGreen},
	"note":  {OpBold, FgLightCyan},
	"light": {FgLightWhite, BgRed},
	"error": {FgLightWhite, BgRed},

	"danger":  {OpBold, FgRed},
	"notice":  {OpBold, FgCyan},
	"success": {OpBold, FgGreen},
	"comment": {OpReset, FgMagenta},
	"primary": {OpReset, FgBlue},
	"warning": {OpBold, FgYellow},

	"question":  {OpReset, FgMagenta},
	"secondary": {FgDarkGray},
}

// some style name alias
var styleAliases = map[string]string{
	"err":  "error",
	"suc":  "success",
	"warn": "warning",
}

// AddStyle add a style
func AddStyle(name string, s Style) {
	Styles[name] = s
}

// GetStyle get defined style by name
func GetStyle(name string) Style {
	if s, ok := Styles[name]; ok {
		return s
	}

	if realName, ok := styleAliases[name]; ok {
		return Styles[realName]
	}

	// empty style
	return New()
}

/*************************************************************
 * quick use style print message
 *************************************************************/

// Infof print message with Info style
func Infof(format string, a ...interface{}) {
	Info.Printf(format, a...)
}

// Infoln print message with Info style
func Infoln(a ...interface{}) {
	Info.Println(a...)
}

// Errorf print message with Error style
func Errorf(format string, a ...interface{}) {
	Error.Printf(format, a...)
}

// Errorln print message with Error style
func Errorln(a ...interface{}) {
	Error.Println(a...)
}

// Warnf print message with Warn style
func Warnf(format string, a ...interface{}) {
	Warn.Printf(format, a...)
}

// Warnln print message with Warn style
func Warnln(a ...interface{}) {
	Warn.Println(a...)
}

