# tools_linestamp

Filter (Pipe) that adds a timestamp to each line 

## Options
```
Usage of tools_linestamp:
  -n	Prefix lines with a 4 digit line number
  -help Display this text
```
## Examples

cat someTextFie.txt | tools_linestamp

```
2018.04.05 23:23:03.581 # tools_linestamp
2018.04.05 23:23:03.581 
2018.04.05 23:23:03.581 Filter (Pipe) that adds a timestamp to each line 
2018.04.05 23:23:03.581 
2018.04.05 23:23:03.581 ## Options
2018.04.05 23:23:03.581 Usage of tools_linestamp:
2018.04.05 23:23:03.581   -n	Prefix lines with a 4 digit line number
```

cat someTextFie.txt | tools_linestamp -n

```
0001:2018.04.05 23:23:26.491 # tools_linestamp
0002:2018.04.05 23:23:26.491 
0003:2018.04.05 23:23:26.491 Filter (Pipe) that adds a timestamp to each line 
0004:2018.04.05 23:23:26.492 
0005:2018.04.05 23:23:26.492 ## Options
0006:2018.04.05 23:23:26.492 Usage of tools_linestamp:
0007:2018.04.05 23:23:26.492   -n	Prefix lines with a 4 digit line number
```
tools_linestamp -help

```
Usage of tools_linestamp:
  -n	Prefix lines with a 4 digit line number
```
