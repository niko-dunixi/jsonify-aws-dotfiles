# jsonify-aws-dotfiles

## What is it?

The AWS config and credentials file (typically located in `~/.aws`) have values that you need to sometimes access before your application starts. Sometimes you run into situations where you need values from the config file and the SDK you are using only read the credentials file and not the config file. The AWS-CLI loads both, but doesn't have options for printing out the values it gets from those files. Fortunately, there are a miriade of tools for dealing with JSON. This command takes both and creates a usable JSON file from both of these files.

## How to install

1. Install the most recent version of [Go](https://golang.org/doc/install) to your machine.
2. [Follow the instructions](https://github.com/golang/go/wiki/GOPATH#integrating-gopath) for adding build go packages to your PATH.
3. Install with `go get -u github.com/paul-nelson-baker/jsonify-aws-dotfiles`

## How to use

Running the command will load the two files (`~/.aws/config` and `~/.aws/credentials`) and will dump the JSON file to standard out. You can manually specify either file with the `--config-file=` or `--credentials-file=`. You can toggle logging with `--verbose`, it will print to stderr as to not interfere with the JSON written to stdout. If desired, an output file can be used instead of writing the JSON string to file instead of stdout with `--output-file=`

Sample JSON output:

```
{
  "config": {
    "paulbaker": {
      "region": "us-east-1",
      "role_arn": "arn:aws:iam::12345:role/something",
      "source_profile": "alternate"
    },
    "alternate": {
      "region": "us-west-2",
      "output": "json"
    }
  },
  "credentials": {
    "alternate": {
      "aws_access_key_id": "...",
      "aws_secret_access_key": "..."
    }
  }
}
```
