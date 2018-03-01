# jsonify-aws-dotfiles

## What is it?

The AWS config and credentials file (typically located in `~/.aws`) have values that you need to sometimes access before your application starts. Sometimes you run into situations where you need values from the config file and the SDK you are using only read the credentials file and not the config file. The AWS-CLI loads both, but doesn't have options for printing out the values it gets from those files. Fortunately, there are a miriade of tools for dealing with JSON. This command takes both and creates a usable JSON file from both of these files.

## How to install

1. Install the most recent version of [Go](https://golang.org/doc/install) to your machine.
2. [Follow the instructions](https://github.com/golang/go/wiki/GOPATH#integrating-gopath) for adding build go packages to your PATH.
3. Install with `go get -u github.com/paul-nelson-baker/jsonify-aws-dotfiles` 


