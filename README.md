
## Полная информация о Swagger

В репозитории библиотеки для работы со swagger на go, подробнее описаны все возможности и правила работы со swgger
`https://github.com/swaggo/swag?tab=readme-ov-file`

## Запуск

Для запуска примера достаточно запустить exe файл из репозитория
Но при таком запуске документация будет не доступна

Для запуска примера с документацией необходимо добавить флаг `-d` со значением `true`

```
swaggerExample.exe -d=true
```

### Добавление новых endpoint - ов и их документирование

Для примера добавим endpont, который будет умножать полученные данные

Для этого создадим новый файл папке service и напишем следующий код
``` Golang
package service

  
import (
    "encoding/json"
    "net/http"
)

type MultiplicationData struct {
    Data []float64
}

type MultiplicationResponce struct {
    Sum float64
} 

// @Summary     Перемножение массива чисел
// @Tags        numbers
// @Description Перемножение массива чисел переданого в структуре MultiplicationData
// @ID          MultiplicationSlice
// @Accept      json
// @Produce     json
// @Param       data        body        MultiplicationData     true    "Данные"
// @Success     200         {object}    MultiplicationResponce         "success"
// @Failure     400,500     {object}    errorResponce                  "error"
// @Router      /Multiplication [post]
func MultiplicationSlice(w http.ResponseWriter, r *http.Request) {
    var data MultiplicationData
  
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        ErrorResponce(err, "invalid data", http.StatusBadRequest, w)
        return
    }
  
    var responce MultiplicationResponce  

    for _, v := range data.Data {
        responce.Sum *= v
    }
  
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(responce); err != nil {
        ErrorResponce(err, "error encode", http.StatusInternalServerError, w)
    }
}
```

На основе комментария над функцией MultiplicationSlice будет составлена документация swagger - ом.

`// @Summary` - Описание что видно до раскрытия подробной информации
`// @Tags` - тег для группировок endpoint по какому либо признаку, тег придумывается самостоятельно и является заголовком группы
`// @Description` - Подробное описание endpoint
`// @ID` - id endpoint
`// @Accept` - тип принимаемого запроса
`// @Produce` - тип отправляемого ответа
`// @Param` - Принимаемы параметры 
	`// @Param       data        body        MultiplicationData     true    "Данные"`
		data - имя параметра 
		body - где находятся данные, в данном случае в теле запроса, если параметры в запросе то пишется query
		MultiplicationData - тип данных ответа
		true - обязательность этого параметра
		"Данные" - Описание
`// @Success` - описание успешного ответа
	`// @Success     200         {object}    MultiplicationResponce         "success"`
		код ответа
		тип данных ответа
		тип данных данных внутри кода
		описание
`// @Failure` - описание неудачных ответов, заполняется аналогично успешному ответу
`// @Router` - путь до этого endpoint, без него swagger не документирует 

>Подробнее о том какие настройки существуют в swagger смотрите в [репозитории](https://github.com/swaggo/swag?tab=readme-ov-file) 

Так же добавим в роутер новый endpoint в файле main.go

``` Golang
r.HandleFunc("/Multiplication", service.MultiplicationSlice).Methods("POST")
```

Запустим генерацию обновленной документации 

1. Убедимся что swagger установлен 
``` cmd
   go install github.com/swaggo/swag/cmd/swag@latest`
```
2. Запустим генерацию
``` cmd
   swag init
```
3. Соберём новый билд 
```cmd
go build
```

Теперь можно запустить и увидеть в документации описание нового endpoint