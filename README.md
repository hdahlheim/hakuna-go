# Hakuna Go

Hakuna Go is an unofficial CLI for the time-tracking tool [Hakuna](https://hakuna.ch).

## Installation

```shell
curl -sLO https://github.com/hdahlheim/hakuna-go/releases/latest/download/hakuna-go-macos-amd64
chmod +x hakuna-go-macos-amd64
mv hakuna-go-macos-amd64 hakuna
```

## Usage

The simplest way to use the CLI is to export two environment variables
with your hakuna company subdomain and your api token.

```shell
export HAKUNA_CLI_SUBDOMAIN="my-company-subdomain"
export HAKUNA_CLI_API_TOKEN="xxxxxxxxxxxxxxxxxxxx"
# then you can use the CLI to start a timer
hakuna timer start --taskId=2 --note="Building cool stuff!"
```

If you don't want to use environment variables you can use the `.hakuna.yaml` config file.
The CLI searches for the config file in two places, the first in the current directory and
after that in the user home directory.

Example using the config file:

```shell
cat << EOF > ~/.hakuna.yaml
subdomain: my-company-subdomain
api_token: xxxxxxxxxxxxxxxxxxxx
default:
  task_id: 2
EOF
```

Now you can start a timer without specifying a taskId

```shell
hakuna timer start
```

The CLI looks for a config in the current directory first.
If no config exits it fallsback to the config in the home directory.

You can check which config is being used by providing the `--debug` flag.

NOTE: Environment variables take presidency over the config file.
