out=//home/adithyaaaa/vsc/calcwebapp/bin


build:
	go build -o $(out)/calc /home/adithyaaaa/vsc/calcwebapp/main.go


run:build
	$(out)/calc

clean:build
	rm -f $(out)/calc