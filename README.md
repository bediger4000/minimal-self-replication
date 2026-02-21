# Minimal programming language that allows non-vacuous self-replicating programs

I found a [quine](https://en.wikipedia.org/wiki/Quine_(computing)),
a self-replicating program,
I wrote no later than 1989.

```
st='echo st=$sq${st}$sq;echo dq=$sq${dq}$sq;echo sq=$dq${sq}$dq;echo $st'
dq='"'
sq="'"

echo st=$sq${st}$sq;echo dq=$sq${dq}$sq;echo sq=$dq${sq}$dq;echo $st
```

That program actually writes out the quine on stdout.
To prove actual self-replicating, you should do something like:

```
$ sh ./quine > q1
$ sh q1 > q2
$ diff q1 q2
```

This quine works with `zsh` v5.9, `bash` 5.3.9, `dash` 0.5.12, `ksh` 2020.0.0,
and even with the "heirloom shell", version 050706.
Weirdly, I wrote `${st}`, variable names surrounded by curly braces,
which aren't strictly necessary.

If an interpreter is to run this program, it has to have:

1. Assignment of a string to a variable.
2. An `echo` command to print to stdout.
3. Single-quoted strings in which no variable substitution takes place.
4. Double-quoted strings so that a string consisting of one single-quote can be assigned.
    - It could also have an escape character to do the same thing.
5. Interpolation of named variables, the `$st` above.
    - This has to happen before the `echo` command executes.
6. Program statements separated by both newlines and semicolons.

Note that the `${name}` variable interpolation notation isn't needed.
I'm going to simplify my 1980s quine to this:

```
st='echo st=$sq$st$sq;echo dq=$sq$dq$sq;echo sq=$dq$sq$dq;echo $st'
dq='"'
sq="'"
echo st=$sq$st$sq;echo dq=$sq$dq$sq;echo sq=$dq$sq$dq;echo $st 
```

That works identically to my original in many shells,
and simplifies variable value interpolation.
