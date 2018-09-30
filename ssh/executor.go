package ssh

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

const (
	AuthByPassword  = "password"
	AuthByPublicKey = "publickey"
)

// Execute the SSH task on remote host.
func (t *Task) Execute() (string, error) {
	config := new(ssh.ClientConfig)
	config.SetDefaults()
	config.User = t.Username
	config.Timeout = t.Timeout
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	var authMethods []ssh.AuthMethod
	for _, m := range t.AuthMethods {
		if m.Type == AuthByPassword {
			authMethods = append(authMethods, ssh.Password(m.Content))
		} else if m.Type == AuthByPublicKey {
			pk, err := publicKey([]byte(m.Content))
			if err != nil {
				return "", err
			}
			authMethods = []ssh.AuthMethod{pk}
		}
	}

	config.Auth = authMethods
	if len(authMethods) < 1 {
		return "", fmt.Errorf("none available auth methods")
	}

	client, err := ssh.Dial("tcp", t.Host, config)
	if err != nil {
		return "", err
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run(t.Script); err != nil {
		return "", fmt.Errorf("failed to run: %v", err)
	}
	return b.String(), nil
}

// publicKey reads the private key's content.
func publicKey(content []byte) (ssh.AuthMethod, error) {
	key, err := ssh.ParsePrivateKey(content)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}
