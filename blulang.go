package blulang

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

type bluetoothctl struct {
	cmd    *exec.Cmd
	stdin  io.Writer
	stdout io.Reader
}

func (b *bluetoothctl) start() error {
	var err error

	b.cmd = exec.Command("bluetoothctl")

	b.stdin, err = b.cmd.StdinPipe()
	if err != nil {
		return err
	}

	b.stdout, err = b.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	return b.cmd.Start()
}

func (b *bluetoothctl) write(btCmd string) error {

	b.stdin.Write([]byte(btCmd))

	return nil
}

func (b *bluetoothctl) read(btMsg string) (string, error) {

	buff, error := ioutil.ReadAll(b.stdout)

	return string(buff), error
}

func (b *bluetoothctl) stop() error {

	b.cmd.Wait()

	b.write("exit")

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

func (a *Adapter) Testmsg() {

	blob, error := ioutil.ReadAll(a.shell.stdout)

	fmt.Println("t")
	fmt.Println(string(blob))
	fmt.Println(error)
}

func (a *Adapter) GetAdapter() (Adapter, error) {

	// check for btshell
	a.shell.start()

	// enable

	// get attributes from shell
	a.Name = "Test name"
	a.Powered = false
	a.Discoverable = false
	a.Pairable = false
	a.Discovering = false

	return *a, nil
}

func (a *Adapter) enable() error {

	return nil
}

func (a *Adapter) disable() error {

	return nil
}

type Device struct {
	Name    string
	Address string
}
