package tests

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/chatterbox-irc/chatterbox/ircc/mock"
)

func TestConnection(t *testing.T) {
	t.Parallel()

	out, ircc, ircd, err := mock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer ircd.Close()

	err = ircc.WaitForConnection(5 * time.Second)
	if err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf(`{"type":"connection","status":"ok","msg":"%s"}`, ircc.Server)
	actual := out.String()

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected '%s' in '%s'", expected, actual)
	}

	ircc.Disconnect()
	mock.PollForEvent(fmt.Sprintf(`{"type":"quit","status":"ok","msg":"%s"}`, ircc.Server), out)
}
