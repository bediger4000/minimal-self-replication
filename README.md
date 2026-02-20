# Minimal programming language that allows non-vacuous self-replicating programs

I found a [quine](),
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
