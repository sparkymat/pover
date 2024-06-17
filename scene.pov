      #include "colors.inc"

      background { color Black }

      camera {
        location <0, 0, -20>
        look_at  <0, 0, 0>
      }

            sphere {
        <0, 0, 0>, 1
        texture {
          pigment { color Green }
        }
      }

      sphere {
        <2, 2, 0>, 1
        texture {
          pigment { color White }
        }
      }

      sphere {
        <0, 2, 0>, 1
        texture {
          pigment { color Blue }
        }
      }

      sphere {
        <2, 0, 0>, 1
        texture {
          pigment { color Red }
        }
      }


      light_source { <0, 20, -20> color White}
