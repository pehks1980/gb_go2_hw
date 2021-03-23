package main

import (
	"flag"
	"fmt"
	"github.com/pehks1980/gb_go2_hw/hw8/fscan"
	Logger "github.com/pehks1980/gb_go2_hw/hw8/logger"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"sync/atomic"
)

var (
	// флаги
	deepScan          = flag.Bool("ds", false, "Deep scan check - check contents of the dub files")
	delDubs           = flag.Bool("del", false, "Delete dub files after scan")
	interactiveDelete = flag.Bool("i", false, "Interactive mode delete dub files after scan")
	// waitgroup
	wg = sync.WaitGroup{}
	// хештаблица структур файлов
	fileSet = fscan.NewRWSet()
	// счетчик гоу поточков
	goProcCounter int64 = 1
	// флаг запускает трассировку
	// cделать и поглядеть трассировку:
	// GOMAXPROCS=1 go run main.go > trace.out
	// go tool trace trace.out
	//
	TraceOn bool = false
)

// main основная функция работы утилиты
func main() {

	// trace code
	if TraceOn {
		trace.Start(os.Stderr)
		defer trace.Stop()
	}

	err := Logger.InitLoggers("log.txt")
	if err != nil {
		fmt.Printf("error opening log. exiting.")
		return
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage\n %s [options] <path>\n\nOptions:\n\t<path>\tpath for scan to start with (default 'working dir')\n ", os.Args[0])

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "\t-%v\t%v (default '%v')\n", f.Name, f.Usage, f.Value)
		})
		fmt.Fprintf(os.Stderr, "\nExamples:\n\n"+
			" '%s -ds -del -i /home/user/go'  \n"+
			"\t- find duplicates in /home/user/go using md5 hash calculation,\n"+
			"\tdelete files in interactive mode - one occurence by one (by pressing Y key).\n\n"+
			" '%s -del /home/user/go'  \n"+
			"\t- find duplicates in /home/user/go using file name and file size\n"+
			"\tdelete all files duplicates after scan is finished.\n\n", os.Args[0], os.Args[0])
	}
	flag.Parse()
	Logger.InfoLogger.Println("1 Starting the application...")

	path := flag.Args()
	if len(path) == 0 {
		pathArg, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		path = append(path, pathArg)
	}

	Logger.InfoLogger.Printf("Program started with pathDir=%s , Deep Scan is %t", path[0], *deepScan)
	wg.Add(1)
	go ScanDir(path[0])
	//ждем пока все перемножится
	wg.Wait()

	Logger.WarningLogger.Printf("scan created %d go procs...", goProcCounter)
	Logger.WarningLogger.Printf("scan found %d unique files with duplicates...", fileSet.FilesHaveDubs)
	fnum := 1
	for _, v := range fileSet.MM {
		if v.DubPaths != nil {
			switch *deepScan {
			case true:
				fmt.Printf("\n%d. File: %s Size: %d (B) Number of Dubs: %d md5: %s \n", fnum, v.FullPath, v.Filesize, len(v.DubPaths), v.FileHash)
			case false:
				fmt.Printf("\n%d. File: %s Size: %d (B) Number of Dubs: %d \n", fnum, v.FullPath, v.Filesize, len(v.DubPaths))
			}

			for i, dub := range v.DubPaths {
				fmt.Printf("%d.%d (DUB) File: %s \n", fnum, i+1, dub)
			}
			fnum++
			// обработка в интерактивном режиме
			if *delDubs && *interactiveDelete && fileSet.FilesHaveDubs != 0 {
			loop:
				for {
					fmt.Printf("\nWhich one you want to KEEP? Press number from 0 to %d, 0 - Keep original file (%s)\n", len(v.DubPaths), v.FullPath)
					var delPrompt int
					_, err := fmt.Scanf("%d", &delPrompt)
					if err != nil {
						// error here
						fmt.Printf("Error enter")
						return
					}
					switch {
					case delPrompt == 0:
						// delete all but original (0)
						for i, dub := range v.DubPaths {
							err := fscan.DeleteDup(dub)
							if err == nil {
								Logger.WarningLogger.Printf("%d. DUB File: %s DELETED", i, dub)
							}
						}
						break loop

					case delPrompt > 0 && delPrompt <= len(v.DubPaths):
						// keep selected dup, delete anything other
						for i, dub := range v.DubPaths {
							if i+1 != delPrompt {
								err := fscan.DeleteDup(dub)
								if err == nil {
									Logger.WarningLogger.Printf("%d. DUB File: %s DELETED", i, dub)
								}
							}

						}
						err := fscan.DeleteDup(v.FullPath)
						if err == nil {
							Logger.WarningLogger.Printf("%d. DUB File: %s DELETED", len(v.DubPaths)-1, v.FullPath)
						}
						break loop
					}

				}
			}
		}
	}
	// обработка в основном режиме
	if *delDubs && !*interactiveDelete && fileSet.FilesHaveDubs != 0 {
		var delPrompt string
		fmt.Println("\nConfirm delete of All duplicates 'Y' (then enter)?")

		_, err := fmt.Scanf("%s", &delPrompt)
		if err != nil {
			// error here
			fmt.Printf("Error enter")
			return
		}

		if delPrompt == "Y" || delPrompt == "y" {
			// delete all Dubs
			i := 0
			for _, v := range fileSet.MM {
				if v.DubPaths != nil {
					for _, dub := range v.DubPaths {
						err := fscan.DeleteDup(dub)
						if err == nil {
							Logger.WarningLogger.Printf("%d. DUB File: %s DELETED", i, dub)
							i++
						}
					}
				}
			}
		} else {
			Logger.WarningLogger.Printf("DUB Files Not DELETED")
		}
	}
	Logger.InfoLogger.Println("Finishing application...")
}
// ScanDir - принимает начальную папку и сканирует все подпапки
// для каждой подпапки запускает саму себя, выделяя новый поточек
func ScanDir(pathDir string) {
	defer wg.Done()
	dirs, err := fscan.IOReadDir(pathDir, fileSet, deepScan)
	if err != nil {
		Logger.ErrorLogger.Println("Error reading dirs", err)
		return
	}

	for _, dir := range dirs {
		wg.Add(1)
		atomic.AddInt64(&goProcCounter, 1)
		sDir := pathDir + "/" + dir
		go ScanDir(sDir)
	}

}
