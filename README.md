
## Features 

1. init
    - this should create the directory strucuture required for the .got directory
    - also initialize the cf file - this keeps track of the number of commits and the current commit we are in (useful for checkouts)

2. add
    - when a file is added for the first time, it should be added to the list of tracked files and also get staged at the same time
    - if a file is already tracked just add to the staged list

3. commit
    - get changes of all staged files
    - construct the file diffs and store the encoded edit string in the com directory under the 
    specific directory
    - these directories are use full when we want to checkout a specific commit
    - remove the staged file, this will clear the list of staged files resetting it to empty

4. status
    - get all files in the directory (ignoring .git and .got)
    - based on the staged and index files, display the status of files
    - currently doesnt check for changes in files - to be done later

5. prev_commit
    - constructs the previos commit of a given file
    - expand this to construct previous commits of all the files
    - eventually expand this to construction of any commit

6. Diff
    - given a file calculate and output its diff with the previous commit
