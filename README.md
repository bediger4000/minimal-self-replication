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
If my minimal programming language can execute this quine,
I consider it finished and successful.

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
Before assignment to a variable,
unquoted and double-quoted strings have any variable's values interpolated.
Variable values must be set before use.
Interpolated variables appear as a dollar sign (`$`) immediately followed by
a name: `$some_name`.

### Statement examples

My small language is definitely influenced by traditional Unix shell syntax.
It should constitute a strict subset of that traditional shell syntax.

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

My programming language does not allow variable interpolation
on the left had side of an `=`.
The variable name, the `=`, and the string literal must be adjacent,
with no white space.

#### Output

All output is via `echo` statements:

```
echo 'now is the time'
v='now is the time
echo $v
```
The 3-line program above produces this output:

```
now is the time
now is the time
```

Double-quoted strings have variables (if any appear) interpolated before output.

```
a="now is the time "
b="for all good men "
c="to come to the aid of their country"
echo "$a$b$c"
```

That 4-line program produces "now is the time for all good men to come to the aid of their country".

## Building the interpreter

```
$ git clone https://github.com/bediger4000/minimal-self-replication.git
$ cd minimal-self-replication
$ go test -v ./...
$ go build -o mpl $PWD
$ ./mpl minquine
$ ./check
```

The script `check` uses `bash`, `dash`, `zsh`, `ksh` and the minimal
programming language interpreter to run the quine.
All outputs are lexically equal.

#### Design

The idea was to implement a small subset of traditional Unix shell language:
variable assignment, string interpolations in string literals,
and `echo` statements.
I did not use a compiler-style lexer/parser, and walk the parse tree to execute programs.
That seemed like overkill.

Because the two kinds of statements can be separated by semicolons,
my lexing and parsing code breaks the input into "command strings",
and then parses each command string as part of execution.
One criticism I've read of traditional Unix shells is that they don't
really have a grammar.
I may have encountered a small example of that.

I do have a separate goroutine breaking the input file contents
into command strings,
passing the command strings to the main goroutine
via a channel.
This style simplifies the code of by avoiding
having to retain enough state to re-start the lexing
when a command string is found.

Parsing without a parser-generator is idiosyncratic.
I include a file named `oddities`, which is valid code in my programming language,
containing some I discovered.
I expect other weirdness exists.

Unlike in traditional Unix shells, the `echo` part of output statements
does not execute a sub process.
`echo` is the only built-in.

Since my programming language does not have looping, if/then/else, recursion,
or `goto`, it is not [Turing complete](https://en.wikipedia.org/wiki/Turing_completeness).
Every program does terminate in a finite number of steps.
