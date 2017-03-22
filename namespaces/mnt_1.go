package main

// (4)
// TODO(sneha) - doesn't entirely work yet
// MNT namespace - gives a process it's own mount table + new filesystem
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

	// TODO(si74) - Need to add own filesystem
	// OPTION 1: get root filesystem, change directly path, and mount empty proc
	// must(syscall.Chroot("/home/rootfs"))
	// must(os.Chdir("/"))
	// must(syscall.Mount("proc", "proc", "proc", 0, ""))

	// OPTION 2: Mount root file system and pivot to root
	// must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	// must(os.MkdirAll("rootfs/oldrootfs", 0700))
	// must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	// must(os.Chdir("/"))

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
