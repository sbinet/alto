alto
====

``alto`` is a naive CLI a la ``vagrant`` for managing ``StratusLab`` VMs.

## Installation

You need the ``StratusLab`` python client installed.
Once that's done:

```sh
$ go get github.com/sbinet/alto
```

and voila.


## Usage

```sh
$ cd somewhere
$ mkdir ubuntu-12.04-64b

# list vms
$ alto vm ls
Vm{Id=DdLv7aHgvg28oCSqpS8RfXVlv_o Tag="archlinux-64b"}
Vm{Id=EnMwelQXv5Z-8_B91ksTlzXoaRT Tag="ubuntu-12.04-32b"}
Vm{Id=AVijhrWdomWxMT5V34bqEPvArCB Tag="ubuntu-12.04-64b"}

# list disks
$ alto disk ls
[...]
:: DISK bf5aa557-1867-441d-bd04-4342fe379a02
	count: 0
	owner: binet
	tag: disk-ubuntu-12.04-64b
	size: 20
[...]

# create a new box (with defaults for CPU+RAM)
$ alto box add \
    -cpu=2 \
    -ram=2048 \
    box-ubuntu-12.04-64b \
        ubuntu-12.04-64b \
    disk-ubuntu-12.04-64b
    
# setup the current dir to host the box
$ alto init box-ubuntu-12.04-64b

# launch the box as a new instance on StratusLab
$ alto up

# connect to that instance
$ alto ssh

# shutdown instance
$ alto down
```

## Documentation

Well... it's embedded:

```sh
$ alto help
Usage:

	alto command [arguments]

The commands are:

    down        shutdown a (running) box on StratusLab
    init        create a box on StratusLab
    ssh         connect to a (running) box on StratusLab
    status      display the status of a box on StratusLab
    up          launch a box on StratusLab

    box         add/remove/edit boxes
    disk        add/remove/list persistent disks
    vm          add/remove/list VMs

Use "alto help [command]" for more information about a command.

Additional help topics:


Use "alto help [topic]" for more information about that topic.


$ alto help box
Usage:

	box command [arguments]

The commands are:

    add         add a box (VM+pdisk) to the repository of boxes
    ls          list boxes from the repository of boxes
    rm          rm a box from the repository of boxes


Use "box help [command]" for more information about a command.

Additional help topics:


Use "box help [topic]" for more information about that topic.

$ alto box help add
Usage: box add [options] <box-name> <vm-name> [<pdisk-name>]

add adds a box (VM+pdisk) on StratusLab.

ex:
 $ alto box add archlinux-64b my-archlinux-vm    # no disk
 $ alto box add archlinux-64b my-archlinux-vm my-archlinux-disk
 $ alto box add -cpu=4 -ram=2048 archlinux-64b my-archlinux-vm my-archlinux-disk

options:
  -cpu=1: number of CPUs for the VM
  -q=true: only print error and warning messages, all other output will be suppressed
  -ram=1024: amount of RAM for the VM (Mb)

```
