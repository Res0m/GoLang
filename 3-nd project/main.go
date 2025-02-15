package main

import "fmt"



type bookmarkMap = map[string]string // ALIAS, тип данных, который является map[string]string
func add(array []string ){
	array[0] = "2"
}

func main(){
	a := []string{"1"}
	add(a)
	fmt.Println(a)
	


// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// key - значение
// NY - New York
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	//  m := make(bookmarkMap, 3) // make(map[string]string, 3) - создание map с 3 элементами
	//  m["A"] = "1"
	//  m["B"] = "2"
	//  m["C"] = "3"
	//  m["D"] = "4"
	//  fmt.Println(len(m))

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// MAP - это структура данных, которая хранит пары ключ-значение:
    // m := map[string]string { // map [key, например float64]valueб например string{key:value}
	// 	"PurpleSchool": "https://purpleschool.ru", // key + value , key:value
	// } 
	// fmt.Println(m)
	// fmt.Println(m["PurpleSchool"])
	// m["PurpleSchool"] = "sigma"
	// fmt.Println(m["PurpleSchool"])
	// m["Google"] = "google.com"
	// m["Yandex"] = "yandex.ru"
	// fmt.Println(m)
	// delete(m, "Yandex") // delete(map, key), удаляет из map ключ и значение
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -




	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	/* Создать приложение, которое сначала выдаёт меню:
	- 1. Посмотреть закладки
	- 2. Добавить закладку
	- 3. Удалить закладку
	- 4. Выход 
	При 1 - Выводит закладки
	При 2 - 2 поля ввода названия и адреса и после добавления
	При 3 - Ввод названия и удаление по нему
	При 4 - Завершение 
	*/
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	dict := bookmarkMap{}
	for {
		menuNotes()
		var number int
		_, err := fmt.Scan(&number)
		if err != nil{
			fmt.Println("Ошибка, введите число")
			fmt.Scanln()
			continue
		}
		switch number {
        case 1:
            watchNotes(dict)
        case 2:
            addNote(dict)
        case 3:
            deleteNote(dict)
        case 4:
            return
        default:
            fmt.Println("Неверный выбор, попробуйте снова")
        }
	}
}

func watchNotes(test bookmarkMap){	
	for key, value := range test{
		fmt.Println(key, value)
	}
}

func addNote(test bookmarkMap) {
    var name, address string
    var err error
    for {
        fmt.Print("Введите название: ")
        _, err = fmt.Scan(&name)
        if err == nil {
            break
        }
        fmt.Println("Ошибка ввода названия, попробуйте снова", err)
    }
    for {
        fmt.Print("Введите адрес:")
        _, err = fmt.Scan(&address)
        if err == nil {
            break
        }
        fmt.Println("Ошибка ввода адреса, попробуйте снова", err)
    }
    test[name] = address
    fmt.Println("Закладка добавлена", name, address)
}

func deleteNote(test bookmarkMap){
	var deleteNote string
	var err error
	if len(test) == 0 {
		fmt.Println("Нет закладок для удаления.")
		return
	}
	fmt.Print("Введите какую заметку надо удалить: ")
	for{
	_, err = fmt.Scan(&deleteNote)
	if err == nil{
		break
	}
	fmt.Println("Ошибка, ты не прав")
	}
	delete(test, deleteNote )
	fmt.Println("Закладка была удалена")
}

func menuNotes(){
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Завершение")
	fmt.Print("Введите номер действия: ")
}
