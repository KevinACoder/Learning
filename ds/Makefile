src = $(wildcard *.c)
obj = $(src:.c=.o)

LDFLAGS = 

ds: $(obj)
	$(CC) -std=c99 -g -o $@ $^ $(LDFLAGS)

.PHONY: clean
clean:
	rm -f $(obj) ds