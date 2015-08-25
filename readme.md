
# Tuts: wiki



This is step by step wiki build tutorial.



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





## change log



V0.01: Basic index.html template.

ref doc:

    http://beego.me/quickstart

	http://v3.bootcss.com/