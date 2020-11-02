# gosho
![demo](demo.gif)

**gosho**, pronounced [goʃo], is a very simple minimalistic cli helper, which gives you a menu with hosts to which you can ssh to.

***
#### Requirements:
- *NIX OS
- the ssh command (ssh client installed and ssh in $PATH)
- gnu make (*optional)
- ssh config file with standard Host definition in your home directory, i.e.:


	`$: cat ~/.ssh/config`

			Host bastion
			    Hostname host.com
			    User donka

***
#### How to use
Clone the repo and run `make build && make install`, then:

	$: gosho

Then select the host you want to ssh into and you're done.

#### Authors
Viktor Penkov [vpenkoff@gmail.com](mailto:vpenkoff@gmail.com)

#### License
[MIT](https://gitlab.com/vpenkoff/gosho/-/blob/master/LICENSE)
