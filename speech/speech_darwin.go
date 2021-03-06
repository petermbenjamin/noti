package speech

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/variadico/noti"
)

// Notify speaks a notification using NSSpeechSynthesizer.
func Notify(n noti.Params) error {
	voice := n.Config.Get(voiceEnv)
	if voice == "" {
		voice = "Alex"
	}
	text := fmt.Sprintf("%s %s", n.Title, n.Message)

	cmd := exec.Command("say", "-v", voice, text)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("speech: %s", err)
	}

	return nil
}
