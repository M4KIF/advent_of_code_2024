# **Advent of _code_ 2024**

## ***Go prerequisites***
`Linux(Arch)`
- Install latest go via `sudo pacman -S go`

## ***Building Go solutions***
- Open Your terminal of choice
- Traverse into the `.../src` directory inside the chosen puzzle dir from the day that You're interested in
- run `go install`

`Linux`

- run `go build -o <preferred-output-name, f.e. gingerbread>` 
- run `./<preferred-output-name, f.e. gingerbread>`

`Windows`

- run `go build -o <preferred-output-name, f.e. gingerbread.exe>`
- run `.\<preferred-output-name, f.e. gingerbread>.exe` or double click on built binary

## ***Python prerequisites***

`Linux(Arch)`
- Install _pipenv_ via `sudo pacman -S python-pipenv`
- Install _pyenv_ via `curl https://pyenv.run | bash`
- Add needed env vars, ie. 
`echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bashrc
echo 'command -v pyenv >/dev/null || export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(pyenv init -)"' >> ~/.bashrc`

## ***Running Python solutions***
- Open Your terminal of choice
- Traverse into the `<>/src` directory inside the chosen puzzle dir from the day that You're interested in
- run `pipenv install`
- run `pipenv shell`
- to run tests, run `pytest`
- to fire up the solution, run `<your python on path> main.py`

## ***Java prerequisites***
