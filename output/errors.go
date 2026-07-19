package output

import "github.com/fatih/color"

func PrintError(value any) {
	// fmt.Println(value)
	intValue, ok := value.(int)
	if ok {
		color.Red("Код помилки: %d", intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}
	color.Red("Невідома помилка: %v", value)
	// switch t := value.(type) {
	// case string:
	// 	color.Red(t)
	// case int:
	// 	color.Red("Код помилки: %d", t)
	// case error:
	// 	color.Red(t.Error())
	// default:
	// 	color.Red("Невідома помилка: %v")
	// }
}
