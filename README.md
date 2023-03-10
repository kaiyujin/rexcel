# rexcel

rexcel is a tool that reads specific rows and cells from an Excel file.

## Install

### For mac
```
brew install kaiyujin/tap/rexcel
```

### Other
[Download file](https://github.com/kaiyujin/rexcel/releases)

## Usage

Specify the file path and sheet name each time, and output the third line.
```
rexcel -f=target.xlsx -s=Sheet1 -r=3
```

Read file path and sheet name from environment variables, and output B1 cell
```
REXCEL_FILE=target.xlsx
REXCEL_SHEET=Sheet1
rexcel -c=B1
```

## Args

- `-f` file path (Required if environment variable REXCEL_FILE is not present)
- `-s` sheet name (Required if environment variable REXCEL_SHEET is not present)
- Either is required
  - `-r` row number
  - `-c` cell
