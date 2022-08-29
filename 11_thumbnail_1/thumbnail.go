package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	// filenames := []string{"img1.jpg", "img2.jpg", "img3.jpg", "img4.jpg", "img5.jpg", "img6.jpg", "img7.jpg"}
	// makeThumbnails1(filenames)
	// makeThumbnails2(filenames)

	filenames := make(chan string, 7)
	filenames <- "img1.jpg"
	filenames <- "img2.jpg"
	filenames <- "img3.jpg"
	filenames <- "img4.jpg"
	filenames <- "img5.jpg"
	filenames <- "img6.jpg"
	filenames <- "img7.jpg"

	// this is very iportant to avoid deadlock!!!
	close(filenames)
	sizes := makeThumbnails3(filenames)
	fmt.Printf("\n\nThe sizes of files is %d\n", sizes)

}

func makeThumbnails1(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(fl string) {
			s, err := ImageFile(fl)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(s)
			ch <- struct{}{}
		}(f) //ignoring errors
	}

	//wait for goroutines to complete
	for range filenames {
		<-ch
	}

}
func makeThumbnails2(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(fl string) {
			var it item
			it.thumbfile, it.err = ImageFile(fl)
			ch <- it
			fmt.Println(it.thumbfile)
		}(f) //ignoring errors
	}

	//wait for goroutines to complete
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, err
}

func makeThumbnails3(filenames <-chan string) int64 {
	// number of iterations unknown ahead (channel as input)
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines

	for f := range filenames {
		fmt.Println("loop starts for file: " + f)
		wg.Add(1)
		fmt.Println("afterwg.Add(): " + f)
		// worker
		go func(f string) {
			defer func() {
				fmt.Println("wg.Done(): " + f)
				wg.Done()
			}()
			fmt.Println("getting thumb: " + f)
			thumb, err := ImageFile(f)
			if err != nil {
				fmt.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			fmt.Println("info: " + info.Name())
			s := fmt.Sprintf("%f", info.Size())
			sizes <- info.Size()
			fmt.Println("size in goroutine1 in file " + f + ": " + s)

		}(f)
	}

	var total int64
	go func() {
		// this goroutine starts only if closing the 'filenames' channel in the goroutin
		// close operation is a signal to the waitgroup
		fmt.Println("before wg.Wait()")
		wg.Wait()
		fmt.Println("before close(sizes)")
		close(sizes)
		fmt.Println("after close(sizes)")
	}()

	for size := range sizes {
		// it adds every size after each wg.Done() in the first goroutine
		total += size
		fmt.Println(fmt.Sprintf("Total size: %d", total))
	}

	return total
}

// Image returns a thumbnail-size version of src.
func Image(src image.Image) image.Image {
	// Compute thumbnail size, preserving aspect ratio.
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // portrait
	} else {
		height = int(128 / aspect) // landscape
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// a very crude scaling algorithm
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream reads an image from r and
// writes a thumbnail-size version of it to w.
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 reads an image from infile and writes
// a thumbnail-size version of it to outfile.
func ImageFile2(outfile, infile string) (err error) {

	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g. "foo.thumb.jpeg".
func ImageFile(infile string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	ext := filepath.Ext(infile) // e.g., ".jpg", ".JPEG"
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return path + "/out/" + outfile, ImageFile2(path+"/out/"+outfile, path+"/img/"+infile)
}
