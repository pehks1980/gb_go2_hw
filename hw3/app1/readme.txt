Коментарий по заданию,

получилось с бубном сделать свой модуль, выставить его на github по имени github делать тест.(testmod_test.go)

не получилось сделать v2 делал по методичке и по google.com,

не понимаю (и не смог починить) почему оно не может найти v2 выдает такое сообщение:

~/go/go2hw/gb_go2_hw/hw3/app1$ go get -v github.com/pehks1980/testmod@v0
go: github.com/pehks1980/testmod v0 => v0.0.1
user@user-HP-Pavilion-17-Notebook-PC:~/go/go2hw/gb_go2_hw/hw3/app1$ go get -v github.com/pehks1980/testmod@v1
go: github.com/pehks1980/testmod v1 => v1.0.0
user@user-HP-Pavilion-17-Notebook-PC:~/go/go2hw/gb_go2_hw/hw3/app1$ go get -v github.com/pehks1980/testmod@v2
go get github.com/pehks1980/testmod@v2: no matching versions for query "v2"

Ps методичка написана не очень ясно, на видео уроке тоже, не очень понятно.

ps перезадал тег удалил старые создал из папки v2 новый
git tag -d v2.0.0
commit..
удаление тега вручную из github...

git tag v2.0.0
git push --tags

после этого модуль начал грузиться
go get github.com/pehks1980/testmod/v2
go: downloading github.com/pehks1980/testmod/v2 v2.0.0

причем также как и раньше такой запрос не находит v2
/go/go2hw/gb_go2_hw/hw3/app1$ go get github.com/pehks1980/testmod@v2
go get github.com/pehks1980/testmod@v2: no matching versions for query "v2"

тест версий проходит
/go/go2hw/gb_go2_hw/hw3/app1$ go test
PASS
ok      github.com/pehks1980/gb_go2_hw/hw3/app1 0.003s

соотвественно вывод - может быть, go get github.com/pehks1980/testmod@v2 это вообще неправильно так запрашивать?

по моему в методичке надо прямо по шагам сделать пример, а то тут не там точечку поставишь все - каюк, и самое смешное
совершенно непонятно как его исправлять!
