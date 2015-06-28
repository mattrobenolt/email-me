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
usage: email-me [flags] [command]
  -max=10000: max bytes to capture for stdout/stderr
  -on-error=false: only notify on a non-0 exit code
  -s="": subject of email (optional)
  -to="": email address to send output to
```

## Example email

```
Cmd: [ls -lah] 
Start: 2015-06-24 18:43:30.195114617 -0700 PDT 
End: 2015-06-24 18:43:30.210375736 -0700 PDT 
Duration: 10.210247ms total 1.824ms user 4.066ms system
ProcessState: exit status 0 
Error: <nil> 
Stderr: 


Stdout: 
total 72 
drwxr-xr-x 14 matt staff 476B Jun 24 18:43 . 
drwxr-xr-x@ 219 matt staff 7.3K Jun 24 18:10 .. 
-rw-r--r-- 1 matt staff 34B Jun 24 18:38 .dockerignore 
drwxr-xr-x 13 matt staff 442B Jun 24 18:43 .git 
-rw-r--r-- 1 matt staff 275B Jun 24 18:38 .gitignore 
-rw-r--r-- 1 matt staff 507B Jun 24 17:32 Dockerfile 
-rw-r--r-- 1 matt staff 397B Jun 24 18:43 README.md 
drwxr-xr-x 12 matt staff 408B Jun 24 18:39 bin 
-rwxr-xr-x 1 matt staff 113B Jun 24 17:32 build.sh 
-rw-r--r-- 1 matt staff 1.5K Jun 24 18:32 command.go 
-rw-r--r-- 1 matt wheel 2.0K Jun 24 18:36 main.go 
-rw-r--r-- 1 matt staff 263B Jun 24 17:25 sendmail.go 
-rw-r--r-- 1 matt staff 193B Jun 24 17:26 smtp.go 
drwxr-xr-x 3 matt staff 102B Jun 24 18:18 src 
```
