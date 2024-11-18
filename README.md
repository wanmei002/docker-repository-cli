## install
```shell
make build
```
if you not install golang, use bin file `delete-image`

## how to use
```shell
./delete-image --help
Usage:
   [command]

Available Commands:
  catalog      Get all catalog images
  completion   Generate the autocompletion script for the specified shell
  delete-image Delete image
  help         Help about any command
  manifest     Get image manifest
  tag          image tags

Flags:
  -h, --help                   help for this command
  -H, --host string            the host to connect to (default "127.0.0.1:5000")
  -u, --user-password string   basic auth user (default "")

Use " [command] --help" for more information about a command.
```
> --user-password value: echo -n username:password | base64

## example
### get catalog
```shell
./delete-image catalog -u username:password-base64
```
### get image tags
```shell
./delete-image tag --repo public/redis -u username:password-base64
```

### delete image tag
```shell
./delete-image delete-image --repo public/redis --tag v1.0.1 -u username:password-base64
```

