src = $(wildcard *.cpp)
obj = $(src:.c=.o)

LDFLAGS =

ds: $(obj)
	$(CXX) -std=c++14 -o $@ $^ $(LDFLAGS)

.PHONY: clean
clean:
	rm -f $(obj) ds