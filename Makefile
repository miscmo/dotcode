all : build


build:
	g++ -o weak_ptr weak_ptr.cpp

clean:
	rm -fr weak_ptr
