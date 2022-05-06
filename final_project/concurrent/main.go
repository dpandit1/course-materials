package main

import(
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main(){
	flag.Parse()
	// check the directory
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	now := time.Now()

	//initiate a channel of type int64
	fileSizes := make(chan int64)
	var n sync.WaitGroup //used for multiple go routines
	for _, root := range roots{
		n.Add(1)
		go walkDir(root, &n, fileSizes) //go routine
	}

	go func(){
		n.Wait()
		close(fileSizes)
	}()
		
	var nfiles, nbytes int64
	for size := range fileSizes{		
		nfiles ++
		nbytes += size
	}

	fmt.Println("Total time taken: ", time.Since(now))
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64){
	fmt.Printf("%d files %.4f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	time.Sleep(100 * time.Millisecond)
	defer n.Done() //indicates the end of the go routines
	for _, entry:= range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes) //go routine
		} else{
			fileSizes <- entry.Size()
		}
	}	
}

//semaphore is to limit concurrncy in the dirents function
var sema = make (chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <-struct{}{}
	defer func() { <- sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	} 
	return entries
}