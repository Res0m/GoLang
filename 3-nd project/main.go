package main

import "fmt"

func main(){
// key - значение
// NY - New York

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
	dict := map[string]string{}
	for {
		menuNotes()
		var number float64
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

func watchNotes(test map[string]string){	
	for key, value := range test{
		fmt.Println(key, value)
	}
}

func addNote(test map[string]string) {
    var name, address string
    var err error
    for {
        fmt.Println("Введите название:")
        _, err = fmt.Scan(&name)
        if err == nil {
            break
        }
        fmt.Println("Ошибка ввода названия, попробуйте снова:", err)
    }
    for {
        fmt.Println("Введите адрес:")
        _, err = fmt.Scan(&address)
        if err == nil {
            break
        }
        fmt.Println("Ошибка ввода адреса, попробуйте снова:", err)
    }
    test[name] = address
    fmt.Println("Закладка добавлена:", name, address)
}

func deleteNote(test map[string]string){
	var deleteNote string
	var err error
	for{
	fmt.Println("Введите какую заметку надо удалить: ")
	_, err = fmt.Scan(&deleteNote)
	if err == nil{
		break
	}else{
		fmt.Println("Ошибка, ты не прав")
	}
	delete(test, deleteNote )
	fmt.Println("Заметка была удалена")
	}
}

func menuNotes(){
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Завершение")
	fmt.Print("Введите номер действия: ")
}
