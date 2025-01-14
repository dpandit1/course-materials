package main

import(
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

	var nfiles, nbytes int64
	for _, root := range roots{
		nf, nb := walkDir(root) //call the walkDir function
		nfiles += nf
		nbytes += nb
	}

	fmt.Println("Total time taken: ", time.Since(now))
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64){
	fmt.Printf("%d files %.4 GB\n", nfiles, float64(nbytes)/1e9)
}

//walkDir uses recursion to scan each file and returns the number and size of the total files
func walkDir(dir string) (numFiles int64, size int64) {
	time.Sleep(100 * time.Millisecond)
	for _, entry:= range dirents(dir) { 
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			nf, fs := walkDir(subdir)
			numFiles += nf
			size += fs
		} else{
			numFiles++
			size += entry.Size()
		}
	}
	return
}

//dirents gives the directories and the files present in the initial directory
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	} 
	return entries
}