# SSH Client
SSH client execute command on remote host.

## Use password to execute
```golang
t := &ssh.Task{
    Username: "ubuntu",
    AuthMethods: []ssh.AuthMethod{
        {
            Type:    ssh.AuthByPassword,
            Content: "pass",
        },
    },
    Host: "ip:22",
    Script: `
ls /
`,
}
result, err := t.Execute()
if err != nil {
    panic(err)
}

fmt.Println(result)
```

## Use `ssh` command line
Sometimes you're not able to ssh the remote host by password or public key but only able to ssh by
ssh client in terminal, so there's another way to create a process calling the shell command `ssh`.

```golang
t := &ssh.Task{
    Username: "ubuntu",
    UseSSHCommand: true,
    Host: "10.0.0.11",
    Script: `
ls /
`,
}
result, err := t.Execute()
if err != nil {
    panic(err)
}

fmt.Println(result)
```
