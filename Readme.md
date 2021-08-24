# SM

A State Machine

Uses Stateless[https://github.com/qmuntal/stateless] to implement simple State Machine.

# One day...

You will be able to add a state-machine member to your struct and allow it to be modified by a state-machine.

# Build
```
go build
```

# Running
 
 Right now it outputs a graphviz definition of the statemachine in `dot` language. To generate the png you can use
 
 ```
 ./sm > a.dot
 dot -Tpng a.dot > a.png
 
 ```
![a](https://user-images.githubusercontent.com/331871/130680480-f100a21d-40ab-4709-9e89-b826f5e8c763.png)
