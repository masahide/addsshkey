package opcli

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

const (
	InitBufSize = 2048
)

type Op struct {
	*Config
	ID            string
	twoFactorFunc func() string
	cmdIn         []byte
	cmdOut        []byte
	cmdErr        []byte
	cmdInBuf      *bytes.Buffer
	cmdOutBuf     *bytes.Buffer
	cmdErrBuf     *bytes.Buffer
	*exec.Cmd
	session       string
	sessionExpiry time.Time
	errEvent      chan error
}

func NewOp(c *Config, id string) *Op {
	o := &Op{}
	o.ID = id
	o.Config = c
	o.cmdIn = make([]byte, 0, InitBufSize)
	o.cmdOut = make([]byte, 0, InitBufSize)
	o.cmdErr = make([]byte, 0, InitBufSize)
	o.cmdInBuf = bytes.NewBuffer(o.cmdIn)
	o.cmdOutBuf = bytes.NewBuffer(o.cmdOut)
	o.cmdErrBuf = bytes.NewBuffer(o.cmdErr)
	return o
}

func (o *Op) SetSession(sess string) {
	o.session = sess
	o.sessionExpiry = time.Now().Add(45 * time.Minute)
}
func (o *Op) resetBuf() {
	o.Cmd = &exec.Cmd{}
	o.Cmd.Path = o.CmdPath
	o.cmdInBuf.Reset()
	o.cmdOutBuf.Reset()
	o.cmdErrBuf.Reset()
	o.Cmd.Stdin = o.cmdInBuf
	o.Cmd.Stderr = o.cmdErrBuf
	o.Cmd.Stdout = o.cmdOutBuf
	o.errEvent = make(chan error)
}

type stdoutWork struct {
	Event chan string
}

func newStdOutWork() *stdoutWork {
	return &stdoutWork{
		Event: make(chan string, 1),
	}
}

func (s *stdoutWork) Write(p []byte) (n int, err error) {
	str := string(p)
	s.Event <- str
	return len(p), nil
}

type stdInWork struct {
	Event chan string
	data  string
}

func newStdInWork() *stdInWork {
	return &stdInWork{
		Event: make(chan string, 1),
	}
}

func (s *stdInWork) Read(p []byte) (n int, err error) {
	if len(s.data) > 0 {
		return s.send(p)
	}
	var ok bool
	select {
	case s.data, ok = <-s.Event:
		if !ok {
			return 0, io.EOF
		}
		return s.send(p)
	}
}

func (s *stdInWork) send(p []byte) (n int, err error) {
	l := len(s.data)
	if l >= len(p) {
		copy(p, []byte(s.data[:len(p)]))
		s.data = s.data[len(p):]
		return len(p), nil
	}
	copy(p, []byte(s.data))
	s.data = s.data[:0]
	return l, nil
}
func (o *Op) runWorker() {
	o.errEvent <- o.Run()
	close(o.errEvent)
}

func (o *Op) twofactor() string {
	return o.twofactor()
}
func (o *Op) AuthCmdHandler(info OpAccountInfo, pass string, twoFactorFunc func() string) (string, error) {
	o.twoFactorFunc = twoFactorFunc
	stdIn := newStdInWork()
	stdErr := newStdOutWork()
	stdOut := newStdOutWork()
	o.resetBuf()
	stdIn.data = pass + "\n"
	o.Cmd.Stdin = stdIn
	o.Cmd.Stderr = stdErr
	o.Cmd.Stdout = stdOut
	o.Cmd.Path = o.CmdPath
	o.Cmd.Args = []string{o.CmdPath, "signin", info.URL, info.Email, info.AccountKey, "--raw", "--config", o.opConfigPath}
	//log.Printf("cmd:%v", o.Cmd.Args)
	go o.runWorker()
	var out string
	var err error
	var errStr string
L1:
	for {
		select {
		case err = <-o.errEvent:
			//log.Printf("<-o.errEvent:%v", err)
			break L1
		case s := <-stdErr.Event:
			//log.Printf("<-stdErr.Event:%v", s)
			if strings.Contains(s, "six-digit authentication code") {
				stdIn.Event <- o.twofactor() + "\n"
				continue
			}
			if strings.Contains(s, "configuration at non-standard location") {
				continue
			}
			if len(errStr) > 0 {
				errStr += "\n"
			}
			errStr += s
			if strings.Contains(s, "ERROR") {
				close(stdIn.Event)
			}
		case out = <-stdOut.Event:
			//log.Printf("out:%s", out)
			close(stdIn.Event)
		}
	}
	//log.Println("end select loop")
	close(stdOut.Event)
	close(stdErr.Event)
	if err != nil {
		return "", fmt.Errorf("Unauthorized. error message:%s, err:%s", errStr, err)
	}
	l := strings.Index(out, "\n")
	if l == -1 {
		return out, nil
	}
	return out[:l], nil

}

type OpCli struct {
	*Config
	Ops []*Op
}

func NewOpCli(c *Config) *OpCli {
	ops := make([]*Op, len(c.AccountIDs))
	for i, id := range c.AccountIDs {
		ops[i] = NewOp(c, id)
	}
	return &OpCli{
		Config: c,
		Ops:    ops,
	}

}

func (opc *OpCli) GetAccountInfo() ([]OpAccountInfo, error) {
	conf, err := loadOpConfig(opc.baseDirPath)
	if err != nil {
		return []OpAccountInfo{}, nil
	}
	res := make([]OpAccountInfo, len(conf.Accounts))
	for i, a := range conf.Accounts {
		res[i] = OpAccountInfo{
			ID:         a.UserUUID,
			URL:        a.URL,
			Email:      a.Email,
			AccountKey: a.AccountKey,
		}
	}
	return res, nil
}

func (opc *OpCli) GetUserID(accountKey string) (string, error) {
	conf, err := loadOpConfig(opc.baseDirPath)
	if err != nil {
		return "", err
	}
	for _, a := range conf.Accounts {
		if a.AccountKey == accountKey {
			return a.UserUUID, nil
		}
	}
	return "", errors.New("not found Account")
}

func (opc *OpCli) addOp(o *Op) {
	opc.AccountIDs = append(opc.AccountIDs, o.ID)
	opc.Ops = append(opc.Ops, o)
}

func (opc *OpCli) AddAccount(info OpAccountInfo, pass string, twoFactorFunc func() string) error {
	o := NewOp(opc.Config, "")
	session, err := o.AuthCmdHandler(info, pass, twoFactorFunc)
	if err != nil {
		return err
	}
	o.SetSession(session)
	id, err := opc.GetUserID(info.AccountKey)
	if err != nil {
		return err
	}
	o.ID = id
	opc.addOp(o)
	return nil

}

func (opc *OpCli) DeleteAccount(o *Op) {
	ids := make([]string, 0, len(opc.AccountIDs))
	for _, id := range opc.AccountIDs {
		if o.ID == id {
			continue
		}
		ids = append(ids, id)
	}
	opc.AccountIDs = ids
	ops := make([]*Op, 0, len(opc.Ops))
	for _, op := range opc.Ops {
		if op.ID == o.ID {
			continue
		}
		ops = append(ops, op)
	}
	opc.Ops = ops
}
