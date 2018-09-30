package ssh

import "time"

type Task struct {
	Username    string        `json:"username"`
	AuthMethods []AuthMethod  `json:"auth_methods"`
	Host        string        `json:"host"`
	Script      string        `json:"script"`
	Timeout     time.Duration `json:"timeout"`
}

type AuthMethod struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
