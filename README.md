# blrun

A spiritual successor to [blbuild](https://github.com/BendyLand/blbuild) (and is somehow working before it 😅), `blrun` is a simple run tool which automates the process of building and running compiled code.

## Usage

`blrun` utilizes a TOML file for the configuration of the build and run steps. The basic structure is as follows:

```toml
compiler = ""
path = ""
files = [""]
extras = ""
run = ""
```

If a file called "blrun.toml" does not exist in the project's root directory, then the program will prompt you for each value and generate one automatically.  

## Examples

```toml
# blrun.toml, placed in the root directory of the project
compiler = "gcc"
path = ""
files = ["main.c"]
extras = ""
run = "./a.out"
# Result:
# gcc main.c
# ./a.out
```

```toml
# blrun.toml
compiler = "g++ --std=c++20"
path = "files"
files = ["one.cpp", "two.cpp", "three.cpp"]
extras = "-o run"
run = "./run"
# g++ --std=c++20 files/one.cpp files/two.cpp files/three.cpp -o run
# ./run
```

```toml
# blrun.toml
compiler = "scalac"
path = "src/main"
files = ["Main.scala", "Utils.scala"]
extras = "-d src/build"
run = "scala -classpath src/build main.run"
# scalac src/main/Main.scala src/main/Utils.scala -d src/build
# scala -classpath src/build main.run
```
