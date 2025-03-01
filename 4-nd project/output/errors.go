package output

import "github.com/fatih/color"

func PrintError(value any) { // any -> любой тип // interface{} -> тоже самое что и any, сам any это алиас для interface{}
	// 2 Способа обработки типа any, проверка на type, 1-ый через some, flag := any.(указываем тип)
	// ... 2-ой через switch -> switch some := any.(type) -> дальше делаем условия для типов, например case int:

	// The first approach
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		color.White(strValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		color.Black(errorValue.Error())
		return
	}
	color.Red("Неизвестный тип ошибки")

	// The second approach
	// switch t := value.(type){
	// case string:
	// 	color.White(t)
	// case int:
	// 	color.Red("Код ошибки: %d", t)
	// case error:
	// 	color.Black(t.Error())
	// default:
	// 	color.Red("Неизвестный тип ошибки")
	// }
}

func sum[T int | int16 | int32 | float32 | float64](a, b T) T{ // Generic тип -> делает мультитип, облегчает работа для функций которые работают с несколькими типами
	return a + b											   // Вместо T, может быть любое название, также Generic типов можно быть несколько, записываются через запятую
}


