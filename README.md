# anew(er)

Append lines from stdin to a file, but only if they don't already appear in the file.
Outputs new lines to `stdout` too, making it a bit like a `tee -a` that removes duplicates.

## Usage Example

Here, the files called `things.txt` and `things2.txt` contain lists of numbers. `newthings.txt` contains a second
list of numbers, some of which appear in `things*.txt` and some of which do not. `anewer` is used
to append the latter to the last file specified as a parameter (or the value set in the `-o` argument.


```
▶ cat things.txt
Zero
One
Two

▶ cat things2.txt
Zero
Six
Three

▶ cat newthings.txt
One
Two
Three
Four

▶ cat newthings.txt | anewer -o added-lines things.txt things2.txt
Four

▶ cat added-lines
Four

▶ cat newthings.txt | anewer things.txt things2.txt
Four

▶ cat things2.txt
Zero
Six
Three
Four

```

Note that the new lines added to `things2.txt` are also sent to `stdout`, this allows for them to
be redirected to another file or command:

```
▶ cat newthings.txt | anewer things.txt things2.txt | echo $(xargs)
Four
```

## Flags

- To view the output in stdout, but not append to the file, use the dry-run option `-d`.
- To append to the file, but not print anything to stdout, use quiet mode `-q`.
- To trim trailing whitespaces before making the comparison, use `-t`.
- To define a file to append the output to, without needing to specify it as a check value, use `-o`.
