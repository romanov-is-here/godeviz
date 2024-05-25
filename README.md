# godeviz

Dependency visualizer for go projects

⚠️ Prototype code ahead. You've been warned ⚠️


### Roadmap

- Drill in в фильтре
- Показывать ошибку получения данных
- Карточка выделенного нода в UI
- Drill in в UI
- Пропускать сгенерированные
- Пропускать тесты
- Статистика по выделенному ноду (список импортов in/out)
- Обсудить формат развертывания. Если запускаем сервис, надо разделять локальный и сервисный запуск
- Возможность делиться графом через файл/pastebin
- Отличать локальные пути и путь до репозитория

### Требования к визуализатору графов

- Можно селектить вершины и ребра
- Можно добавлять тултип на ховер по вершинам и ребрам
- Понятный лэйаут без вермишели ()
- ? Возможность построить Path по вершинам (для визуализации циклов)
- ? Можно двигать вершины (если лэйаут хороший, можно и не двигать)

### Исправление плохого кода

- Зафиксировать тестами работу
- Рефакторинг фронта