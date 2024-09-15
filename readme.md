## Инструкция по сборке и запуску
Весь кодлежит в /src
Небольшая документация к коду с помощью go doc и pandoc documentation.html
Запуск:
- 1) из папки test_ozon выполнить команду **docker compose up --build**
    * 1.1) перейти на localhost:8080 будет playground для тестирования запросами
- 2) для запуска тестов отдельно выполнить  **docker compose run test** (после выполнения шага 1).
### materials
- в файле db.sql можно посмотреть как сутроена БД
- projectTree.txt дерево проекта
