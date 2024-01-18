# Экзамен &mdash; варинат №17
## Студент
Козлов Егор, 1/278

## Как пользоваться?
Чтобы выбрать нужный вариант, переключитесь на ветвь `case_<n>`, где `<n>` &mdash; номер варианта. 

Ветка `master` хранит обобщенную версию задания, на основе которой можно построить решение своего варината.

**Заметка:** для переименовывания переменных/структур в `VS Code` заложена функция на клавишу `<F2>`.

## Задание

Создайте структуру для предметной области "Самолет" с полями "Модель", "Тип" и "Год выпуска". Установите ограничения для года выпуска от 1900 до текущего года. Запретите устанавливать значения вне этого диапазона, возвращая ошибку при попытке. Тип самолета должен автоматически определяться на основе его модели, и не должен иметь собственного сеттера.

Создайте срез, содержащий несколько экземпляров этой структуры. Напишите метод для вычисления среднего года выпуска самолетов, который принимает срез структур в качестве входных данных и возвращает сумму годов выпуска.

Реализуйте метод с именем tryAdd, который принимает указатель на срез структур и экземпляр самолета. Этот метод должен добавить самолет в срез и вернуть true, если он еще не присутствует в срезе, иначе вернуть false.

## Решение
Решения находится в пакетах `airplanes` и `main`.

Чтобы запустить программу, выполоните:
```sh
go run ./
```
