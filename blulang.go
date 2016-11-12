package blulang

import "os/exec"
import "io"
import "io/ioutil"
import "fmt"

type Adapter struct {
	Name string
	Powered, Discoverable, Pairable, Discovering bool
	btshell btShell
}

type Device struct {
	Name string
}







type btShell struct {
	stdin io.Writer
	stdout io.Reader
}

func (b *btShell) start() error {

	btCtl := exec.Command("bluetoothctl")

	b.stdin, _ = btCtl.StdinPipe()
	b.stdout, _ = btCtl.StdoutPipe()

	btCtl.Start()
	return nil
}

func (b *btShell) write(btCmd string) error {
	b.stdin.Write([]byte(btCmd))
	return nil
}

func (b *btShell) read(btMsg string) (string, error) {
	buff, error := ioutil.ReadAll(b.stdout)
	return string(buff), error
}

func (b *btShell) stop() error {
	b.write("exit")
	return nil
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
	blob, error := ioutil.ReadAll(a.btshell.stdout)
	fmt.Println("t")
	fmt.Println(string(blob))
	fmt.Println(error)
}

func (a *Adapter) GetAdapter() (Adapter, error) {
	// check for btshell
	a.btshell.start()
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


