<p>
<div align="center">
<img src="logo.png" width="630" alt="mvrep logo"></p>
</div>

# mvrep

`mvrep` is a powerful command-line tool for renaming files and directories based on a pattern or fixed string substitution. It supports various options for dry-run, fixed string substitution, interactive mode, and more.

## Features

- **Regex and Fixed String Substitution**: Rename files and directories using regex or fixed string substitution.
- **Substring Search**: Both regex and fixed string modes perform substring searches by default.
- **Dry-Run Mode**: Preview changes without making actual modifications.
- **Interactive Mode**: Confirm each renaming operation.
- **Force Overwrite**: Automatically overwrite existing files.
- **Shell Command Output**: Output shell commands that would perform the renaming.
- **Back-References in Regex**: Support for back-references in regex substitution (e.g., `\1` for the first capture group).

## Installation

### Using `go install`

To install `mvrep` using `go install`, run the following command:

```sh
go install github.com/nmeilick/mvrep@latest
```

### Downloading Binaries

You can download pre-built binaries from the [release page](https://github.com/nmeilick/mvrep/releases).

## Usage

```
mvrep [options] <pattern> <replacement> <files...>
```

### Options

- `-n, --dry-run`: Only show what would be changed.
- `-F, --fixed`: Use fixed string substitution instead of regex.
- `-i, --interactive`: Interactive mode (y/n).
- `-f, --force`: Force overwrite of existing files.
- `-s, --shell`: Output shell commands that would perform the renaming.
- `-V, --version`: Show version.

### Examples

#### Rename Files Using Regex

```sh
mvrep 'image_(\d+)' 'photo_\1' *.jpg
```
This will rename files like "image_001.jpg" to "photo_001.jpg". Note that the pattern matches anywhere in the filename
and you need to anchor the search pattern explicitely with ^ or $.

#### Rename Files Using Fixed String Substitution

```sh
mvrep -F 'old_prefix_' 'new_prefix_' *
```
This will rename files like "old_prefix_document.txt" to "new_prefix_document.txt". The substitution occurs wherever
the string is found in the filename, not just at the beginning.

#### Preview Changes Without Making Actual Modifications

```sh
mvrep -n '\.jpeg' '.jpg' *
```
This will show how files containing ".jpeg" would be renamed to use ".jpg" without actually renaming them.

#### Interactive Mode

```sh
mvrep -i 'report_(\d{4})' 'annual_report_\1' *
```
This will prompt for confirmation before renaming each file matching the pattern.

#### Force Overwrite of Existing Files

```sh
mvrep -f 'draft_' 'final_' *
```
This will rename files even if the new names already exist, potentially overwriting files.

#### Output Shell Commands That Would Perform the Renaming

```sh
mvrep -s '(\w+)_backup\.(\w+)' '\1.\2' *
```
This will output the shell commands to rename backup files without the "_backup" suffix.

#### Use Back-References in Regex Substitution

```sh
mvrep '(\d{4})-(\d{2})-(\d{2})_log\.txt' 'log_\1\2\3.txt' *.txt
```
This will rename files like "2023-05-15_log.txt" to "log_20230515.txt".

## Building

To build the project for multiple platforms, use the provided `Makefile`:

```sh
make all
```

The binaries will be placed in the `bin` directory.

## Cleaning

To clean the build artifacts, use:

```sh
make clean
```

## License

This project is licensed under the MIT License.
