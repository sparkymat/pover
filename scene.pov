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

      cylinder {
        <3, 3, 0>, <3, 3, 1>, 1
        texture {
          pigment { color Red }
        }
      }

      box {
        <-1, -1, 0>, <1, 3, 0>
        texture {
          pigment { color Blue }
        }
      }


      light_source { <0, 0, -20> color White}
