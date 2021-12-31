# Hakuna Go

Hakuna Go is a cli for the timetracking tool [Hakuna](https://hakuna.ch).

## Usage

The simplest way to use the cli is to export two environment variables
with your hakuna company subdomain and your api token.

```shell
export HAKUNA_CLI_SUBDOMAIN="my-company-subdomain"
export HAKUNA_CLI_API_TOKEN="xxxxxxxxxxxxxxxxxxxx"

hakuna timer start --taskId=2 --note="Building cool stuff!"
```

Or using the `.hakuna.yaml` config file

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

The cli looks for a config in the current directory first.
If no config exits it fallsback to the config in the home directory.

you can check which config it uses by providing the `--debug` flag.

Environment variables take presidency over the config file.
