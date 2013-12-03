
NAME=		"../init.tar"
FILES=		$(shell find . -depth 1)

all:		tar

$(NAME):
	tar -cvf $@ $(FILES)

tar:		$(NAME)

.PHONY:		all tar

