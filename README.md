# ppmap
A simple scanner/exploitation tool written in GO which automatically exploits known and existing gadgets (checks for specific variables in the global context) to perform XSS via Prototype Pollution. NOTE: The program only exploits known gadgets, but does not cover code analysis or any advanced Prototype Pollution exploitation, which may include custom gadgets.

## Requirements
Make sure to have Chromium installed.

## Installation
- Run the following command to clone the repo: 
 ```bash
git clone https://github.com/MrChiz/ppmap
cd ppmap
go install .
sudo cp ~/go/bin/ppmap /usr/local/bin
cd ~
 ```
 - 
That's it. Enjoy using ppmap!

## Usage

Using the program is very simple, you can either:
```ppmap scan -u https://target.com/ppmap/test.html```

## Demo

![Screenshot_2024-09-04_05-55-21](https://github.com/user-attachments/assets/c44c9fab-1549-4352-bc4f-5b4a90fb8251)
