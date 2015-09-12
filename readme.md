
# Tuts: wiki



This is step by step wiki build tutorial.
Check tags out.



##Start up steps



 Do not forget go get ...


    go get github.com/astaxie/beego
    
    go get github.com/beego/bee


Make sure you have a usable $GOPATH, and add $GOPATH/bin to your path.



In my case, ~/.bashrc is end up with:



    export GOPATH=$HOME/go
    
    export PATH=$PATH:$GOPATH/bin



So bee is in your path. 



 Suggest git clone repo to $GOPATH/src/

 cd into tuts_wiki, and "bee run", open browser with "http://127.0.0.1:8080"


## Todo

* Limit model cache num

* limit history length

* Crash if DB file not exist


## change log
v0.09: Optimize word history logic;

v0.08: Add link wiki word function.

v0.07: Choose a pretty editor.

v0.06: Add modification indication

v0.05: RESTful API support, implement content modification. 

* New router "/words/xxxx" to get wikiWord content; 

* Implement content modification via PUT method.

v0.04: Word content, relations, attributions views are done. Implement an usable history strategy

v0.03: Add login logic and history logic

* Use cookie to store email and password (this is suck, I know).

* Use session feature in beego to implement history(check it out @controllers/home.go).

v0.02: Add cover&login pages/controller and a simple logo.

v0.01: Basic index.html template.

ref doc:

    http://beego.me/quickstart

	http://v3.bootcss.com/