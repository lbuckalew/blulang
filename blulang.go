package blulang

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
)

type bluetoothctl struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

type response struct {
	kind string
	body string
}

func NewResponse(b []byte) *response {
	r := new(response)

	if m, e := regexp.Match("^Compile", b); m && (err == nil) {
		r.kind = "Info"
	} else if m, e = regexp.Match("[NEW]", b); m && (err == nil) {
		r.kind = "Device - New"
	} else if m, e = regexp.Match("[CHG]", b); m && (err == nil) {
		r.kind = "Device - Change"
	}

}

func (b *bluetoothctl) start() {

	b.cmd = exec.Command("bluetoothctl")

	b.stdin, _ = b.cmd.StdinPipe()
	b.stdout, _ = b.cmd.StdoutPipe()

	s := bufio.NewScanner(b.stdout)

	b.cmd.Start()

	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func (b *bluetoothctl) write(btCmd string) error {

	b.stdin.Close()
	b.stdin.Write([]byte(btCmd))
	b.stdin.Close()

	return nil
}

func (b *bluetoothctl) read(btMsg string) (string, error) {

	buff, error := ioutil.ReadAll(b.stdout)

	return string(buff), error
}

func (b *bluetoothctl) stop() error {

	b.write("exit")
	b.cmd.Wait()

	return nil
}

type Adapter struct {
	Name                                         string
	Powered, Discoverable, Pairable, Discovering bool
	shell                                        bluetoothctl
}

/* Attempts to connect to any known hosts
 */
func (a Adapter) ConnectKnown() error {

	// av_devs = adapter.startScan()
	// bd_devs = adapter.getBondedDevices()
	// if av_dev in bd_devs
	// connect
	// error
	return nil
}

/* Makes controller invisible and unpairable
 */
func (a Adapter) Lock() error {

	return nil
}

/* Allows for device to make pair request and connect
 */
func (a Adapter) FindDevice() error {

	return nil
}

func (a *Adapter) Init() error {
	// check for btshell
	a.shell.start()
	a.shell.Enable()

	// get attributes from shell
	a.Name = "Test name"
	a.Powered = false
	a.Discoverable = false
	a.Pairable = false
	a.Discovering = false

	return nil
}

func (a *Adapter) GetInfo() error {

}

func (a *Adapter) Enable() error {

	a.shell.write("power on")
	return nil
}

func (a *Adapter) Disable() error {

	a.shell.write("power off")
	return nil
}

func (a Adapter) isEnabled() bool {
	return true
}

func (a *Adapter) KillShell() error {

	a.stop()
	return nil
}

type Device struct {
	Name    string
	Address string
}
