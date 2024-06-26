# golang junior dev



## Задачи

### Exchanges

Написать REST сервис по расчету всех вариантов размена для указанной суммы денег. На вход принимается HTTP запрос в формате:

```json
{
  "amount": 400,
  "banknotes": [5000, 2000, 1000, 500, 200, 100, 50]
}
```

где
* **amount** – _сумма денег_
* **banknotes** – _доступные номиналы банкнот_

Формат ответа:

```json
{
  "exchanges": [
    [200, 200],
    [200, 100, 100],
    [200, 100, 50, 50],
    [200, 50, 50, 50, 50],
    [100, 100, 100, 100],
    [100, 100, 100, 50, 50],
    [100, 100, 50, 50, 50, 50],
    [100, 50, 50, 50, 50, 50, 50],
    [50, 50, 50, 50, 50, 50, 50, 50]
  ]
}
```

## Требования к оформлению

- конфигурация (хост, порт, уровень логирования) 
- graceful shutdown
- оформлен в общедоступном git репозитории
