## Запуск сервера ##

```bash
docker-compose up
go run github.com/cosmtrek/air
```
dump базы данных находится в файле [dump.sql](dump.sql)

docker запускает сервисы на портах:
- postgres: 5432
- pgadmin: 5050
- swagger: 8080

Тестировать можно через [swagger](http://localhost:8080/) 

Комментарии по коду:
>- Метод признания выручки – списывает из резерва деньги 
> (принимает id пользователя, услуги, заказа и его стоимость).
> Я передавала только orderId, так как его достаточно (уникален в рамках системы).

------

`POST /funds/reserve`- выдает ошибку. Не получилось найти причину, т.к. это первый опыт написания сервиса на Go.
Изначально запрос работал, но после ряда изменений в коде route начал падать. 
```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x7d19e9]
```
