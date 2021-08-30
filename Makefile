all : build


build:
	g++ -o WeakPtr WeakPtr.cpp
	g++ -o MyString MyString.cpp

clean:
	rm -fr WeakPtr 
	rm -fr MyString 
