Cross-platform development with Go can be achieved by using the GOOS and GOARCH environment variables. The GOOS environment variable specifies the operating system you want to target, and the GOARCH environment variable specifies your target architecture.

To compile code for Linux, you would set the GOOS environment variable to linux and the GOARCH environment variable to amd64.

`GOOS=linux GOARCH=amd64 go build`

This command will compile the code for Linux. 

The command to compile these program are respectively:

`GOOS=windows go build -o app.exe`
`GOOS=linux go build -o app`

When we execute the app within a Linux environment, it should print This is Linux!. Meanwhile, on a Windows system, running app.exe will display This is Windows!. 
