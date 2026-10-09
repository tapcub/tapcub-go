# tapcub-go

Server-side Go SDK for [TapCub](https://tapcub.com), MIT licensed. Requires Go 1.21+ and uses only the standard library.

> This repository is a read-only mirror. Development happens in the TapCub main repository and is synced here as snapshots. Please open an issue for bugs and suggestions.

```bash
go get github.com/tapcub/tapcub-go
```

```go
import tapcub "github.com/tapcub/tapcub-go"

client := tapcub.New("<server API key>", "https://api.pathclue.com")
err := client.Track("purchase", map[string]any{"login_id": "u_1001", "properties": map[string]any{"revenue": 19.9}})
```

## Releases

Versions are git tags such as `v0.1.0`, pushed from the TapCub main repository with `pnpm sdk:sync -- --push --tag go-v0.1.0`. The Go module proxy fetches tags directly, so there is no separate publish step.

## License

MIT, see [LICENSE](LICENSE).
