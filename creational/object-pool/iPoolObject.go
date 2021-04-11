package main

type iPoolObject interface {
	getID() string //This is any id which can be used to compare two different pool objects
}
