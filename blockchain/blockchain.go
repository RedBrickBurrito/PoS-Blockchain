package Blockchain

import "sync"

var Blockchain []Block

var Announcements = make(chan string)

var Mutex = &sync.Mutex{}

var Validators = make(map[string]int)
