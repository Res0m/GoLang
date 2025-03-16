package geo_test

import (
	"Golang/weather/geo"
	"testing" // Встроенная библиотека в GoLang для тестирования
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка теста, expected result, данные для функции
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата с expected
	if err != nil {
		t.Error(err) // Ошибка теста
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

//go test geo/geo_test.go
