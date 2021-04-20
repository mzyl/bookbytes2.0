## BookBytes 2.0
#### Short book passages to pique your interest
Things are coming together.
Interface is HEAVILY borrowed from the clickbaiter project by cbrgm found here: https://github.com/cbrgm/clickbaiter/

### Goals
- [x] Better data passing through functions
- [x] New buttons for displaying title and author
- [x] Build "booktext" out of "fulltext"
- [x] Maybe begin random selection of paragraphs and/or books
- [x] Random file selection
- [x] Select new book from ui
- [ ] Address some "problems"
- [x] Better layout - Buttons on the side, maybe
- [x] Next paragraph button
- [x] Previous paragraph button
- [x] Beginning of chapter button
- [ ] Display title/author and the end at beginning and end of book

### Goals for Release
- [x] Able to display title and author
- [x] Random selection of books and paragraph
- [x] Basic navigation through books i.e. forward, backward, beginning of chapter
- [x] Solution for reading more of the book i.e. forward-backward are not so great for dialog
  - Field is now scrollable
- [ ] Basic web interface that isn't completely borrowed
- [ ] Work with whole book library
- [ ] Serve locally on server

### Problems
- ~~Franklin Autobiography is a formatting mess.~~
- Chapters are not all marked the same.....
- ~~Title formatting starts after first colon, makes bad things happen.~~
  - Frankenstein has a br tag that messes things up a bit..
- ~~Random file selection sometimes returns no file.~~
- Same book is served to all users and changes made, i.e. new book/paragaph, effect all users
- Can't get to final chapter because adding 2 becomes out of range

### Thoughts
- Maybe we don't have a "scroll the whole book option"
  - Instead, we have a chapter by chapter option
    - "Read this chapter" | "Next Chapter" | "Previous Chapter" | "Start at the Beginning" (First Chapter)
    - Later we can have accounts to save places in books
- Should we see if struct construction can be done using go routines?
