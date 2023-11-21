package kitt

import (
	"fmt"
	"time"
)

func Animate(frames []string, delay time.Duration, done chan bool) {
	fmt.Printf("\r%s", "\033[?25l")
	defer fmt.Print("\033[?25h") // Show the cursor again

	reverse := false
	i := 0

	for {
		select {
		case <-done:
			fmt.Printf("\r%s", "\033[0m")
			return
		default:
			if !reverse {
				if i == len(frames) {
					reverse = true
				} else {
					fmt.Printf("\r%s", frames[i])
					time.Sleep(delay)
					i++
				}
			} else {
				if i == 4 {
					reverse = false
				} else {
					i--
					fmt.Printf("\r%s", frames[i])
					time.Sleep(delay)
				}
			}
		}
	}

}

func PrepareFrames() []string {
	black := "\033[30m"
	r0 := "\033[38;5;52m"
	r1 := "\033[38;5;88m"
	r2 := "\033[38;5;124m"
	r3 := "\033[31m" // or "\033[38;5;160m"
	r4 := "\033[38;5;196m"

	// Manually define the frames
	frames := []string{
		r0 + "»" + black + "»",
		r1 + "»" + r0 + "»" + black + "»",
		r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",

		r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		black + "»" + r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		black + "»" + black + "»" + r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		black + "»" + black + "»" + black + "»" + r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		black + "»" + black + "»" + black + "»" + black + "»" + r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + r0 + "»" + black + "»",
		black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + r1 + "»" + black + "»", //Start the end
		black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + r2 + "»" + black + "»",
		black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + r3 + "»" + black + "»",
		black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + black + "»" + r0 + "»" + r1 + "»" + r2 + "»" + r3 + "»" + r4 + "»" + black + "»",
	}
	return frames
}
