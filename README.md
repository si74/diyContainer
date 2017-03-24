# diyContainer

Containers: What, Why, How?

This repo is the companion to this [talk](https://github.com/si74/diyContainer).

## Preparation

Many of the examples here will work only a Linux based distribution.
As a result, I recommend using the vagrant manager along with the VirtualBox hypervisor.

Download vagrant and virtual box on the MacOSX if you have not already done so:

`brew cask install virtualbox`

`brew cask install vagrant`

The vagrant file is in the repository so all you need to do is:

`vagrant up`

Note: as per the docs, the current working directory is shared in the /vagrant folder of the guest.

Install go and set the $GOPATH:

`sudo apt-get install golang
source env.sh`

## DIY

1. Run the basic container file, `basic.go`:

   ```
   go run basic.go run echo "Hello" //Will print echo  
   go run basic.go run /bin/bash //Process will open a bash shell  
   hostname <FAKENAME> //Can set hostname from our program - NOT GOOD!!!  
   exit  
   hostname
   ```

   We'll see that our hostname is now the FAKENAME. Not good!

2. Run `uts.go` - we've now restricted hostname access using the UTS namespace:

   ```  
   sudo go run namespaces/uts.go run /bin/bash  
   hostname <FAKENAME>  
   exit  
   hostname
   ```  

   Using the UTS namespace, we've isolated our host hostnames, etc. We cannot affect/change them from within the "container" process.

3. Run `pid.go` - we've now restricted PID access using the PID namespace:

   ```  
   sudo go run namespaces/pid.go run /bin/bash  
   ps  
   ```  

   Hmm. Why are we seeing the processes from our host system? Let's try to fork a child process instead (step 4).

4. Run `pid_1.go`

   ```
   sudo go run namespaces/pid_1.go run /bin/bash  
   ps  
   ```

   Good news! We see that our child process has been assigned a PID of 1! Unfortunately when we use `ps` we see that all the processes on the host machine. Why?! Well, it turns out that ps checks the /proc directory to determine running processes.

5. Run `mnt.go`

   So let's try to use the MNT namespace as well (which gives the process it's own mount table and enables one to swap out the filesystem a container sees.).

   ```
   sudo go run namespaces/mnt.go run /bin/bash
   ps
   ```

   Uh oh! What happened. Though the process is in the MNT namespace, we're still mounting the root filesystem. In order to see only local processes, we'd also need to swap out a root filesystem...

## Next Steps

1. Add a filesystem

   Note the file `namespaces/pid_1.go`. This file demo's two ways of mounting different root filesystem. However, I did not have enough time to actually finish the demo (i.e. have a usable extra filesystem in our vagrant-controlled VM).

2. Add cgroups

   This demo primarily used namespaces for process isolation. However, for the fully
   containerized experience, we will need to leverage cgroups for resource sharing.

## Sources

[Build a Container In Less Than 100 Lines of Go](https://www.infoq.com/articles/build-a-container-golang)

[What is a Container, really?](https://www.youtube.com/watch?v=HPuvDm8IC-4)

[A tutorial for isolating your Namespace](https://www.toptal.com/linux/separation-anxiety-isolating-your-system-with-linux-namespaces)

[Unprivileged Containers in Go (a 4-part series)](https://lk4d4.darth.io/)
