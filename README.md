# Conventional Commit CLI

A command-line tool to help you create conventional commits with ease.

## Features

- 📝 Interactive prompts for creating conventional commits
- 🎨 Beautiful colored output
- ✨ Emoji support
- 🔄 Automatic git staging and committing
- 💡 Smart commit message formatting

## Installation

```bash
go install github.com/modecode22/cmg/cmd/cmg@latest
```

Make sure your Go installation is up to date and your `$GOPATH/bin` is in your system's `PATH`.

## Usage

Simply run:

```bash
cmg
```

The tool will guide you through creating a conventional commit with interactive prompts.

## Development

1. Clone the repository:

```bash
git clone https://github.com/modecode22/cmg.git
cd cmg
```

1. Install dependencies:

```bash
go mod download
```

1. Build the project:

```bash
make build
```

1. Install locally:

```bash
make install
```

## License

MIT License
