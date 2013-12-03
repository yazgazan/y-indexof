
NAME=		y-indexof

MAKE=		gmake
RM=		rm -vf

INITDIR=	init_base
INITFILE=	init.tar
TESTDIR=	test

all:		make_init build install test

make_init:
	$(MAKE) -C $(INITDIR)

build:
	go build

install:
	go install

clean:
	$(RM) $(INITFILE)
	$(RM) $(NAME)
	$(RM) -r $(TESTDIR)

test:
	sh test.sh

re:		clean all

.PHONY:		all make_init build install clean test

