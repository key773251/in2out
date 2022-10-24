# in2out
Simple file format converter between jsonnet, json, and yaml. 

</br>

## Input File Types
The following input file types are currently supported:
| File Extension |
| -------------- |
|    *.json      |
|  *.jsonnet     |
|    *.yaml      |

</br>

## Output File Types
The following output file types are currently supported:
| File Extension |
| -------------- |
|    *.json      |
|    *.yaml      |

</br>

# Building
```
$ make 
```

</br>

# Usage
```
$ ./build/in2out
in2out version:0.1.0
  -d    Enable debugging output
  -e string
        External Variables for jsonnet substitution
  -i string
        Input file path (Required)
  -o string
        Output file path (Required)
```

</br>

# Contributing
If you already can't tell from the code, I am a complete novice when it comes to using go. 
I am open to any feedback/collaboration that anyone would like to provide :)