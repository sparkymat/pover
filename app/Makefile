all: watch-rb

watch-rb:
	reflex -s -r '\.rb$$' -- bash -c "bundle exec ruby scene.rb > scene.pov"

watch-pov:
	reflex -s -r '\.pov$$' -- povray scene.pov

watch:
	live-server
