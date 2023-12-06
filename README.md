# To-Do-List-App-Golang
Ini merupakan penyelesaian tugas membuat To-Do List App Golang

Run main.go

open PostMan

GET http://localhost:8080/tasks --> Send

POST http://localhost:8080/tasks Choose raw --> JSON { "title": "masukkan_tugas_baru", "done": false } Send

UPDATE http://localhost:8080/tasks/ID Choose raw --> JSON { "id": 1, "title": "edit tugas", "done": true }

DELETE http://localhost:8080/tasks/ID Send