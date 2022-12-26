# FileToHex
 Convert file data to hex string

# Objective
- Print file data in hex string with 0x prefix with comma separated.
- This output will allow you to get a []uint8 literal snippets.

## How to use

- Input file path and maximum byte size in hex.
- Input negative integer if entire file wants to be printed.

```
...\FileToHex>type HelloWorld.txt
Hello, World!
Hello, World!
Hello, World!

...\FileToHex>filetohex.exe
=== File to Hex ===
File Path: HelloWorld.txt
Bytes in Hex: -1

00000000:0x48,0x65,0x6c,0x6c,0x6f,0x2c,0x20,0x57,0x6f,0x72,0x6c,0x64,0x21,0x0d,0x0a,0x48,
00000010:0x65,0x6c,0x6c,0x6f,0x2c,0x20,0x57,0x6f,0x72,0x6c,0x64,0x21,0x0d,0x0a,0x48,0x65,
00000020:0x6c,0x6c,0x6f,0x2c,0x20,0x57,0x6f,0x72,0x6c,0x64,0x21,0x0d,0x0a,
...\FileToHex>filetohex.exe
=== File to Hex ===
File Path: HelloWorld.txt
Bytes in Hex: 10

00000000:0x48,0x65,0x6c,0x6c,0x6f,0x2c,0x20,0x57,0x6f,0x72,0x6c,0x64,0x21,0x0d,0x0a,0x48,
...\FileToHex>
```