# email-me

sometimes you run processes and you want to be emailed when they're done running

```
$ email-me -to=m@robenolt.com ls -lah
```

## Installation

Compiled binaries can be found under: https://github.com/mattrobenolt/email-me/releases, or
if you're familiar with Go, you can install via `go get`:

```bash
$ go get github.com/mattrobenolt/email-me
```

## Usage

```
NAME:
email-me - email me when a thing is done

USAGE:
email-me [global options] [arguments...]

VERSION:
0.3.0

GLOBAL OPTIONS:
--to 		email address to send output to [$EMAIL_ME_TO]
--subject, -s 	subject of email (optional)
--max "10000"	max bytes to capture for stdout/stderr [$EMAIL_ME_MAX]
--on-error		only notify on a non-0 exit code
```

## Example email

```
Cmd: [ls -lah]
Start: 2015-06-28 08:10:46.880877935 -0700 PDT
End: 2015-06-28 08:10:46.899524717 -0700 PDT
Duration: 18.646782ms total 1.981ms user 7.601ms system
ProcessState: exit status 0
Error: <nil>
Stderr:


Stdout:
total 11416
drwxr-xr-x   16 matt  staff   544B Jun 28 08:10 .
drwxr-xr-x@ 219 matt  staff   7.3K Jun 24 18:11 ..
-rw-r--r--    1 matt  staff    34B Jun 24 18:38 .dockerignore
drwxr-xr-x   13 matt  staff   442B Jun 28 08:10 .git
-rw-r--r--    1 matt  staff   275B Jun 24 18:38 .gitignore
-rw-r--r--    1 matt  staff   507B Jun 24 17:32 Dockerfile
-rw-------    1 matt  staff   1.4K Jun 28 06:55 LICENSE
-rw-r--r--    1 matt  staff   1.5K Jun 28 08:06 README.md
drwxr-xr-x   12 matt  staff   408B Jun 28 08:03 bin
-rwxr-xr-x    1 matt  staff   109B Jun 24 18:48 build.sh
-rw-r--r--    1 matt  staff   1.6K Jun 25 10:42 command.go
-rwx------    1 matt  staff   5.5M Jun 28 07:06 email-me
-rw-r--r--    1 matt  staff   2.3K Jun 28 08:01 main.go
-rw-r--r--    1 matt  staff   263B Jun 24 17:25 sendmail.go
-rw-r--r--    1 matt  staff   193B Jun 24 17:26 smtp.go
drwxr-xr-x    3 matt  staff   102B Jun 24 18:18 src
```
