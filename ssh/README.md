# SSH Client
SSH client execute command on remote host.

```golang
t := &ssh.Task{
    Username: "ubuntu",
    AuthMethods: []ssh.AuthMethod{
        {
            Type: ssh.AuthByPassword,
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