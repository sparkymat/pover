require_relative './pover'

scene do
  # line [-8, -2], [-7, 0]
  # line [-7, 0], [-6, -2]
  # line [-7.5, -1], [-6.5, -1]
  #
  # line [-4.5, -2], [-5.5, 0]
  # line [-4.5, -2], [-3.5, 0]
  #
  # line [-3,0],[-2,-1]
  # line [-1,0],[-2,-1]
  # line [-2,-2],[-2,-1]

  line [-20, 0], [20, 0]
  line [0, -20], [0, 20]

  rectangle [-8, 0], [-7, 3], 'Red'
  rectangle [-6, 0], [-5, 3], 'Orange'
  rectangle [-4, 0], [-3, 3], 'Yellow'
  rectangle [-2, 0], [-1, 3], 'Green'
  rectangle [0, 0], [1, 3], 'Blue'
  rectangle [2, 0], [3, 3], 'MidnightBlue'
  rectangle [4, 0], [5, 3], 'Violet'
  rectangle [6, 0], [7, 3], 'SpicyPink'
end
