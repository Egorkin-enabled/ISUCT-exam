# Экзамен &mdash; варинат обобщенный
## Студент
Козлов Егор, 1/278

## Как пользоваться?
Чтобы выбрать нужный вариант, переключитесь на ветвь `case_<n>`, где `<n>` &mdash; номер варианта. 

Ветка `master` хранит обобщенную версию задания, на основе которой можно построить решение своего варината.

**Заметка:** для переименовывания переменных/структур в `VS Code` заложена функция на клавишу `<F2>`.

## Задание

Создайте структуру для предметной области "Автомобиль" с полями "Марка", "Год выпуска" и "Скорость". Установите ограничения для года выпуска от 1900 до текущего года. Запретите устанавливать значения вне этого диапазона, возвращая ошибку при попытке. Скорость автомобиля должна автоматически определяться на основе его марки и года выпуска, и не должна иметь собственного сеттера.

Создайте срез, содержащий несколько экземпляров этой структуры. Напишите метод для вычисления средней скорости, который принимает срез структур в качестве входных данных и возвращает сумму скоростей.

Также реализуйте метод с именем tryAdd, который принимает указатель на срез структур и экземпляр автомобиля. Этот метод должен добавить автомобиль в срез и вернуть true, если он еще не присутствует в срезе, иначе вернуть false.

## Решение
Решения находится в пакетах `cars` и `main`.

Чтобы запустить программу, выполоните:
```sh
go run ./
```
