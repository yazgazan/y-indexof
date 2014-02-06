
NAME=		"../init.tar"
FILES=		$(shell find . -maxdepth 1)

all:		tar

$(NAME):
	tar --exclude-vcs -cvf $@ $(FILES)

tar:		$(NAME)

.PHONY:		all tar

