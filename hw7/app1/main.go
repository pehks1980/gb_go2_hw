package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)
// функция распечатки структуры из методы
func PrintStruct(in interface{}) {
	// проверка что не пустой тип
	if in == nil {
		return
	}

	val := reflect.ValueOf(in)
	// если указатель то разымененвываем его
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// если не структура то возврат
	if val.Kind() != reflect.Struct {
		return
	}
	// число полей структуры (рутовых)
	max_field := val.NumField()

	for i := 0; i < max_field; i++ {
		// берем очередное поле структуры
		typeField := val.Type().Field(i)
		// если вложенная структура
		if typeField.Type.Kind() == reflect.Struct {
			log.Printf("nested field: %v", typeField.Name)
			// заходим во вложенную структуру рекурсивно и тд
			nested_struct_interface := val.Field(i).Interface()
			PrintStruct(nested_struct_interface)
			// выходим из рекурсии (из подкаталога) идем к след. полю
			continue
		}

		log.Printf("\tname=%s, type=%s, value=%v, tag=`%s`\n",
			typeField.Name,
			typeField.Type,
			val.Field(i),
			typeField.Tag,
		)
	}
}
// функция приведения типа к строковому
func ToString(any interface{}) string {
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return any.(string)
	default:
		return "???"
	}
}
// функция приведения типа к Integer
func ToInt(any interface{}) int {
	switch v := any.(type) {
	case int:
		return any.(int)
	case string:
		i, _ := strconv.Atoi(v)
		return i
	default:
		return 0
	}
}
/*
Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и values map[string]interface{}
(key - название поля структуры, которому нужно присвоить value этой мапы).
Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
Функция может возвращать только ошибку error. Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).

*/
func functionOne(in interface{}, values map[string]interface{}) {
	// проверка что не пустой тип
	if in == nil {
		return
	}

	val := reflect.ValueOf(in)
	// если указатель то разымененвываем его
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// если не структура то возврат
	if val.Kind() != reflect.Struct {
		return
	}
	//val структура
	// число полей структуры (рутовых)
	max_field := val.NumField()

	for i := 0; i < max_field; i++ {
		// берем очередное поле структуры val
		typeField := val.Type().Field(i)

		for key, v := range values {
			if typeField.Name == key {
				takenElem := val.Field(i)
				if takenElem.CanSet(){
					// если элемент можно перезадать перезадаем его в зависимости от типчика
					vInterface := takenElem.Interface()
					switch vInterface.(type) {
					case int:
						fmt.Printf("\n****altered int field %s = %d with new val %d\n",typeField.Name,takenElem,ToInt(v))
						takenElem.SetInt(int64(ToInt(v)))
					case string:
						fmt.Printf("\n****altered string field %s = %s with new val %s\n",typeField.Name,takenElem,ToString(v))
						takenElem.SetString(ToString(v))
					default:
						fmt.Printf("couldnt find appropitate type to set")
					}

				} else {
					fmt.Printf("\n****field %s = %s cannot be set\n",typeField.Name,val.Field(i))
				}

			}
		}

		// если вложенная структура
		if typeField.Type.Kind() == reflect.Struct {
			//log.Printf("nested field: %v", typeField.Name)
			// заходим во вложенную структуру рекурсивно и передаем адресс чтобы можно было менять данные субструктуры
			nested_struct_iface := val.Field(i).Addr().Interface()
			functionOne(nested_struct_iface, values)
			// выходим из рекурсии (из подкаталога) идем к след. полю
			continue
		}

	}
}

// ВНИМАНИЕ СТРУКТУРЫ В ГО ДОЛЖНЫ ИМЕТЬ КАПИТАЛ ВО ВСЕХ ФИЛДАХ
func main() {

	v := struct {
		FieldString string `json:"field_string"`
		FieldInt    int
		Slice       []int
		Object      struct {
			NestedField string
		}
	}{
		FieldString: "STROKA",
		FieldInt:    107,
		Slice:       []int{112, 107, 207},
		Object:      struct{ NestedField string }{NestedField: "KUKA"},
	}
	fmt.Println("Initial struct")
	PrintStruct(v)
	fmt.Println()
	/*
		interface{} is the "any" type, since all types implement the interface with no functions.
		map[string]interface{} is a map whose keys are strings and values are any type.
	*/
	fmt.Printf("Process struct..\n")
	values := map[string]interface{}{"NestedField":"BYAKA","FieldString": "OSAKA","FieldInt":"111"}

	functionOne(&v,values)

	fmt.Println("Processed struct")
	PrintStruct(v)


}
