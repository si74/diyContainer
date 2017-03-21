# diyContainer

Containers: What, Why, How?

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

`sudo apt-get install golang`

`source env.sh`

## DIY

(1) Run the basic container file:

`go run basic.go run echo "Hello"` //Will print echo

`go run basic.go run /bin/bash` //Process will open a bash shell
`hostname <FAKENAME>` //Can set hostname from our program - NOT GOOD!!!

(2) Run uts.go - we've now restricted hostname access using the UTS namespace:

`sudo go run basic.go `
