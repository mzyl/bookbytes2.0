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
- [ ] Better layout - Buttons on the side, maybe
- [x] Next paragraph button
- [x] Previous paragraph button
- [ ] Beginning of chapter button
- [ ] Whole book in scrollable field

### Goals for Release
- [x] Able to display title and author
- [x] Random selection of books and paragraph
- [ ] Basic navigation through books i.e. forward, backward, beginning of chapter
- [ ] Solution for reading more of the book i.e. forward-backward are not so great for dialog
  - Scrollable field would be ideal
  - Provide where to find the book and roughly where the user is i.e. 38% through
- [ ] Basic web interface that isn't completely borrowed

### Problems
- Franklin Autobiography is a formatting mess.
- ~~Title formatting starts after first colon, makes bad things happen.~~
  - Frankenstein has a br tag that messes things up a bit..
- ~~Random file selection sometimes returns no file.~~
