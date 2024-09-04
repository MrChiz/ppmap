# ppmap
A simple scanner/exploitation tool written in GO which automatically exploits known and existing gadgets (checks for specific variables in the global context) to perform XSS via Prototype Pollution. NOTE: The program only exploits known gadgets, but does not cover code analysis or any advanced Prototype Pollution exploitation, which may include custom gadgets.

## Requirements
Make sure to have Chromium installed.

## Installation
- Run the following command to clone the repo: 
 ```bash
go install github.com/MrChiz/ppmap/v2@latest
sudo cp ~/go/bin/ppmap /usr/local/bin
 ```
 - 
That's it. Enjoy using ppmap!

## Usage

Using the program is very simple, you can either:
```ppmap scan -u https://target.com/ppmap/test.html```

## Demo

![Screenshot_2024-09-04_11-18-30](https://github.com/user-attachments/assets/da8f540c-ea68-4378-b40c-7469e0fe3b37)

