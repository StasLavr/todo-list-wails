## Запуск 
Windows
Требование:
gcc c cgo (https://jmeubank.github.io/tdm-gcc/) из-за sqlite
```
git clone https://github.com/StasLavr/todo-list-wails.git
cd todo-list-wails
go env -w CGO_ENABLED=1
wails dev
```
## Building

To build a redistributable, production mode package, use `wails build`.

## Стек ##
Go vuejs tailwind css sqlite wails 
