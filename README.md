# Fuck Git

Шутим и практикуемся

## Как поднять локально

### 1. Backend
Можно запустить через docker-compose или через makefile

```bash
make up
 ```
или аналогично
```bash
docker-compose up -d
```


### 2. Frontend
для запуска фронта нужно установить зависимости и запустить сервер
```bash
make install_front
```
```bash
make run_front
```

наверное?


## Ручка `/questions`
возвращает в  `payload` бинарное дерево из объектов Question.

`root` — объект `Question`, корень дерева;

`left` и `right` — тоже обекты `Question`, потомки корня на ответ "нет" и "да" соответственно.

```json
{
  "status": "ok",
  "payload": {
    "root": ...
  }
}
```

### Пример

```json
{
  "payload": {
    "root": {
      "left": {
        "left": null,
        "right": null,
        "solution": {
          "url": "/",
          "command": "",
          "description": ""
        },
        "name": "unknown",
        "body": "Тогда не знаю, чем помочь, друг"
      },
      "right": {
        "left": {
          "left": {
            "left": null,
            "right": null,
            "solution": {
              "url": "/",
              "command": "",
              "description": ""
            },
            "name": "unknown",
            "body": "Тогда не знаю, чем помочь, друг"
          },
          "right": {
            "left": null,
            "right": null,
            "solution": {
              "url": "/solution",
              "command": "git push --force blabla",
              "description": "Откатить изменения в удаленном репозитории"
            },
            "name": "remotely_true",
            "body": ""
          },
          "solution": null,
          "name": "remotely",
          "body": "Нужно откатить изменения в удаленном репозитории???"
        },
        "right": {
          "left": null,
          "right": null,
          "solution": {
            "url": "/solution",
            "command": "git reset --hard HEAD~1",
            "description": "Откатить изменения локально, cтерев их из истории"
          },
          "name": "locally_true",
          "body": ""
        },
        "solution": null,
        "name": "locally",
        "body": "Нужно откатить изменения локально?"
      },
      "solution": null,
      "name": "start",
      "body": "Нужно откатить изменения?"
    }
  },
  "status": "ok"
}

```

## Модели

### Solution
`solution` содержит `url` для ссылки на документацию, `description` — описание решения, и `command` — команду для решение проблемы.
```go
type Solution struct {
	URL         string `json:"url"`
	Command     string `json:"command"`
	Description string `json:"description"`
}
```
### Пример
```json
{
    "url": "/solution",
    "command": "git reset --hard HEAD~1",
    "description": "Откатить изменения локально, cтерев их из истории"
}
```

### Question
`question ` содержит в себе `name`, `body`, `solution` и `left` и `right` поля.

- `body` - текст вопроса
- `name` - внутреннее название вопроса (для удобства и возможно поиска в будущем)
```go
type Question struct {
    Left     *Question `json:"left"`
    Right    *Question `json:"right"`
    Solution *Solution `json:"solution"`
    Name     string    `json:"name"`
    Body     string    `json:"body"`
}
```
### Пример
```json
{
    "left": null,
    "right": null,
    "solution": {
            "url": "https://git-scm.com/docs/git-reset",
            "command": "git reset --hard HEAD~1",
            "description": "Откатить изменения локально, cтерев их из истории"
        },,
    "name": "end",
    "body": "Нужно откатить только локальные изменения?"
}
```
