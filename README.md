# gfnls
A tool to generate a listing of all Nvidia GeForce Now games using their JSON data set as well as archive.org's copy.

## Usage
Usage of ./gfnls:
  -source string
        Source for the data, either filename or archive.org timestamp

Feel free to suggest a better usage or other features in the issues this is a very early test

### Comparing games added or removed

On a system with bash you can use gfnls to show what games have been added or removed by piping the output through diff

`diff -urNp <(./gfnls -source 20200205085751) <(./gfnls)`

This will show you what games have been added or removed since february 5th 2020 until the current time in unified diff format.
