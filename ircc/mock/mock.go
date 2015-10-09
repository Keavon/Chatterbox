package mock

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chatterbox-irc/chatterbox/ircc/irc"
)

var (
	port      = 10000
	portMutex = sync.Mutex{}
)

// IRCD holds a connection to a ircd
type IRCD []string

// Close stops an test ircd.
func (m IRCD) Close() {
	err := exec.Command(m[0], m[1:]...).Run()
	fmt.Println(err)
}

// New returns a mock IRC setup.
func New() (*bytes.Buffer, *irc.IRC, *IRCD, error) {
	reader := bytes.Buffer{}

	portMutex.Lock()
	p := strconv.Itoa(port)
	port++
	portMutex.Unlock()

	cmd := exec.Command("docker", "run", "--name=cbx-ircd-test-"+p, "-d", "-p", p+":6667", "xena/elemental-ircd")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err := cmd.Run()

	if err != nil {
		return nil, nil, nil, err
	}

	// Need to wait for the ircd to startup. Otherwise we get connection errors.
	time.Sleep(50 * time.Millisecond)

	ircc, err := irc.New("cbx", "test", "localhost:"+p, "", false, &reader)

	if err != nil {
		return nil, nil, nil, err
	}

	ircd := IRCD{"docker", "rm", "-f", "cbx-ircd-test-" + p}

	return &reader, ircc, &ircd, nil
}

// PollForEvent polls for an event and blocks until it gets the event or times out.
func PollForEvent(event string, out *bytes.Buffer) error {
	start := time.Now()
	timeout := 2 * time.Second
	recieved := ""

	for {
		recieved += out.String()

		if strings.Contains(recieved, event) {
			return nil
		}

		if time.Since(start) > timeout {
			return fmt.Errorf("Expected '%s' in '%s'", event, recieved)
		}
	}
}
