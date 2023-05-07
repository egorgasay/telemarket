# telemarket - online-store bot

### ⚙️ Installation

```bash
git clone https://github.com/egorgasay/telemarket
cd telemarket
make run
```

### 🔍️ Purpose

With this bot, you can easily sell clothes via Telegram.  

### 👕 Change items list

You can start selling your own products by changing the default values in the config.json file.

```json
config.json
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
