# Why as a Service (WaaS)

Ever needed a reason and could not be bothered to come up with one yourself? WaaS has you covered. It is a tiny Go API that answers the eternal question "why?" with a random reason pulled by [extremely advanced algorithms](https://github.com/tb12as/waas/blob/main/internal/app/reason.go#L19). The answers are not always good, but they are always confident.

Live at [waas.syfq.my.id](https://waas.syfq.my.id), open 24/7 and asking no questions.

## How It Works

1. You ask "why?"
2. The database shuffles its pile of excuses and picks one.
3. You receive that reason and act like you came up with it.

## API

There is exactly one endpoint, so it is hard to call the wrong one.

### `GET /`

Returns one random reason.

```json
{
    "status": 200,
    "message": "Success",
    "data": {
        "id": 56,
        "reason": "Because you can always apologize later."
    }
}
```

### Rate Limiting

Each IP gets 10 requests per second, with bursts of up to 20. If you need more reasons than that, the problem is probably not a lack of reasons. Go over the limit and you get:

```json
{ "error": "rate limit exceeded" }
```

That one is a reason too, if you think about it.

## Running It Yourself

### Prerequisites

- Go 1.26 or newer
- MySQL

### Configuration

Copy `.env.example` to `.env` and fill it in:

| Variable      | What it does                    | Default     |
|---------------|---------------------------------|-------------|
| `APP_ENV`     | Environment name                | `local`     |
| `APP_PORT`    | Port the API listens on         | `8080`      |
| `DB_NAME`     | MySQL database name             | `waas`      |
| `DB_USERNAME` | Database user                   | `root`      |
| `DB_PASSWORD` | Database password               | `root`      |
| `DB_HOST`     | Database host                   | `localhost` |
| `DB_PORT`     | Database port                   | `3306`      |

Please do not use `root`/`root` in production (BTW I know one server that does this aowkoakokw, f in stupid devs). Hackers need reasons to hack you, so don't give them one.

### Start

```bash
go run cmd/api/main.go
```

The schema is migrated automatically on startup. The table starts out empty, and an empty reasons table is a sad thing to watch, so insert your own reasons lol. 

## Built With

- **Go**, because it is fast and I need to learn it.
- **Gin**, because it is also fast and I also need to learn it.
- **GORM**, because writing SQL by hand is also a choice, which I avoided here.
- **MySQL**, because that's the only database exists in my poor shared hosting server :(
- **`ORDER BY RAND()`**, the most advanced randomness technology available.
- **Tears**, because I hate programming.
