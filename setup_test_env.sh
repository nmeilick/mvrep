#!/bin/bash

set -eu

# Create test directory
mkdir -p test

# Navigate to test directory
cd test

# Create various test files
touch file1.txt
touch file2.txt
touch file3.txt
touch file_with_special_chars_!@#.txt
touch file_with_numbers_123.txt
touch file_with_spaces.txt
touch file_with_underscores.txt
touch file_with-dashes.txt

# Create a subdirectory and files within it
mkdir subdir
touch subdir/file_in_subdir.txt

# Create a file with no permissions
touch no_permissions.txt
chmod 000 no_permissions.txt

# Create a file with read-only permissions
touch read_only.txt
chmod 444 read_only.txt

# Create a file with write-only permissions
touch write_only.txt
chmod 222 write_only.txt

# Create a file with execute-only permissions
touch execute_only.txt
chmod 111 execute_only.txt

# Create a hidden file
touch .hidden_file.txt

# Create a symbolic link
ln -s file1.txt symlink_to_file1.txt

echo "Test environment setup complete."
