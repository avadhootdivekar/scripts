# scripts
This repository hosts some simple scripts. Intention is to let people pick up some readily available scripts which they can use in their projects.

## Number conversion script. 
This script changes number from one base to other base. It also recognizes and formats output based on delimeter and prefix requests. 
User can also provide the files for input and output. 
Can be used to convert between decimal,binary, octal , hexadecimal etc upto bsae 36. 

## Usage

```
python3 num_conv.py  -d " " --out-delimeter="," --in-base=10 --out-base="2"  --out-prefix="0b" --s "165 170 255 1024" 
```

## helper.sh
This is a helper script which anyone can easily pull and start using. 
Use following command to pull and load the helper script. 
#### WARNING : Please understand what you rae doing. Following command will download a file and source it.
```
wget https://raw.githubusercontent.com/avadhootdivekar/scripts/refs/heads/master/bash/.helper.sh -O ~/.helper.sh && source ~/.helper.sh
```
