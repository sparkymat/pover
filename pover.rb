Bundler.require

class Position
  attr_reader :x, :y, :z

  def initialize(x: 0, y: 0, z: 0)
    @x = x || 0
    @y = y || 0
    @z = z || 0
  end

  def render
    "<#{@x}, #{@y}, #{@z}>"
  end
end

class Camera
  attr_accessor :position, :lookat
end

class Light
  attr_accessor :position
end

class Sphere
  attr_accessor :center, :radius, :color

  def render
    <<-EOS
      sphere {
        #{@center.render}, #{@radius}
        texture {
          pigment { color #{@color} }
        }
      }
    EOS
  end
end

class Scene
  def initialize
    @default_fg_color = 'Yellow'
    @camera = Camera.new
    @camera.position = Position.new(x: 0, y: 0, z: -20)
    @camera.lookat = Position.new(x: 0, y: 0, z: 0)

    @objects = []

    default_light = Light.new
    default_light.position = Position.new(x: 0, y: 20, z: -20)
    @lights = [
      default_light
    ]
  end

  def camera(position, lookat)
    @camera.position = Position.new(x: position[0], y: position[1], z: position[2])
    @camera.lookat = Position.new(x: lookat[0], y: lookat[1], z: lookat[2])
  end

  def light(position)
    l = Light.new
    l.position = Position.new(x: position[0], y: position[1], z: position[2])
    @lights << l
  end

  def sphere(center, radius, color = '')
    s = Sphere.new
    s.center = Position.new(x: center[0], y: center[1], z: center[2])
    s.radius = radius
    s.color = color.nil? || color == '' ? @default_fg_color : color
    @objects << s
  end

  def circle(center, radius, color = '')
    sphere(center, radius, color)
  end

  def render
    <<-TEMPL
      #include "colors.inc"

      background { color Black }

      camera {
        location #{@camera.position.render}
        look_at  #{@camera.lookat.render}
      }

      #{@objects.map(&:render).join("\n")}

      #{
        @lights.map do |light|
          "light_source { #{light.position.render} color White}"
        end.join("\n")
      }
    TEMPL
  end
end

def scene(&block)
  s = Scene.new
  s.instance_eval(&block)
  puts s.render
end
