      #include "colors.inc"

      background { color Black }

      camera {
        location <0, 0, -20>
        look_at  <0, 0, 0>
      }

            cylinder {
        <-20, 0, 0>, <20, 0, 0>, 0.1
        texture {
          pigment { color White }
        }
      }

      cylinder {
        <0, -20, 0>, <0, 20, 0>, 0.1
        texture {
          pigment { color White }
        }
      }

      box {
        <-8, 0, 0>, <-7, 3, 0>
        texture {
          pigment { color Red }
        }
      }

      box {
        <-6, 0, 0>, <-5, 3, 0>
        texture {
          pigment { color Orange }
        }
      }

      box {
        <-4, 0, 0>, <-3, 3, 0>
        texture {
          pigment { color Yellow }
        }
      }

      box {
        <-2, 0, 0>, <-1, 3, 0>
        texture {
          pigment { color Green }
        }
      }

      box {
        <0, 0, 0>, <1, 3, 0>
        texture {
          pigment { color Blue }
        }
      }

      box {
        <2, 0, 0>, <3, 3, 0>
        texture {
          pigment { color MidnightBlue }
        }
      }

      box {
        <4, 0, 0>, <5, 3, 0>
        texture {
          pigment { color Violet }
        }
      }

      box {
        <6, 0, 0>, <7, 3, 0>
        texture {
          pigment { color SpicyPink }
        }
      }


      light_source { <0, 0, -20> color White}
