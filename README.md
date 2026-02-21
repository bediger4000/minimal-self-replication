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
and even with the [heirloom shell](https://sourceforge.net/projects/heirloom/files/heirloom-sh/),
version 050706.
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

The quine doesn't strictly need `${name}` variable interpolation.
I simplified my 1980s quine to this:

```
st='echo st=$sq$st$sq;echo dq=$sq$dq$sq;echo sq=$dq$sq$dq;echo $st'
dq='"'
sq="'"
echo st=$sq$st$sq;echo dq=$sq$dq$sq;echo sq=$dq$sq$dq;echo $st 
```

The above code replicates in all the shells I tried,
and simplifies variable value interpolation code.

## My programming language

My programming language (`mpl`) has two kinds of statements:

1. Variable assignment: `name=string`
2. Print to stdout: `echo string`

Statements are executed in the order in which they appear in the interpreter's
single input file.

Variable names are upper and lower case letters, digits and underscores.
The have a letter as their first character.

Strings are consecutive characters.
They may appear as single-quoted (`'letters in a string'`)
or double-quoted (`"letters in a string"`),
or even unquoted.
Before assigment to a variable,
unquoted and double-quoted strings have any variable's values interpolated.
Variable values must be set before use.
Interpolated variables appear as a dollar sign (`$`) immediately followed by
a name: `$some_name`.

### Statement examples

#### Assignment

```
a='now is the time for all good men'
b=now is the time for all good men
c="now is the time for all good men"
```

Variables `a`, `b` and `c` will all have a value of "now is the time for all good men"
(without quotes) when interpolated.

```
d=defgh
a="abc$d"
b='abc$d'
c=abc$d
```

Variables `a` and 'c' will evaluate to "abcdefgh" on interpolation,
Variable `b` will evaluate to "abc$d".
Since the string assigned to `b` has single-quotes, no variable interpolation takes place.

#### Output

## Building the interpreter
