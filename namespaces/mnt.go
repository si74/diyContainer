package main

// (4)
// MNT namespace - gives a process it's own mount table
// can mount/unmount directories, swap out filesystem container sees
import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}
}

func run() {

	// link to currently running process
	// http://unix.stackexchange.com/questions/333225/which-process-is-proc-self-for
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...) // link to currently running process
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as pid%v\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//must(syscall.Chroot("/home/rootfs")) //where to get an extra filesystem
	//must(os.Chdir("/"))
	//must(syscall.Mount("proc", "proc", "proc", 0, ""))

	must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	must(os.MkdirAll("rootfs/oldrootfs", 0700))
	must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	must(os.Chdir("/"))

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
