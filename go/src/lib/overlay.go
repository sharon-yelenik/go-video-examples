package lib

import(
    "fmt"
    "github.com/cloudinary/cloudinary-go"
)

    func Overlay(cld *cloudinary.Cloudinary) string {
      // Instantiate an object for the video with public ID "exercise1" in folder "docs/sdk/go"
      v_exercise, err := cld.Video("docs/sdk/go/exercise1")
      if err != nil {
          fmt.Println("error")
      }

      // Add the transformation
      v_exercise.Transformation = "c_scale,w_300/l_video:exercise2/c_fit,w_80/bo_2px_solid_blue/fl_layer_apply,g_north_east,so_2.0"

      // Generate and print the delivery URL
      myURL, err := v_exercise.String()
      if err != nil {
          fmt.Println("error")
      }
      fmt.Println(myURL)

      // Output: https://res.cloudinary.com/demo/video/upload/c_scale,w_300/l_video:exercise2/c_fit,w_80/bo_2px_solid_blue/fl_layer_apply,g_north_east,so_2.0/v1/docs/sdk/go/exercise1  
    return myURL
}