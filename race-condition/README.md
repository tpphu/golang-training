
## Tips: Command line
> [How do I save terminal output to a file?](https://askubuntu.com/questions/420981/how-do-i-save-terminal-output-to-a-file)

```sh
go build -o race2 race2.go && for i in {1..100}; do ./race2 >> race2.txt 2>&1; done;
```