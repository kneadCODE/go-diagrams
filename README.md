# Go-Diagrams

## Fast and easy application diagrams

Go-Diagrams is a loose port of [diagrams](https://github.com/mingrammer/diagrams).

## Contents

- [Features](#features)
- [Usage](#usage)

## Features

Turn this:

```golang
d, err := diagram.New(diagram.Filename("app"), diagram.Label("App"), diagram.Direction("LR"))
if err != nil {
    log.Fatal(err)
}

dns := gcp.Network.Dns(diagram.NodeLabel("DNS"))
lb := gcp.Network.LoadBalancing(diagram.NodeLabel("NLB"))
cache := gcp.Database.Memorystore(diagram.NodeLabel("Cache"))
db := gcp.Database.Sql(diagram.NodeLabel("Database"))

dc := diagram.NewGroup("GCP")
dc.NewGroup("services").
    Label("Service Layer").
    Add(
        gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 1")),
        gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 2")),
        gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 3")),
    ).
    ConnectAllFrom(lb.ID(), diagram.Forward()).
    ConnectAllTo(cache.ID(), diagram.Forward())

dc.NewGroup("data").Label("Data Layer").Add(cache, db).Connect(cache, db)

d.Connect(dns, lb, diagram.Forward()).Group(dc)

if err := d.Render(); err != nil {
    log.Fatal(err)
}
```

Into this:

![app-diagram](images/app-diagram.png)

## Usage

```sh
go get github.com/blushft/go-diagrams
```

Create a diagram:

```golang
d, err := diagram.New(diagram.Label("my-diagram"), diagram.Filename("diagram"))
if err != nil {
    log.Fatal(err)
}

fw := generic.Network.Firewall().Label("fw")
sw := generic.Network.Switch().Label("sw")

d.Connect(fw, sw)
```

Render the output:

```golang
if err := d.Render(); err != nil {
    log.Fatal(err)
}
```

Go-Diagrams will create a folder in the current working directory with the graphviz DOT file and any image assets.

Create an ouput image with any graphviz compatible renderer:

```sh
dot -Tpng diagram.dot > diagram.png
```

## Development

### Prerequisites

- Go 1.18 or later (for generics support)
- [Task](https://taskfile.dev/) (optional, for running automation tasks)
- Graphviz (for rendering diagrams)

### Asset Embedding

This project uses [fileb0x](https://github.com/UnnoTed/fileb0x) to embed PNG icon assets as binary data in Go source files. When you update PNG files in `nodes/assets/`, you need to regenerate the embedded Go files.

### Regenerating Embedded Assets

After updating any PNG files in the `nodes/assets/` directory:

#### Using Taskfile (Recommended)

```sh
task generate
```

This will regenerate all embedded asset files (`b0xfile_*.go`) in the `nodes/assets/` directory.

#### Manual Approach

```sh
cd nodes/assets
go generate
```

**Note:** The `go generate` command will update all `b0xfile_*.go` files (1100+ files). If you only changed one asset, you can use:

```sh
task generate:only-storage
```

This task regenerates all assets but only commits the storage account changes, restoring all other files to their previous state.

### Available Tasks

Run `task help` or `task --list` to see all available tasks:

- `generate` - Regenerate embedded asset files from source PNGs
- `generate:only-storage` - Regenerate only Azure storage account asset (for testing)
- `test` - Run tests
- `lint` - Run linter
- `clean` - Clean generated files

### Building

```sh
go build ./...
```

### Testing

```sh
go test ./...
```

Or using Task:

```sh
task test
```
