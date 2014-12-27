#Stash

Simple utility to store and retrieve key:value pairs on the command line. Deeply inspired by the awesome: [boom](https://github.com/holman/boom)

To see a list of all your entries simply run: 

    $ stash   
    server1 => ssh ubuntu@host

  
To add a new entry called simple to your stash:

    $ stash simple test data  
    Added: simple => test data
    
To retrieve an entry called simple and have it copied to your clipboard:

    $ stash simple  
    Copied: test data
