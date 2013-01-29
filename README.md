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
