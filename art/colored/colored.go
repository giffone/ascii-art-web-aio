// package colored contains all functions related to coloring
package colored

import (
	"ascii-art-web/art/check"
	"ascii-art-web/art/convert"
	"ascii-art-web/logging"
	"ascii-art-web/models"
	"strings"
)

// Lib consists color libruary
func Lib(name string) string {
	switch name {
	case "black":
		return "<code style=\"color: black;\">"
	case "red":
		return "<code style=\"color: red;\">"
	case "green":
		return "<code style=\"color: green;\">"
	case "yellow":
		return "<code style=\"color: yellow;\">"
	case "blue":
		return "<code style=\"color: blue;\">"
	case "purple":
		return "<code style=\"color: purple;\">"
	case "cyan":
		return "<code style=\"color: cyan;\">"
	case "white":
		return "<code style=\"color: white;\">"
	case "orange":
		return "<code style=\"color: orange;\">"
	default:
		return "</code>"
	}
}

func Find(a *models.Art) {
	a.Color.MethodColoring = "none"

	if a.Page.ColorCurr != "" {
		if a.Page.ColorParam != "" {
			if strings.HasPrefix(a.Page.ColorParam, "{") {
				// if founded parameters for coloring
				Param(a, "{}")
			} else if strings.HasPrefix(a.Page.ColorParam, "[") {
				// if founded parameters for coloring
				Param(a, "[]")
			}
		} else {
			// if no parameters - color all words
			a.Color.MethodColoring = "all"
		}
	}
}

func Param(a *models.Art, bracket string) {
	value := a.Page.ColorParam
	lValue := len(value)
	if check.Brackets(value, lValue, "{}", "[]") {
		arrParam, param, addLast := []string{}, "", false
		if lValue == 2 {
			// if parameter like [],{} , color whole text
			a.Color.MethodColoring = "all"
		} else if lValue == 3 && check.ByteConsist(value[1], ':', '+', '-', '.', ',', ';') {
			// if parameter like (:) or [+] or (-)
			a.Color.MethodColoring = "all"
		} else {
			if check.ByteConsist(value[1], ':', '+', '-') {
				// if parameter like (:d...) or [:5...]
				if bracket == "{}" {
					// add first symbol of accepted ascii
					arrParam = append(arrParam, " ")
					a.Color.MethodColoring = "and"
				} else {
					// add first "0" index
					arrParam = append(arrParam, "0")
					a.Color.MethodColoring = "and"
				}
			} else if check.ByteConsist(value[lValue-2], ':', '+', '-') {
				// if parameter like(...d:) or [...5:]
				addLast = true
			}
			for _, elem := range value {
				if elem == '}' || elem == ']' {
					if param != "" {
						arrParam = append(arrParam, param)
					}
					break
				}
				if elem == '{' || elem == '[' {
					continue
				}
				if elem == ':' || elem == '-' || elem == '+' || elem == '.' || elem == ',' || elem == ';' {
					if param != "" {
						arrParam = append(arrParam, param)
					}
					param = ""
					if elem == ':' || elem == '-' || elem == '+' {
						a.Color.MethodColoring = "and"
						continue
					}
					a.Color.MethodColoring = "or"
					continue
				}
				if elem != 0 {
					param += string(elem)
				}
			}
			if addLast {
				if bracket == "{}" {
					// add last symbol of accepted ascii
					arrParam = append(arrParam, "~")
				} else {
					// count all letters in all arguments to make max possible range index to add Color.ByIndex.Range2
					a.Color.ByIndex.MaxIndex = true
				}
			}
			// save range of letters and delete from args
			if bracket == "{}" {
				a.Color.BySymbol.Range = arrParam
				a.Color.MethodBy = "bySymbol"
			} else {
				var err error
				count := 0
				a.Color.ByIndex.Range, count, err = convert.ArrAtoi(arrParam)
				if err != nil {
					logging.WarningInd(a, count)
				}
				a.Color.MethodBy = "byIndex"
			}
			// correct method to "or" if parameters less than 2
			if len(arrParam) < 2 {
				a.Color.MethodColoring = "or"
			}
		}
	} else {
		logging.WarningBrc(a, value)
	}
}

func ChooseColor(a *models.Art) {
	// start color
	a.Color.Case1 = Lib(strings.ToLower(a.Page.ColorCurr))
	// stop color
	a.Color.Case2 = Lib("")
}

func AddStartColor(a *models.Art, b *models.Buf) {
	for i := 0; i < len(b.ArrStr1); i++ {
		b.ArrStr1[i] += a.Color.Case1
	}
}
