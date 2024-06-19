require_relative './pover'

scene do
  line [-20, 0], [20, 0]
  line [0, -20], [0, 20]

  circle [3, 3], 1, 'Red'
  rectangle [-1, -1], [1, 3], 'Blue'
end
