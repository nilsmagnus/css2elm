# Work in progress

generate elm-source from your css-files

NOT TESTED PROPERLY, USE WITH CAUTION


## Install/use

	# install
    go install github.com/nilsmagnus/css2elm
    
	# use
	css2elm -input mystyles.css > ElmFromCss.elm
	

# Development & build

Requirements:

* go installed
* elm installed
* make installed

Build:

	# run go-tests and build binary:
    make 
	
	# run integration-test
	make integration-test
