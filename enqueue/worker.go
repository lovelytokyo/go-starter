package main

import (
        "os/exec"
        "os"
        "fmt"
        "errors"
        "github.com/benmanns/goworker"
)

//RPUSH resque:queue:default '{"class":"MyClass","args":["http://banner-mtb.dspcdn.com/mtbimg/video/bb5/bb59adddc40801a2f9fa10f2116d4185c56a0213"]}'
func myFunc(queue string, args ...interface{}) error {
        fmt.Printf("myFunc - From %s, %v\n", queue, args)

        movieUrl, ok := args[0].(string)
        if ok == false {
                msg := fmt.Sprintf("failed read testID. testID: %s\n", args[1])
                return errors.New(msg)
        }

        fmt.Printf("movieUrl: %s\n", movieUrl)
        cmd := fmt.Sprintf("eval $(ffprobe -v error -of flat=s=_ -select_streams v:0 -show_entries stream=height,width %s) && size=${streams_stream_0_width}x${streams_stream_0_height} && echo $size",
                movieUrl,
        )

        fmt.Printf("cmd: %s", cmd)

        out, err := exec.Command("sh", "-c", cmd).Output()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        fmt.Printf("== == string : %s\n", string(out))
        return nil
}

func yourFunc(queue string, args ...interface{}) error {
        fmt.Printf("yourFunc - From %s, %v\n", queue, args)
        return nil
}

func init() {
        // RPUSH resque:queue:default '{"class":"MyClass","args":["a123456","z95-MyClass"]}'
        goworker.Register("MyClass", myFunc)
        // RPUSH resque:queue:default '{"class":"YourClass","args":["a123456","z95-YourClass"]}'
        goworker.Register("YourClass", yourFunc)
}

func main() {
        if err := goworker.Work(); err != nil {
                fmt.Println("Error:", err)
        }
}
