# HyperDbg Script

This repo is a collection of useful HyperDbg scripts. HyperDbg uses the "**.ds**" extension (which is stands for **D**ebugger **S**cript).

You can use the '[.script](https://docs.hyperdbg.org/commands/meta-commands/.script)' command to run these scripts. For more examples, take a look at [here](https://docs.hyperdbg.org/commands/scripting-language/debugger-script).

## Usage

You can use scripts in the HyperDbg environment:

```
HyperDbg> .script c:\users\sina\desktop\script.ds
```

or you can directly run them :
```
C:\Users\sina\Desktop\HyperDbg>hyperdbg-cli.exe --script c:\users\sina\desktop\script.ds
```
## Description

### Basics
- **hello-world.ds**: The *Hello World!* script.

### DFIR (Digital Forensics and Incident Response)
- **process-behavior-logger.ds**: Gathering information about different behavior of a process like system calls, kernel memory allocations, CPUIDs, etc.

### Network
- **all-connections-ip-port.ds**: Creates a log from all of the network accesses (IP address and port number) for user-mode applications.
- **process-specific-connections-ip-port.ds**: Creates process-specific logs of the network accesses (IP address and port number).

### Memory
- **user-mode-memory-allocations.ds**: Creates a log from memory allocations of a user-mode process (e.g., mallocs).

## Contributing
Pull requests are super welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
