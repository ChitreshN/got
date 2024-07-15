# Todos

1. Figure out how to read files [DONE]
2. Make directory [DONE]
3. Calculating diffs [DONE]
    - Base case -empty file
    - Non tivial case
        - Howww
4. get changes save [DONE]
5. when asked for it build the thing to that commit

## Today [Really?, yeah not that day :)]

1. Commits
    - How save state of files [Done]
        - save the recent commited state
    - How calculate diff [Done]
        - compare current saved version with latest commit
    - How to save commit info
        - this turned out interesting
            - so create a string, all the info other than the identical line stuff is there
    - How to build [Trivial once the above two are done][well not really trivial, well it is, but cumbersome] [Done]

2. Types of files [Done]
    - untracked
    no saved version
    - Tracked
    has a saved version
        - staged
        there are some changes and they are added
        - unstaged
        changes but not added

3. each commit should get its own directory
    - this would help in constructing previous commits efficiently [Efficiently? nah easier] [Done]

4. Previous commit calculation must happen for all the tracked files [ ]

5. Checkout by specifiying a commit number [ ]



## Future? [Ill prolly never get to these, but lets write them down anyway]

1. Branches
2. Encoding files to keep the size small
3. More?
