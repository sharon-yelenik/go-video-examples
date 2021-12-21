package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Resize_crop(cld *cloudinary.Cloudinary) string {
    // Instantiate an object for the video with public ID "play_time" in folder "docs/sdk/go"
    v_play_time, err := cld.Video("docs/sdk/go/play_time")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    v_play_time.Transformation = "b_blurred:400:15,c_pad,h_320,w_480"


    // Generate and print the delivery URL
    myURL, err := v_play_time.String()
    if err != nil {
        fmt.Println("error")
    }
    fmt.Println(myURL)

    // Output: https://res.cloudinary.com/demo/video/upload/b_blurred:400:15,c_pad,h_320,w_480/v1/docs/sdk/go/play_time
  return myURL
}