package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

func Trim(cld *cloudinary.Cloudinary) string {
    // Instantiate an object for the video with public ID "dog_mirror" in folder "docs/sdk/go"
    v_dog, err := cld.Video("docs/sdk/go/dog_mirror")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    v_dog.Transformation = "c_scale,w_450/du_5.0,so_3.5"


    // Generate and print the delivery URL
    myURL, err := v_dog.String()
    if err != nil {
        fmt.Println("error")
    }
    fmt.Println(myURL)

    // Output: https://res.cloudinary.com/demo/video/upload/c_scale,w_450/du_5.0,so_3.5/v1/docs/sdk/go/dog_mirror  
  return myURL
}