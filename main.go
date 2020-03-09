package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"gopkg.in/ini.v1"
	"github.com/briandowns/spinner"
)

func runCommand(paths []string, imgResolution string, extName string) string {
	wg := &sync.WaitGroup{}

	flag := 0 // 标记为，初始化为0

	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			//log.Println(runtime.NumGoroutine()) // goroutine的数
			defer wg.Done()
			ext := strings.LastIndex(path, ".") // . 的index位置
			// PDF判断.后面是否是pdf，如果是则处理
			if path[ext+1:] == extName {
				flag = 1
				pdfDir := filepath.Dir(path)            // PDF文件的目录路径
				filename := getFileNameWithoutExt(path) // 文件名称
				saveDir := pdfDir + "/" + filename      // 图片保存的路径
				err1 := os.Mkdir(saveDir, 0777)         // 创建目录
				if err1 != nil {
					panic(err1)
				}
				// mutoolを叩いて画像出力
				imageFile := saveDir + "/" + "p%04d.png"
				cmd := exec.Command("mutool.exe", "draw", "-o", imageFile, "-r", imgResolution, path)
				err2 := cmd.Run()
				if err2 != nil {
					panic(err2)
				}
				fmt.Println("\n", path) // 处理文件名称
				// 运行的状态
				state := cmd.ProcessState
				fmt.Printf("  %s\n", state.String())               // 状态
				fmt.Printf("    Pid: %d\n", state.Pid())           // pid
				fmt.Printf("    System: %v\n", state.SystemTime()) // system时间
				fmt.Printf("    User: %v\n", state.UserTime())     // user时间
			}
		}(path)
	}
	wg.Wait()

	// PDF不存在
	if flag == 0 {
		fmt.Println("PDF file is missing.")
		os.Exit(1)
	}

	return "\nDone."
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func main() {

	cfg, err := ini.Load("pdf2png.ini")
    if err != nil {
        fmt.Printf("Fail to read config file: %v", err)
        os.Exit(1)
	}

	dir := cfg.Section("").Key("dir").String()
	imgResolution := cfg.Section("").Key("imgResolution").Validate(func(in string) string{
		if len(in) == 0 {
			return "300"
		}
		return in
	})
	extName := cfg.Section("").Key("extname").Validate(func(in string) string{
		if len(in) == 0 {
			return "pdf"
		}
		return in
	})

	paths := dirwalk(dir) // 递归目录

	// 判断目录列表是否为空
	if paths == nil {
		fmt.Println("File is missing.")
		os.Exit(1)
	}

	fmt.Println("Processing...")

	//设置可视化窗口
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Color("green")
	s.Start()

	// 启动转换
	result := runCommand(paths, imgResolution, extName)

	// 结束可视化窗口
	s.Stop() 

	fmt.Println(result)
}
