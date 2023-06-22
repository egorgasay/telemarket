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
1. Rejection of the SDK ✅
2. CI ✅
3. Deploy✅
4. Replacing photos in existing messages
5. Personal Data Storage Agreement
6. Administrator Mode (Add, Remove, Change)
7. Analytics Mode (Watch stats)
8. Uploading statistics to Excel
9. Order tracking
10. Discount system (as module)
11. Bonus system (as module)

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
