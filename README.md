# Got

1. A custom vcs.
2. Better than available things? Nope.
3. Fastet? Prolly not.
4. Efficient in terms of file usage? If i tried to make it.
5. Whats the point of this then? It was fun to build, plus i had other things to do, doing this helped me not feel guilty about not doing those.

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
    - the previous here is prev commit - 1,i.e the last but one commit, the last commit is stored in the objFile (change name later)
    - expand this to construction of any commit

6. Diff
    - given a file calculate and output its diff with the previous commit

7. Revert
    - Revert files to the previous commit

8. Checkout
    - Given a commit number, changes the file contents of all tracked files to that commit state
    - Currently only works backwards, need to change it so that works other way too 
    

## How i built this

1. The first thought I had was something along these lines
    - The most important thing is to capture the changes in files
    - then storing these in a way, such that, any previous version can be calculated programatically

2. How to store changes
    - The obvious thing was to implement a file diff algorithm
    - and then design some sort of encoding mechanism to store these diffs

3. The Diff Algorithm
    - The base of this is the longest common subsequence problem
    - once that is calculated, we can use the dag, thats built when calculating the lcs and use it to get the changes [code here](lib/diff.go)

4. Adding and staging files
    - Done with the help of 2 files
    - index and staged, when a file is added for the first time its added to both these files. (check if first time by looping over the index file)
    - what should happen when we commit is explained in that section

5. Commit
    - The initial version was quite simple, get the current state of all the files in the staged file, and store them in the obj directory under the same filename
    - Once this is done remove the staged file, resetting the state.
    - But this sort of storage becomes difficult if we wish to retrace back to a previous state of the file, which is the purpose of building this after all

6. Storing commits
    - The solution i came up with is the following
    - Have some sort of way to track, how many number of commits are made to the repo
    - Using this, create a directory with the current commit number as its name
    - Under this directory, store the encoded diffs of the files which were staged for that particular commit [encoding algorith is here](lib/diff.go) [decoding is done seperately, should i refactore these? nah](lib/revert.go)
    - And then also update the data in the obj directory (obj folder contains the file in the most recent commit)

5. Revert
    - Just copy the contents in the obj file and update the correspondig files in the working directory

6. Prev_commit
    - Implementing this is, is the second most interesting thing
    - Get all files that are changed in the previous commit
    - Then construct the prev commit - 1, by decoding the commit string
    - change file based on the decoded data

7. Checkout
    - Same thing as prev commit, but keep track how many commits we have walked back
    - when we reach the required commit, call revert, this will update all the files to the required state
