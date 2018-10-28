godocker-server-info
====================

Simple image that runs a local GO server to display docker info.

It uses `Go SDK`, so you need to mount the docker socket into container.


Start Tutorial
--------------

- [https://golang.org/doc/articles/wiki/](https://golang.org/doc/articles/wiki/)


NOTE
----

### Assets and Binary


Secondo la struttura attuale tutti i files reali sono in `./go/github.com/SimoSca/godocker-server-info` , che a sua volta contiene la `assets/templates/...`, cosa ganza in quanto io adoro l'uso dei templates!

Pero' ora inizia la difficolta':

se nel `main.go` uso un comando come 

````go
template.ParseFiles( "assets/templates/home.html" )
````

questo mi restituisce errore se lo compilo con `make build`, questo perche' di fatto la directory di esecuzione del mio comando `make` e' `./`, ma la directory del package invece e' `./go/src/github.com/SimoSca/godocker-server-info`, e proprio per questo non funzia!

Una soluzione per ovviare a questo problema e' creare la `assets/templates/...` nella `<project root>`, o meglio: dove verra' eseguito il comando d'avvio del server, in modo che si possano continuare a usare i path relativi!
Ad esempio potrei usare anche `go/src/github.com/SimoSca/godocker-server-info/assets/templates/...`, ma cosi' sarei vincolato a utilizzare il comando di istanziazione del server SOLO ed esclusivamente da questa cartella, quando eseguo `make build`; ad esempio lanciarlo dal `Desktop` porterebbe al solito errore!


Ora bisogna pensarci bene, perche' devo capire esattamente cosa devo fare col server, in quanto vi sono due approcci che posso utilizzare:

- pensare che sia da considerare per applicazioni esterne, ovvero prevedere la presenza dei vari files di template, ma lasciando che sia l'utilizzatore a crearli (ovvero che si crei la propria `assets/templates/...`), quindi come se fosse una applicazione embedded.

- pensare che il server sia a utilizzo monodirezionale (tipo `MailHog`), e che quindi gli assets devono in qualche modo rimanere vincolati alla procedura di compilazione, e quindi essere considerati come facenti parte di una applicazione standalone.


Questo se lo considero come isolato nel SO dell'utente, ma eventualmente si puo' pensare di svolgere un ibrido qualora si concepisse il tutto in un ambiente `docker`.


Attualmente visto che sto sviluppando il tutto come un server `standalone`, usero' [go packr](https://github.com/gobuffalo/packr), che consente di creare dei `box` che poi verranno buildati col package stesso.