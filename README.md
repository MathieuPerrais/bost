# B[logp]ost

Interact with jekyll posts because I don't always remember what's the date of my blogpost and I don't remember either what's the template for an empty post.

## Usage

```
bost [global options] command [command options] [arguments...]
```

## Commands

- *create*
- *open* without really know the name of the file. ```bost open my-first``` will open ```2016-01-29-my-first-post```
- *search*


## Help

```
NAME:
   bost - interact with jekyll posts

USAGE:
   bost [global options] command [command options] [arguments...]

VERSION:
   0.2.0

COMMANDS:
   create, c	create a blog post
   open, o	open blog post lazily
   search, s	search for posts
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --directory, -d "_posts"	directory where to look for posts
   --help, -h			show help
   --version, -v		print the version
```
