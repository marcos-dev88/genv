package genv

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

const (
	notFindIndex       = -1
	indexMatchedValue  = 1
	nullMatch          = 0
	commentCharacter   = "#"
	separatorCharacter = "="
)

// New: This function define environment variables by file.
//
// * NOTE: .env file is defined by default, so don't need to send anything, if you want to use one or more files, you must define the all env files.
func New(envFiles ...string) error {
	var outErr error

	if len(envFiles) == 0 {
		return defineEnvs(".env")
	} else {
		for i := 0; i < len(envFiles); i++ {
			if err := defineEnvs(envFiles[i]); err != nil {
				outErr = err
				break
			}
		}
	}

	return outErr
}

func defineEnvs(filename string) error {

	file, err := os.Open(filename)

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			os.Stderr.WriteString(err.Error())
		}
	}(file)

	if err != nil {
		return err
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		indexComment := strings.Index(sc.Text(), commentCharacter)
		if indexComment != notFindIndex && len(strings.TrimSpace(sc.Text()[:indexComment])) == nullMatch {
			continue
		}
		envEqualSign := strings.Index(sc.Text(), separatorCharacter)
		if envEqualSign != notFindIndex {
			envMatchKey := sc.Text()[:envEqualSign]
			envMatchValue := sc.Text()[envEqualSign+indexMatchedValue:]
			if len(envMatchKey) != nullMatch || len(envMatchValue) != nullMatch {
				err := os.Setenv(envMatchKey, envMatchValue)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// NewFast: This function define environment variables by file using go routines.
//
// * NOTE: .env file is defined by default, so don't need to send anything, if you want to use one or more files, you must define the all env files.
func NewFast(envFiles ...string) error {

	var (
		reading   = make(chan bool, len(envFiles))
		errDefEnv = make(chan error, len(envFiles))
		outErr    error
		wg        = new(sync.WaitGroup)
	)

	if len(envFiles) == 0 {
		wg.Add(1)
		go defineEnvsFast(".env", reading, errDefEnv, wg)
	} else {
		for i := 0; i < len(envFiles); i++ {
			wg.Add(1)
			go defineEnvsFast(envFiles[i], reading, errDefEnv, wg)
		}
	}

	go func() {
		wg.Wait()
		close(reading)
		close(errDefEnv)
	}()

	for range reading {
		if err := <-errDefEnv; err != nil {
			outErr = err
			break
		}
	}

	return outErr
}

func defineEnvsFast(filename string, reading chan bool, chErr chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	reading <- true

	file, err := os.Open(filename)

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			os.Stderr.WriteString(err.Error())
		}
	}(file)

	if err != nil {
		chErr <- err
		return
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		indexComment := strings.Index(sc.Text(), commentCharacter)
		if indexComment != notFindIndex && len(strings.TrimSpace(sc.Text()[:indexComment])) == nullMatch {
			continue
		}
		envEqualSign := strings.Index(sc.Text(), separatorCharacter)
		if envEqualSign != notFindIndex {
			envMatchKey := sc.Text()[:envEqualSign]
			envMatchValue := sc.Text()[envEqualSign+indexMatchedValue:]
			if len(envMatchKey) != nullMatch || len(envMatchValue) != nullMatch {
				err := os.Setenv(envMatchKey, envMatchValue)
				if err != nil {
					chErr <- err
					break
				}
			}
		}
	}
}
