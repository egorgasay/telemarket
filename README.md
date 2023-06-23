# telemarket - online-store bot  [![CI](https://github.com/egorgasay/telemarket/actions/workflows/go.yml/badge.svg)](https://github.com/egorgasay/telemarket/actions/workflows/go.yml)

### ⚙️ Installation 

```bash
git clone https://github.com/egorgasay/telemarket
cd telemarket
export TELEGRAM_BOT_KEY=YOUR_BOT_KEY
make run
```

### 🔍️ Purpose

With this bot, you can easily sell clothes via Telegram.  

### 🏆 Challenges:
1. CI ✅
2. Deploy ✅
3. Messages with photos ✅
4. Personal Data Storage Agreement
5. Administrator Mode (Add, Remove, Change)
6. Analytics Mode (Watch stats)
7. Uploading statistics to Excel
8. Order tracking
9. Discount system (as module)
10. Bonus system (as module)
11. Mystery Box system (as module)

### 👕 Change items list

You can start selling your own products by changing the default values in the items.json file.

```json
items.json
[
    {
        "name": "t-shirt black",
        "description": "100% cotton",
        "price": 1500.00,
        "quantity":  1
    },
    {
        "name": "t-shirt white",
        "description": "100% cotton",
        "price": 1500.00,
        "quantity":  1
    }
]
```

### ✅ Run tests

```bash
make test
```
